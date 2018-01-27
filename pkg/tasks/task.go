package tasks

type Task interface {
	GetKind() string
	GetStatus() string
}
