package persistence

import "github.com/owainlewis/kcd/pkg/types"

// Datastore defines the interface for persisting types
type Datastore interface {
	GetJob(id string) (*types.Job, error)
	SaveJob(job types.Job) error
}
