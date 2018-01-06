package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	builder "github.com/owainlewis/kcd/pkg/client"
	"github.com/owainlewis/kcd/pkg/controller"
	"github.com/owainlewis/kcd/pkg/orchestrator"
	"github.com/owainlewis/kcd/pkg/types"

	"k8s.io/client-go/kubernetes"
)

var kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")

func main() {

	flag.Parse()

	client, err := builder.Build(*kubeconfig)
	if err != nil {
		glog.Errorf("Failed to build client: %s", err)
		return
	}

	startPodController(client)

	orch := orchestrator.NewOrchestrator(client)

	stage := types.Stage{
		Name:  "hello-kcd",
		Image: "ubuntu:latest",
		Commands: []string{
			"echo \"Hello World\"",
		},
	}

	err = orch.ExecuteStage("default", stage)

	if err != nil {
		glog.Errorf("Orchestration failed: %s", err.Error())
	}

	http.Handle("/", http.FileServer(http.Dir("./ui/src")))
	http.ListenAndServe(":3000", nil)
}

func startPodController(client *kubernetes.Clientset) {
	ctrl := controller.NewController(client)
	stop := make(chan struct{})

	defer close(stop)
	go ctrl.Run(1, stop)
}
