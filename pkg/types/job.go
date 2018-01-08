package types

// Job is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Job struct {
	Name      string   `json:"name"`
	Workspace string   `json:"workspace"`
	Image     string   `json:"image"`
	Commands  []string `json:"commands"`
	Build     *Build   `json:"build"`
}

// EnsureDefaults will set default values on a job
func (job *Job) EnsureDefaults() {
	if job.Workspace == "" {
		job.Workspace = "/workspace"
	}
}
