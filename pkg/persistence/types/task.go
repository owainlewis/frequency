package persistence

import (
	tasks "github.com/owainlewis/frequency/pkg/tasks"
)

type Task struct {
	Created  string
	Started  string
	Finished string
	Spec     tasks.Task
}
