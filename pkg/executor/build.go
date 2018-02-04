package executor

import (
	kubernetes "k8s.io/client-go/kubernetes"

	"github.com/owainlewis/frequency/pkg/types"
)

type DefaultBuildExecutor struct {
	Client kubernetes.Interface
}

func NewDefaultBuildExecutor(clientset kubernetes.Interface) DefaultBuildExecutor {
	return DefaultBuildExecutor{Client: clientset}
}

func (e DefaultBuildExecutor) ExecuteBuild(task types.Build) error {
	return nil
}
