package executor

import (
	"fmt"

	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	tasks "github.com/owainlewis/frequency/pkg/tasks"
)

// TaskExecutor ...
type TaskExecutor struct {
	Client kubernetes.Interface
}

// NewTaskExecutor creates a properly configured TaskExecutor
func NewTaskExecutor(clientset kubernetes.Interface) TaskExecutor {
	return TaskExecutor{Client: clientset}
}

// Execute will execute a single job
func (e TaskExecutor) Execute(task tasks.Task) error {
	glog.Infof("Executing task: %+v", task)

	taskPod := e.newPod(task)

	// TODO which namespace to run in (must be configurable)
	_, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(taskPod)
	if err != nil {
		glog.Infof("Failed to create Pod: %s", err)
		return err
	}

	return nil
}

func (e TaskExecutor) newPod(task tasks.Task) *v1.Pod {
	primary := v1.Container{
		Name:       "primary",
		Image:      task.Image,
		WorkingDir: task.Workspace,
		Env:        task.Env,
		Command:    task.Run.Command,
		Args:       task.Run.Args,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "workspace",
				MountPath: "/workspace",
			},
		},
	}

	// 	// // When a source is declared as part of a job, we use an init container
	// 	// // to go and fetch that source code from a VCS such as github.com
	// 	// var initContainers []v1.Container
	var initContainers []v1.Container

	if task.Source != nil {

		glog.Infof("Cloning code from %s", task.Source.GitURL)

		cloneCommand := fmt.Sprintf("git clone %s /workspace", task.Source.GitURL)
		sourceCloneContainer := v1.Container{
			Name:  "setup",
			Image: "alpine/git",
			VolumeMounts: []v1.VolumeMount{{
				Name:      "workspace",
				MountPath: task.Workspace,
			}},
			Command: []string{"ash", "-c", cloneCommand}}

		initContainers = append(initContainers, sourceCloneContainer)
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{},
		},
		Spec: v1.PodSpec{
			Containers:     []v1.Container{primary},
			InitContainers: initContainers,
			RestartPolicy:  v1.RestartPolicyNever,
			Volumes: []v1.Volume{
				{
					Name: "workspace",
				},
			},
		},
	}

	pod.SetGenerateName("task-")

	return pod
}
