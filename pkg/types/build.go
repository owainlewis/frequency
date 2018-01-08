package types

// Build describes the source code VCS information (e.g. Github branch and commit SHA)
type Build struct {
	Message string `json:"message"`
	Branch  string `json:"branch"`
	Commit  string `json:"commit"`
}
