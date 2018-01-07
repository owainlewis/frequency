package executor

import (
	types "github.com/owainlewis/kcd/pkg/types"
)

type JobExecutor interface {
	Execute(job types.Job) error
}
