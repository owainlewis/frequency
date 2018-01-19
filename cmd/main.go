package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/owainlewis/kcd/api"
	builder "github.com/owainlewis/kcd/pkg/client"
	"github.com/owainlewis/kcd/pkg/controller"
	"github.com/owainlewis/kcd/pkg/executor"
)

var kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")

var banner = `

|\  \     |\  \|\   __  \|\  \    /  /|\  ___ \
\ \  \    \ \  \ \  \|\  \ \  \  /  / | \   __/|
 \ \  \  __\ \  \ \   __  \ \  \/  / / \ \  \_|/__
  \ \  \|\__\_\  \ \  \ \  \ \    / /   \ \  \_|\ \
   \ \____________\ \__\ \__\ \__/ /     \ \_______\
    \|____________|\|__|\|__|\|__|/       \|_______|

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

	apiHandler := api.New(executor.NewExecutor(client))

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/jobs", apiHandler.CreateJob).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	glog.Info("Starting API server...")
	glog.Info(banner)

	log.Fatal(http.ListenAndServe(":3000", router))
}
