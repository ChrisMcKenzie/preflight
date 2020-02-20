package plugin

import "context"

const (
	TaskPackageIsVersion1 = true
)

type Task interface {
	Evaluate(ctx context.Context) (*ResultChange, error)
	Apply(ctx context.Context) (*ResultChange, error)
}

type ResultChange struct{}

type TaskFunc func(Meta) Task
