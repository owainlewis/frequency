package types

// Stage is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Stage struct {
	Name     string
	Image    string
	Commands []string
}

// StageResult contains information about the execution of a stage.
// TODO StagePassed, StageFailed, StageRunning ADT
type StageResult struct {
}

// Pipeline is a series of stages to be run in some order.
// Stages can be run in parallel.
type Pipeline struct {
}
