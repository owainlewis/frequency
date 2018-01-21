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

// FindJobByID will return a job if a matching ID is found
func (s InMemoryStore) FindJobByID(ID string) (*types.Job, error) {
	for _, job := range s.Jobs {
		if job.ID == ID {
			return job, nil
		}
	}
	return nil, fmt.Errorf("failed to find job with ID %s", ID)
}

// CreateJob will persist a job to an in memory collection
func (s InMemoryStore) CreateJob(job *types.Job) error {
	s.Jobs = append(s.Jobs, job)
	return nil
}
