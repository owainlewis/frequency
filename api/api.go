package api

import (
	"github.com/owainlewis/kcd/pkg/executor"
	"github.com/owainlewis/kcd/pkg/persistence"
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
