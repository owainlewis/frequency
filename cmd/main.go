package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	builder "github.com/owainlewis/kcd/pkg/client"
	"github.com/owainlewis/kcd/pkg/controller"
	"github.com/owainlewis/kcd/pkg/types"
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

	exec := executor.NewExecutor(client)

	job := types.Job{
		Name:  "hello-kcd",
		Image: "golang",
		Commands: []string{
			"go build -v main.go",
			"mv ./main $OUTPUT_DIR",
			"ls $OUTPUT_DIR",
		},
	}

	err = exec.Execute("default", job)
	if err != nil {
		glog.Errorf("Execution failed: %s", err.Error())
	}

	http.Handle("/", http.FileServer(http.Dir("./ui/src")))
	http.ListenAndServe(":3000", nil)
}
