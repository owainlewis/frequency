package types

import (
	v1 "k8s.io/api/core/v1"
)

// Command describes the command to run inside a Pod container
type Command struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

// Job is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the frequency.yml manifest file.
type Job struct {
	// ID is the UID of the pod created for this job and gets assigned
	// after a job has been submitted to Kubernetes
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Image     string      `json:"image"`
	Workspace string      `json:"workspace"`
	Env       []v1.EnvVar `json:"env,omitempty"`
	Command   Command     `json:"command"`
	Source    *Source     `json:"source,omitempty"`
}

// EnsureDefaults will set default values on a job
func (job *Job) EnsureDefaults() {
	if job.Workspace == "" {
		job.Workspace = "/workspace"
	}
}
