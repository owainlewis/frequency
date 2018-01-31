package tasks

import (
	v1 "k8s.io/api/core/v1"
)

// Task runs a single Kubernetes pod exactly once
type Task struct {
	Image     string      `json:"image"`
	Workspace string      `json:"workspace"`
	Env       []v1.EnvVar `json:"env,omitempty"`
	Run       run         `json:"run"`
	Source    *Source     `json:"source"`
}

// Run describes the command to run inside a Pod container
type run struct {
	Command []string `json:"command"`
	Args    []string `json:"args"`
}

// Source describes the source code VCS information (e.g. Github branch and commit SHA)
type Source struct {
	GitURL    string `json:"git_url"`
	GitBranch string `json:"git_branch"`
}

func (t Task) SetDefaults() {
	if t.Workspace == "" {
		t.Workspace = "/"
	}
}
