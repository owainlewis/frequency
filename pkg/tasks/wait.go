package tasks

type Wait struct {
	DurationSeconds int
}

func (t Wait) GetKind() string {
	return "WAIT"
}

func (t Wait) GetStatus() string {
	return "PENDING"
}
