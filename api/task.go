package api

import (
	"encoding/json"
	"net/http"

	"github.com/owainlewis/frequency/pkg/tasks"
)

// CreateTask ...
func (api Api) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	task.SetDefaults()

	err = api.Executor.Execute(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
