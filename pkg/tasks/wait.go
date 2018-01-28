package tasks

type WaitTask struct {
	Duration int
}

func (t WaitTask) GetKind() string {
	return "WAIT"
}

func (t WaitTask) GetStatus() string {
	return ""
}
