package orchestrator

import (
	glog "github.com/golang/glog"
	v1 "k8s.io/api/batch/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

type Orchestrator struct {
	clientset kubernetes.Interface
}

func NewOrchestrator(clientset kubernetes.Interface) Orchestrator {
	return Orchestrator{clientset: clientset}
}

func (o Orchestrator) NewJob(namespace string, job *v1.Job) (*v1.Job, error) {
	job, err := o.clientset.BatchV1().Jobs(namespace).Create(job)

	if err != nil {
		return nil, err
	}

	glog.Infof("Job is %v", job)

	return job, nil
}
