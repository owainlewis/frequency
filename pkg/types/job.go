package types

// Job is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Job struct {
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Workspace   string            `json:"workspace"`
	Environment map[string]string `json:"environment"`
	Steps       []string          `json:"steps"`
}

// EnsureDefaults will set default values on a job
func (job *Job) EnsureDefaults() {
	if job.Workspace == "" {
		job.Workspace = "/workspace"
	}
}
