package executor

import (
	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	tasks "github.com/owainlewis/frequency/pkg/tasks"
)

// PodTaskExecutor ...
type PodTaskExecutor struct {
	Client kubernetes.Interface
}

// NewPodTaskExecutor creates a properly configured PodTaskExecutor
func NewPodTaskExecutor(clientset kubernetes.Interface) PodTaskExecutor {
	return PodTaskExecutor{Client: clientset}
}

// Execute will execute a single job
func (e PodTaskExecutor) Execute(task tasks.PodTask) error {
	glog.Infof("Executing Pod task: %+v", task)

	taskPod := e.newPod(task)

	// TODO which namespace to run in (must be configurable)
	_, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(taskPod)
	if err != nil {
		glog.Infof("Failed to create Pod: %s", err)
		return err
	}

	return nil
}

func env(k, v string) v1.EnvVar {
	return v1.EnvVar{Name: k, Value: v}
}

func (e PodTaskExecutor) newPod(task tasks.PodTask) *v1.Pod {
	primary := v1.Container{
		Name:       "primary",
		Image:      task.Image,
		WorkingDir: task.Workspace,
		Env:        task.Env,
		Command:    []string{task.Command.Cmd},
		Args:       task.Command.Args,
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{},
		},
		Spec: v1.PodSpec{
			Containers:    []v1.Container{primary},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}

	pod.SetGenerateName("task-")

	return pod
}
