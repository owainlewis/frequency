package api

import (
	"github.com/owainlewis/frequency/pkg/executor"
)

type Api struct {
	Executor executor.TaskExecutor
}

func New(executor executor.TaskExecutor) Api {
	return Api{
		Executor: executor,
	}
}
