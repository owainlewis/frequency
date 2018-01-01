package main

import (
	"flag"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/owainlewis/kcd/pkg/batch/orchestrator"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "KCD")
}

func main() {
	var kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")

	flag.Parse()

	client, err := buildClient(*kubeconfig)
	if err != nil {
		glog.Errorf("Failed to build client: %s", err)
		return
	}

	orchestrator := orchestrator.NewOrchestrator(client)

	_, err = orchestrator.CreateJob("default", "ubuntu", []string{"echo", "Hello KCD!"})
	if err != nil {
		glog.Errorf("Failed to create batch job: %s", err.Error())
	}

	http.Handle("/", http.FileServer(http.Dir("./ui/src")))
	http.ListenAndServe(":3000", nil)
}

func buildClient(cnf string) (*kubernetes.Clientset, error) {
	config, err := getKubeConfig(cnf)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getKubeConfig(cnf string) (*rest.Config, error) {
	if cnf != "" {
		return clientcmd.BuildConfigFromFlags("", cnf)
	}

	return rest.InClusterConfig()
}
