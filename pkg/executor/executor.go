package executor

import "github.com/owainlewis/frequency/pkg/tasks"

type TaskExecutor interface {
	Execute(task tasks.PodTask) error
}
