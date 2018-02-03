package executor

import "github.com/owainlewis/frequency/pkg/types"

type Executor interface {
	ExecuteTask(task types.Task) error
}
