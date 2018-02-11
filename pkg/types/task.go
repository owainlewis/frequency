package types

import (
	v1 "k8s.io/api/core/v1"
)

// Task runs a single Kubernetes pod exactly once
type Task struct {
	Name      string      `json:"name"`
	Image     string      `json:"image"`
	Workspace string      `json:"workspace"`
	Env       []v1.EnvVar `json:"env"`
	Run       run         `json:"run"`
	Checkout  *Checkout   `json:"checkout"`
}

func (t *Task) Validate() []error {
	var errs []error
	return errs
}

// Run describes the command to run inside a Pod container
type run struct {
	Command []string `json:"command"`
	Args    []string `json:"args"`
}

// Checkout describes the source code VCS information (e.g. Github branch and commit SHA)
type Checkout struct {
	URL  string   `json:"url"`
	Post []string `json:"post"`
}

// If set this will checkout the source code into a different working directory
// Destination string `json:"destination"`
// Branch      string `json:"branch"`
// Commit      string `json:"commit"`

// SetDefaults ensures that sensible default values are applied to a task
func (t Task) SetDefaults() {
	if t.Workspace == "" {
		t.Workspace = "/"
	}
}
