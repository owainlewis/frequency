package types

// Source describes the source code VCS information (e.g. Github branch and commit SHA)
type Source struct {
	GitURL    string `json:"git_url"`
	GitBranch string `json:"git_branch"`
	// This is used to mount a deploy SSH key so that
	// we can git clone private repos
	DeployKeySecret string `json:"deploy_key_secret"`
}
