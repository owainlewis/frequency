package types

type Step struct {
	Name    string
	Command string
}

// Stage is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Stage struct {
	Name     string
	Image    string
	Commands []string
}

// Pipeline is a series of stages to be run in some order.
// Stages can be run in parallel.
type Pipeline struct {
	Name   string
	Stages []Stage
}
