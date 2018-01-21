package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/owainlewis/frequency/pkg/types"
)

type MySQLStore struct {
}

func NewMySQLStore() *MySQLStore {
	db, err := sql.Open("mysql", "user:password@/dbname")

	return nil
}

// GetJob will return a job if a matching ID is found
func (s *MySQLStore) GetJob(ID string) (*types.Job, error) {
	return nil, fmt.Errorf("failed to find job with ID %s", ID)
}

// SaveJob will persist a job to an in memory collection
func (s *MySQLStore) SaveJob(job *types.Job) error {
	return nil
}
