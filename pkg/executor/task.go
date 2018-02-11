package executor

import (
	"fmt"

	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"strings"

	"github.com/owainlewis/frequency/pkg/types"
)

// DefaultTaskExecutor ...
type DefaultTaskExecutor struct {
	Client kubernetes.Interface
}

// NewDefaultTaskExecutor creates a properly configured DefaultTaskExecutor
func NewDefaultTaskExecutor(clientset kubernetes.Interface) DefaultTaskExecutor {
	return DefaultTaskExecutor{Client: clientset}
}

// ExecuteTask will execute a single task and return an error if it cannot be executed
func (e DefaultTaskExecutor) ExecuteTask(task types.Task) error {
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

func (e DefaultTaskExecutor) newPod(task types.Task) *v1.Pod {
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
				MountPath: task.Workspace,
			},
		},
	}

	var initContainers []v1.Container

	if task.Checkout != nil {
		checkoutCommand := prepareCheckout(task.Checkout)

		sourceCloneContainer := v1.Container{
			Name:  "setup",
			Image: "alpine/git:latest",
			VolumeMounts: []v1.VolumeMount{{
				Name:      "workspace",
				MountPath: task.Workspace,
			}},
			WorkingDir: task.Workspace,
			Command:    []string{"ash", "-c"},
			Args:       []string{checkoutCommand},
		}

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

func envVar(name, value string) v1.EnvVar {
	return v1.EnvVar{
		Name:  name,
		Value: value,
	}
}

func buildEnvironmentVariables(task *types.Task) []v1.EnvVar {
	var env []v1.EnvVar

	env = append(env, envVar("FREQUENCY_TASK_WORKSPACE", task.Workspace))

	if task.Checkout != nil {
		env = append(env, envVar("FREQUENCY_CHECKOUT_URL", task.Checkout.URL))
	}

	return env
}

// Builds the correct checkout command including any post actions
func prepareCheckout(checkout *types.Checkout) string {
	var args []string

	cloneCommand := fmt.Sprintf("git clone %s", checkout.URL)
	args = append(args, cloneCommand)

	for _, post := range checkout.Post {
		args = append(args, post)
	}

	return strings.Join(args, "; ")
}
