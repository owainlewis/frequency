package api

import (
	"encoding/json"
	"net/http"

	"github.com/owainlewis/frequency/pkg/types"
	"github.com/owainlewis/frequency/pkg/validation"
)

type errorResponse struct {
	Error string `json:"error"`
}

// CreateTask will create a new task. It will be executed asynchronously
// POST /api/v1/tasks
func (api Api) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task types.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

	task.SetDefaults()

	errs := task.Validate()
	if len(errs) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{Error: validation.ConsolidateErrors(errs)})
		return
	}

	err = api.Executor.TaskExecutor.ExecuteTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// GetTask will return a single task if it exists
// GET /api/v1/task/:id
func (api Api) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

// GetTasks returns a list of all active tasks
// GET /api/v1/tasks
func (api Api) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}
