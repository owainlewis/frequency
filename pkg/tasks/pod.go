package tasks

import (
	v1 "k8s.io/api/core/v1"
)

// PodTask is a Task that runs a single Kubernetes pod exactly once
type PodTask struct {
	Image     string      `json:"image"`
	Workspace string      `json:"workspace"`
	Env       []v1.EnvVar `json:"env,omitempty"`
	Command   Command     `json:"command"`
}

// Command describes the command to run inside a Pod container
type Command struct {
	Cmd  []string `json:"cmd"`
	Args []string `json:"args"`
}

func (t PodTask) GetKind() string {
	return "PodTask"
}
