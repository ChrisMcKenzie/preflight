package plugin

import (
	"context"
	fmt "fmt"
)

const (
	TaskPackageIsVersion1 = true
)

type TaskItem struct {
	Id   string
	Task Task
	Meta *Meta
}

func (ti *TaskItem) String() string {
	return fmt.Sprintf("%s.%s", ti.Meta.URL, ti.Meta.Name)
}

type Task interface {
	Evaluate(ctx context.Context) (*ResultChange, error)
	Apply(ctx context.Context) (*ResultChange, error)
}

type ResultChange struct{}

type TaskFunc func(Meta) Task
