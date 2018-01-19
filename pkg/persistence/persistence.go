package persistence

import "github.com/owainlewis/kcd/pkg/types"

type JobStore interface {
	GetJob(id string) (*types.Job, error)
	SaveJob(job types.Job) error
}

type ProjectStore interface {
}

// Datastore defines the interface for persisting types
type Datastore interface {
	JobStore
}
