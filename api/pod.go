package api

import (
	"encoding/json"
	"net/http"

	"github.com/owainlewis/frequency/pkg/tasks"
)

// CreatePodTask ...
func (api Api) CreatePodTask(w http.ResponseWriter, r *http.Request) {
	var task tasks.PodTask

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = api.Executor.Execute(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//task.ID = string(pod.GetUID())

	w.WriteHeader(http.StatusAccepted)
	//json.NewEncoder(w).Encode(&pod)
}
