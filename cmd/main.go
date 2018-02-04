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

const (
	host = ":9000"
)

var kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")

var banner = `
_____
_/ ____\______   ____  ________ __   ____   ____   ____ ___.__.
\   __\\_  __ \_/ __ \/ ____/  |  \_/ __ \ /    \_/ ___<   |  |
 |  |   |  | \/\  ___< <_|  |  |  /\  ___/|   |  \  \___\___  |
 |__|   |__|    \___  >__   |____/  \___  >___|  /\___  > ____|
                    \/   |__|           \/     \/     \/\/
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

	// Tasks
	router.HandleFunc("/api/v1/tasks", API.CreateTask).Methods("POST")

	// Builds
	router.HandleFunc("/api/v1/projects/{id:[0-9]+}/builds", API.CreateBuild).Methods("POST")

	glog.Info("Starting API server...")
	glog.Info(banner)

	log.Fatal(http.ListenAndServe(host, router))
}

func buildAPI(client kubernetes.Interface) api.Api {
	ex := executor.NewDefaultExecutor(client)
	return api.New(ex)
}
