package persistence

import (
	"github.com/owainlewis/frequency/pkg/types"
)

// Task represents the information that needs to be persisted for a
// given task. Notably this includes the task definition but also metadata
// about the start/end times and information about the underlying K8s Pod.
type Task struct {
	CreatedAt  string
	StartedAt  string
	FinishedAt string
	PodID      string
	Spec       types.Task
}
