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

	router := mux.NewRouter()

	exec := executor.NewExecutor(client)

	apiHandler := api.New(exec)

	router.HandleFunc("/api/v1/jobs", apiHandler.CreateJob).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

	http.ListenAndServe(":3000", nil)
}
