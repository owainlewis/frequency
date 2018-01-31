package main

import (
	"flag"
	"log"
	"net/http"

	"k8s.io/client-go/kubernetes"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/owainlewis/frequency/api"
	builder "github.com/owainlewis/frequency/pkg/client"
	"github.com/owainlewis/frequency/pkg/controller"
	"github.com/owainlewis/frequency/pkg/executor"
)

var kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")

var banner = `
#
#    ___
#   /
#  (___  ___  ___  ___       ___  ___  ___
#  |    |   )|___)|   )|   )|___)|   )|    \   )
#  |    |    |__  |__/||__/ |__  |  / |__   \_/
#                     |                      /
#
`

func main() {

	flag.Parse()

	client, err := builder.Build(*kubeconfig)
	if err != nil {
		glog.Errorf("Failed to build client: %s", err)
		return
	}

	ctrl := controller.NewController(client)

	stop := make(chan struct{})

	defer close(stop)
	go ctrl.Run(1, stop)

	API := buildAPI(client)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/tasks", API.CreateTask).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	glog.Info("Starting API server...")
	glog.Info(banner)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func buildAPI(client kubernetes.Interface) api.Api {
	ex := executor.NewTaskExecutor(client)
	return api.New(ex)
}
