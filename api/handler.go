package api

import (
	"encoding/json"
	"net/http"

	"github.com/owainlewis/kcd/pkg/types"
)

// CreateJob is a HTTP handler invoked to trigger the execution of a job
func (api Api) CreateJob(w http.ResponseWriter, r *http.Request) {
	var job types.Job
	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = api.Executor.Execute(&job)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&job)
}
