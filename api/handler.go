package api

import (
	"encoding/json"
	"net/http"

	"github.com/owainlewis/kcd/pkg/types"
)

func (api Api) CreateJob(w http.ResponseWriter, r *http.Request) {

	job := types.Job{
		Name:  "hello-kcd",
		Image: "golang",
		Commands: []string{
			"go build -v main.go",
			"mv ./main $OUTPUT_DIR",
			"ls $OUTPUT_DIR",
		},
	}

	err := api.Executor.Execute(&job)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&job)
}

// // TriggerJob is the API endpoint to start the execution of a single job
// func (api *Api) TriggerJob(w *http.ResponseWriter, r *http.Request) {
// 	exec := executor.NewExecutor(api.Client)

// 	job := types.Job{
// 		Name:  "hello-kcd",
// 		Image: "golang",
// 		Commands: []string{
// 			"go build -v main.go",
// 			"mv ./main $OUTPUT_DIR",
// 			"ls $OUTPUT_DIR",
// 		},
// 	}

// 	err := exec.Execute(&job)
// 	if err != nil {
// 		glog.Errorf("Execution failed: %s", err.Error())
// 	}
// }
