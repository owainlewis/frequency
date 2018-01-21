package api

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/owainlewis/frequency/pkg/types"
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

	job.ID = string(pod.GetUID())

	// We store the job to our persistent store
	// The Kubernetes controller for Pods will watch the state of this
	// pod as it progresses and update the job status
	glog.Infof("Saving job %s to database", job.ID)
	api.Store.CreateJob(&job)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&pod)
}
