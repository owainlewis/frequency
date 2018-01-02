package orchestrator

import (
	"github.com/golang/glog"

	types "github.com/owainlewis/kcd/pkg/types"
	v1 "k8s.io/api/core/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

type Orchestrator struct {
	clientset kubernetes.Interface
}

// NewOrchestrator creates a properly configured Job orchestrator
func NewOrchestrator(clientset kubernetes.Interface) Orchestrator {
	return Orchestrator{clientset: clientset}
}

func (o Orchestrator) ExecuteStage(namespace string, stage types.Stage) error {
	glog.Infof("Executing stage: ", stage.Name)
	template := newPod(stage.Image, stage.Commands)

	// TODO which namespace to run in (must be configurable)
	_, err := o.clientset.CoreV1().Pods(namespace).Create(template)

	return err
}

func (o Orchestrator) createPod(namespace string, image string, commands []string) (*v1.Pod, error) {
	template := newPod(image, commands)

	return o.clientset.CoreV1().Pods(namespace).Create(template)
}

func newPod(image string, commands []string) *v1.Pod {
	primary := v1.Container{
		Name:    "primary",
		Image:   image,
		Command: []string{"/bin/sh", "-c", "tail -f /dev/null"},
	}

	agent := v1.Container{
		Name:    "agent",
		Image:   "ubuntu",
		Command: []string{"/bin/sh", "-c", "tail -f /dev/null"},
	}

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers:    []v1.Container{primary, agent},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}

	pod.SetGenerateName("kcd-")

	return pod

}
