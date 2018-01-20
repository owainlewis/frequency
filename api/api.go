package api

import (
	"github.com/owainlewis/frequency/pkg/executor"
	"github.com/owainlewis/frequency/pkg/persistence"
)

type Api struct {
	Executor executor.Executor
	Store    persistence.Datastore
}

func New(executor executor.Executor, store persistence.Datastore) Api {
	return Api{
		Executor: executor,
		Store:    store,
	}
}
