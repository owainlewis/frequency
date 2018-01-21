package persistence

import (
	"fmt"

	"github.com/owainlewis/frequency/pkg/types"
)

// InMemoryStore defines an in memory implementation of
// the persistence inteface
type InMemoryStore struct {
	Jobs []*types.Job
}

// NewInMemoryStore constructs a new InMemoryStore
func NewInMemoryStore() InMemoryStore {
	jobs := []*types.Job{}
	return InMemoryStore{Jobs: jobs}
}

// FindJobByID will return a job if a matching ID is found
func (s InMemoryStore) FindJobByID(ID string) (*types.Job, error) {
	return nil, fmt.Errorf("failed to find job with ID %s", ID)
}

// CreateJob will persist a job to an in memory collection
func (s InMemoryStore) CreateJob(job *types.Job) error {
	s.Jobs = append(s.Jobs, job)
	return nil
}

func (s InMemoryStore) UpdateJob(job *types.Job) (*types.Job, error) {
	return nil, nil
}
