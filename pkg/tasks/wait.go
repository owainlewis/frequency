package tasks

type WaitTask struct {
	Duration int
}

func (t WaitTask) GetKind() string {
	return "WAIT"
}
