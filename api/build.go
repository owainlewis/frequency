package api

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// CreateBuild will create a new build.
// POST /api/v1/project/:id/builds
func (api Api) CreateBuild(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	project, ok := vars["project"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	glog.Infof("Building project %s", project)

	w.WriteHeader(http.StatusAccepted)
}
