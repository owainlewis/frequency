package executor

import (
	"github.com/golang/glog"
	types "github.com/owainlewis/kcd/pkg/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// Should be configurable (part of the job?)
const shell = "/bin/bash"

// Executor controls how jobs are executed inside Kubernetes
type Executor struct {
	Client kubernetes.Interface
}

// NewExecutor creates a properly configured Job Executor
func NewExecutor(clientset kubernetes.Interface) Executor {
	return Executor{Client: clientset}
}

// Execute will execute a single job
func (e Executor) Execute(job *types.Job) (*v1.Pod, error) {
	job.EnsureDefaults()

	glog.Infof("Executing job: %s", job.Name)
	template := e.NewJobExecutionPod(job)

	// TODO which namespace to run in (must be configurable)
	pod, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(template)
	if err != nil {
		glog.Infof("Failed to create pod")
		return nil, err
	}

	return pod, nil
}

// FormatSteps will format the steps passed for execution in a shell
// Hacky way to do this. TODO can we refine here?
func (e Executor) FormatSteps(steps []string) []string {
	stepsStr := ""
	for _, c := range steps {
		stepsStr = stepsStr + c + ";"
	}

	return []string{shell, "-c", stepsStr}
}

func env(k, v string) v1.EnvVar {
	return v1.EnvVar{Name: k, Value: v}
}

// NewJobExecutionPod will construct the pod template in which a job will run
//
// The pod has two primary roles
//
// 1. A sidecar container will check out the project from Git
// 2. A primary pod is created with the correct shared directories and env
//
//
func (e Executor) NewJobExecutionPod(job *types.Job) *v1.Pod {
	defaultEnv := []v1.EnvVar{
		env("WORKSPACE", job.Workspace),
		env("OUTPUT_DIR", "/output"),
	}
	primary := v1.Container{
		Name:       "primary",
		Image:      job.Image,
		WorkingDir: job.Workspace,
		Env:        defaultEnv,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "workspace",
				MountPath: job.Workspace,
			},
			{
				Name:      "output",
				MountPath: "/output",
			}},
		Command: e.FormatSteps(job.Steps),
	}

	buildEnv := []v1.EnvVar{
		env("GIT_PROJECT", "https://github.com/owainlewis/sample-project.git"),
		env("GIT_BRANCH", "master"),
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
				Env:   append(defaultEnv, buildEnv...),
				VolumeMounts: []v1.VolumeMount{{
					Name:      "workspace",
					MountPath: job.Workspace,
				}},
				Command: []string{
					"ash", "-c", "git clone -b $GIT_BRANCH $GIT_PROJECT $WORKSPACE",
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
