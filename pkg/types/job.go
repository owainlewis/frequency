package types

// Job is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Job struct {
	Name      string
	Workspace string
	Image     string
	Commands  []string
}

func (job *Job) EnsureDefaults() {
	if job.Workspace == "" {
		job.Workspace = "/workspace"
	}
}
