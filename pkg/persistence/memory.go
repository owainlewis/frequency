package persistence

import "github.com/owainlewis/kcd/pkg/types"

type InMemoryStore struct {
	jobs []types.Job
}

func (s InMemoryStore) GetJob(id string) (*types.Job, error) {
	return nil, nil
}

func (s InMemoryStore) SaveJob(job types.Job) error {
	return nil
}
