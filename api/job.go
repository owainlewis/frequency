package api

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/owainlewis/kcd/pkg/types"
)

// CreateJob is a HTTP handler invoked to trigger the execution of a single job
func (api Api) CreateJob(w http.ResponseWriter, r *http.Request) {
	var job types.Job

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	job.EnsureDefaults()

	pod, err := api.Executor.Execute(&job)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pid := pod.GetUID()

	job.ID = string(pid)

	glog.Infof("Saving job %v", job)

	// Write job to database and set the pod ID as the

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&pod)
}
