package orchestrator

import (
	"io"
	"os"

	glog "github.com/golang/glog"
	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

type Orchestrator struct {
	clientset kubernetes.Interface
}

// NewOrchestrator creates a properly configured Job orchestrator
func NewOrchestrator(clientset kubernetes.Interface) Orchestrator {
	return Orchestrator{clientset: clientset}
}

// CreateJob will launch a new Kubernetes job in which to execute the user commands which are
// derived from a kcd.yaml file
func (o Orchestrator) CreateJob(namespace string, image string, commands []string) (*v1.Job, error) {
	template := newJob(image, commands)
	job, err := o.clientset.BatchV1().Jobs(namespace).Create(template)

	if err != nil {
		return nil, err
	}

	glog.Infof("Job is %v", job)

	err = o.Logs(namespace, "mycontainer")

	if err != nil {
		glog.Errorf("Failed to get logs: %s", err.Error())
	}

	return job, nil
}

// newJob creates a new Kubernetes Job which is the primary unit of execution in KCD
func newJob(image string, commands []string) *v1.Job {
	return &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "myjob3",
		},
		Spec: v1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-job-pod",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:    "mycontainer",
						Image:   image,
						Command: commands,
					}},
					RestartPolicy: corev1.RestartPolicyOnFailure,
				},
			},
		},
	}
}

// Logs will return a stream of logs from the job contatiner
// TODO this will require a lot of design
func (o Orchestrator) Logs(namespace, container string) error {
	req := o.clientset.CoreV1().Pods(namespace).GetLogs("mycontainer", &corev1.PodLogOptions{})
	readCloser, err := req.Stream()
	if err != nil {
		return err
	}

	defer readCloser.Close()
	_, err = io.Copy(os.Stdout, readCloser)
	return err
}
