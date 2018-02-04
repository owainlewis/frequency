package persistence

import (
	"github.com/owainlewis/frequency/pkg/types"
)

// Build represents the execution of a Task
type Build struct {
	ID         string
	CreatedAt  string
	StartedAt  string
	FinishedAt string
	PodID      string
	Spec       types.Task
}
