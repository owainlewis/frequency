package types

// Source describes the source code VCS information (e.g. Github branch and commit SHA)
type Source struct {
	Branch  string `json:"branch"`
	Commit  string `json:"commit"`
	Message string `json:"message"`
}

// Job is the smallest unit of execution in KCD.
// It represents the execution of a Kubernetes pod with a series of
// commands specified by the user as part of the kcd.yml manifest file.
type Job struct {
	Name      string   `json:"name"`
	Workspace string   `json:"workspace"`
	Image     string   `json:"image"`
	Commands  []string `json:"commands"`
	Source    *Source  `json:"source"`
}

// EnsureDefaults will set default values on a job
func (job *Job) EnsureDefaults() {
	if job.Workspace == "" {
		job.Workspace = "/workspace"
	}
}
