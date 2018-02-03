package executor

import (
	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"github.com/owainlewis/frequency/pkg/types"
)

// TaskExecutor ...
type TaskExecutor struct {
	Client kubernetes.Interface
}

// NewTaskExecutor creates a properly configured TaskExecutor
func NewTaskExecutor(clientset kubernetes.Interface) TaskExecutor {
	return TaskExecutor{Client: clientset}
}

// ExecuteTask will execute a single task and return an error if it cannot be executed
func (e TaskExecutor) ExecuteTask(task types.Task) error {
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

func (e TaskExecutor) newPod(task types.Task) *v1.Pod {
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

	glog.Infof("Task %+v", task)

	if task.Source != nil {

		glog.Infof("Cloning code from %s", task.Source.GitURL)

		cloneCommand := "git clone $GIT_SOURCE $WORKSPACE"
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

func buildEnvironmentVariables(task *types.Task) []v1.EnvVar {
	var env []v1.EnvVar
	// Project and general information
	// FREQUENCY_PROJECT_NAME
	// Task information
	// FREQUENCY_TASK_ID
	// FREQUENCY_TASK_WORKSPACE
	// Git Information
	// FREQUENCY_GIT_DOMAIN="github.com"
	// FREQUENCY_GIT_OWNER="oracle"
	// FREQUENCY_GIT_REPOSITORY="terraform-kubernetes-installer"
	// FREQUENCY_GIT_BRANCH="master"
	// FREQUENCY_GIT_COMMIT="4fc26b093db08a6079e27016d1903b66aa93604b"
	return env
}
