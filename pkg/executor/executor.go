package executor

import (
	"github.com/owainlewis/frequency/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type TaskExecutor interface {
	ExecuteTask(task types.Task) error
}

type BuildExecutor interface {
	ExecuteBuild(build types.Build) error
}

type Executor struct {
	BuildExecutor BuildExecutor
	TaskExecutor  TaskExecutor
}

func NewDefaultExecutor(client kubernetes.Interface) *Executor {
	buildExecutor := NewDefaultBuildExecutor(client)
	taskExecutor := NewDefaultTaskExecutor(client)

	return &Executor{
		BuildExecutor: buildExecutor,
		TaskExecutor:  taskExecutor,
	}
}
