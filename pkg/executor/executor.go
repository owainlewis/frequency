package executor

import (
	types "github.com/owainlewis/frequency/pkg/types"
)

type JobExecutor interface {
	Execute(job types.Job) error
}
