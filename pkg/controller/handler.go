package controller

import (
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
)

// podUpdateHandler is called whenever a pod state update is seen.
// This will have information about the state of the pod lifecycle which can be
// pushed back to the upstream API server
func (c *Controller) podUpdateHandler(pod *v1.Pod) {

	podCopy := pod.DeepCopy()

	pid := string(podCopy.GetUID())
	glog.Infof("Pod %s %s is in phase %s", pid, pod.GetName(), pod.Status.Phase)

	if c.store == nil {
		glog.Warning("Datastore is unbound. Refusing to progress")
		return
	}

	job, err := c.store.FindJobByID(pid)
	if err != nil {
		glog.Errorf("Failed to find job in data store: %s", err)
		return
	}

	glog.Infof("Found matching job %s. Attempting to update status", job.ID)

	//c.store.UpdateJob(job)

	// Get the job where ID == pid and update the status
	// Update the status of the job in the API
	// LOGS

}
