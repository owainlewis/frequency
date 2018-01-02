package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Build will construct a Kubernetes clientset from a given
// config path. If the path is empty then it will default to use
// an in-cluster configuration
func Build(configpath string) (*kubernetes.Clientset, error) {
	config, err := getKubeConfig(configpath)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getKubeConfig(configpath string) (*rest.Config, error) {
	if configpath != "" {
		return clientcmd.BuildConfigFromFlags("", configpath)
	}

	return rest.InClusterConfig()
}
