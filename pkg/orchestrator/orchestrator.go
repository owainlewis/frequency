package orchestrator

import (
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

func (o Orchestrator) CreatePod(namespace string, image string, commands []string) (*v1.Pod, error) {
	template := newPod(image, commands)

	return o.clientset.CoreV1().Pods(namespace).Create(template)
}

func newPod(image string, commands []string) *v1.Pod {
	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name:    "mycontainer",
				Image:   image,
				Command: commands,
			}},
			RestartPolicy: v1.RestartPolicyOnFailure,
		},
	}

	pod.SetGenerateName("kcd-")

	return pod

}
