package executor

import (
	"fmt"

	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	tasks "github.com/owainlewis/frequency/pkg/tasks"
)

// PodTaskExecutor ...
type WaitTaskExecutor struct {
	Client kubernetes.Interface
}

// NewWaitTaskExecutor creates a properly configured WaitTaskExecutor
func NewWaitTaskExecutor(clientset kubernetes.Interface) WaitTaskExecutor {
	return WaitTaskExecutor{Client: clientset}
}

// Execute will execute a single job
func (e WaitTaskExecutor) Execute(task tasks.Task) error {
	glog.Infof("Executing Wait task: %+v", task)

	if task.GetKind() != "WaitTask" {
		return fmt.Errorf("Invalid task kind for WaitTaskExecutor")
	}

	podTask := task.(tasks.WaitTask)
	taskPod := e.newPod(podTask)

	// TODO which namespace to run in (must be configurable)
	_, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(taskPod)
	if err != nil {
		glog.Infof("Failed to create Pod: %s", err)
		return err
	}

	return nil
}

func (e WaitTaskExecutor) newPod(task tasks.WaitTask) *v1.Pod {
	primary := v1.Container{
		Name:    "wait",
		Image:   "busybox",
		Command: []string{"sh", "-c", "sleep " + string(task.Duration)},
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
