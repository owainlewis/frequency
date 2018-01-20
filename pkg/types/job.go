package types

// Command describes the command to run inside a Pod container
type Command struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

// Job is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Job struct {
	// ID is the UID of the pod created for this job and gets assigned
	// after a job has been submitted to Kubernetes
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Workspace   string            `json:"workspace"`
	Environment map[string]string `json:"environment"`
	Command     Command           `json:"command"`
	Source      *Source           `json:"source"`
}

// EnsureDefaults will set default values on a job
func (job *Job) EnsureDefaults() {
	if job.Workspace == "" {
		job.Workspace = "/workspace"
	}
}
