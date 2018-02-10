package store

import (
	"github.com/owainlewis/frequency/pkg/types"
)

type TaskStore interface {
	GetTasks() ([]types.Task, error)
	CreateTask(tasks types.Task) error
	UpdateTask(id int, task types.Task) error
}

type Store interface {
	TaskStore
}
