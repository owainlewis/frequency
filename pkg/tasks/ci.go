package tasks

// Source describes the source code VCS information (e.g. Github branch and commit SHA)
type Source struct {
	GitURL    string `json:"git_url"`
	GitBranch string `json:"git_branch"`
}

type CITask struct {
	PodTask

	Source *Source
}
