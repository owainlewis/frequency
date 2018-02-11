package types

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

// Task runs a single Kubernetes pod exactly once
type Task struct {
	Name      string      `json:"name"`
	Image     string      `json:"image"`
	Workspace string      `json:"workspace"`
	Env       []v1.EnvVar `json:"env"`
	Checkout  *Checkout   `json:"checkout"`
	Run       run         `json:"run"`
	// A simplified implementation of the run command that gets post compiled
	Steps []string `json:"steps"`
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

// Validate a task is suitable to submit to the API server
func (t Task) Validate() []error {
	var errs []error

	if t.Image == "" {
		errs = append(errs, fmt.Errorf("Missing image"))
	}

	if (len(t.Run.Command) != 0 || len(t.Run.Args) != 0) && len(t.Steps) != 0 {
		errs = append(errs, fmt.Errorf("Cannot declare both Run and Steps for a task"))
	}

	return errs
}
