package orchestrator

import (
	"github.com/golang/glog"
	types "github.com/owainlewis/kcd/pkg/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// Should be configurable
const shell = "/bin/bash"

// Executor controls how jobs are executed inside Kubernetes
type Executor struct {
	clientset kubernetes.Interface
}

// NewExecutor creates a properly configured Job Executor
func NewExecutor(clientset kubernetes.Interface) Executor {
	return Executor{clientset: clientset}
}

// Execute will execute a single job
func (o Executor) Execute(job types.Job) error {
	glog.Infof("Executing job: ", job.Name)
	template := newPod(job.Image, job.Commands)

	// TODO which namespace to run in (must be configurable)
	pod, err := o.clientset.CoreV1().Pods(v1.NamespaceDefault).Create(template)
	glog.Infof("Created pod %s for execution", pod.Name)

	return err
}

func (o Executor) createPod(namespace string, image string, commands []string) (*v1.Pod, error) {
	template := newPod(image, commands)

	return o.clientset.CoreV1().Pods(namespace).Create(template)
}

// Hacky way to do this. TODO can we refine here?
func formatCommands(commands []string) []string {
	cmdStr := ""
	for _, c := range commands {
		cmdStr = cmdStr + c + ";"
	}

	return []string{shell, "-c", cmdStr}
}

func newPod(image string, commands []string) *v1.Pod {
	primary := v1.Container{
		Name:       "primary",
		Image:      image,
		WorkingDir: "/workspace",
		Env: []v1.EnvVar{{
			Name:  "OUTPUT_DIR",
			Value: "/output",
		}},
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "workspace",
				MountPath: "/workspace",
			},
			{
				Name:      "output",
				MountPath: "/output",
			}},
		Command: formatCommands(commands),
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"owner": "kcd",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{primary},
			InitContainers: []v1.Container{{
				Name:  "setup",
				Image: "alpine/git",
				VolumeMounts: []v1.VolumeMount{{
					Name:      "workspace",
					MountPath: "/workspace",
				}},
				Command: []string{
					"ash", "-c", "git clone https://github.com/owainlewis/hello-spinnaker.git /workspace/",
				},
			}},
			RestartPolicy: v1.RestartPolicyNever,
			Volumes: []v1.Volume{
				{
					Name: "workspace",
				},
				{
					Name: "output",
				},
			},
		},
	}

	pod.SetGenerateName("kcd-")

	return pod

}
