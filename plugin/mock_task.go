package plugin

import (
	"context"
)

type MockTask struct {
	Meta *TaskMeta `json:"meta,omitempty"`
}

func (m *MockTask) Evaluate(ctx context.Context) (*ResultChange, error) {
	return nil, nil
}

func (m *MockTask) Apply(ctx context.Context) (*ResultChange, error) {
	return nil, nil
}
