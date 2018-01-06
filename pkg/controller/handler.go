package controller

import (
	"strings"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
)

// podUpdateHandler is called whenever a pod state update is seen.
// This will have information about the state of the pod lifecycle which can be
// pushed back to the upstream API server
func (c *Controller) podUpdateHandler(pod *v1.Pod) {
	// TODO (remove me)
	if strings.HasPrefix(pod.Name, "kcd-") {
		glog.Infof("Pod %s is in phase %s", pod.GetName(), pod.Status.Phase)

		// Update the status of the job in the API

		// LOGS

	}
}
