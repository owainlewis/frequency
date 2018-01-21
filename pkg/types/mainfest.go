package types

// Manifest is a type that defines that top level
// user facing contract for defining builds
type Manifest struct {
	Version int32          `json:"version"`
	Jobs    map[string]Job `json:"jobs"`
}

func (m *Manifest) EnsureDefaults() *Manifest {
	if m.Version == 0 {
		m.Version = 1
	}

	return m
}
