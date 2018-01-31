package executor

import "github.com/owainlewis/frequency/pkg/tasks"

type Executor interface {
	Execute(task tasks.Task) error
}
