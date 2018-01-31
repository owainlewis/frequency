package api

import (
	"github.com/owainlewis/frequency/pkg/executor"
)

type Api struct {
	Executor executor.Executor
}

func New(executor executor.Executor) Api {
	return Api{
		Executor: executor,
	}
}
