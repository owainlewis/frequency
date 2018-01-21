package persistence

import "github.com/owainlewis/frequency/pkg/types"

type JobStore interface {
	FindJobByID(id string) (*types.Job, error)
	CreateJob(job *types.Job) error
}

type ProjectStore interface {
}

// Datastore defines the interface for persisting types
type Datastore interface {
	JobStore
}
