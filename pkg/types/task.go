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
	Run       run         `json:"run"`
	Source    *Source     `json:"source"`
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

// Source describes the source code VCS information (e.g. Github branch and commit SHA)
type Source struct {
	Domain     string `json:"domain"`
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
	Commit     string `json:"commit"`
}

func (s *Source) GetPublicCloneURL() string {
	return fmt.Sprintf("https://%s/%s/%s.git", s.Domain, s.Owner, s.Repository)
}

// SetDefaults ensures that sensible default values are applied to a task
func (t Task) SetDefaults() {
	if t.Workspace == "" {
		t.Workspace = "/"
	}
}
