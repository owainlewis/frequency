package exec

import (
	"bytes"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	scheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	remotecommand "k8s.io/client-go/tools/remotecommand"
)

type Executor struct {
	config    *rest.Config
	clientset kubernetes.Interface
}

func NewExecutor(config *rest.Config, clientset kubernetes.Interface) Executor {
	return Executor{config: config, clientset: clientset}
}

// Command executes arbitrary commands inside a pod
func (e Executor) Command(namespace string, podName string, command ...string) (string, string, error) {

	var execOut bytes.Buffer
	var execErr bytes.Buffer

	pod, err := e.clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return "", "", fmt.Errorf("could not get pod info: %v", err)
	}

	req := e.clientset.Core().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")
	req.VersionedParams(&v1.PodExecOptions{
		Container: pod.Spec.Containers[0].Name,
		Command:   command,
		Stdout:    true,
		Stderr:    true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(e.config, "POST", req.URL())
	if err != nil {
		return "", "", fmt.Errorf("failed to init executor: %v", err)
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: &execOut,
		Stderr: &execErr,
	})

	if err != nil {
		return execOut.String(), execErr.String(), fmt.Errorf("could not execute: %v", err)
	}

	return execOut.String(), execErr.String(), err
}
