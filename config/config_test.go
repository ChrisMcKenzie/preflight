package config

import (
	fmt "fmt"
	"testing"

	"github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/funkygao/golib/dag"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		config   *Config
		expected *dag.Dag
	}{
		{
			config: &Config{
				Tasks: map[string]*plugin.TaskItem{
					"mock.task.world": &plugin.TaskItem{
						Meta: &plugin.TaskMeta{
							Name:         "world",
							Dependencies: []string{"mock.task.hello"},
							URL:          "mock.task",
						},
						Task: &plugin.MockTask{
							Meta: &plugin.TaskMeta{
								Name:         "world",
								Dependencies: []string{"mock.task.hello"},
								URL:          "mock.task",
							},
						},
					},
					"mock.task.hal": &plugin.TaskItem{
						Meta: &plugin.TaskMeta{
							Name:         "hal",
							Dependencies: []string{"mock.task.hello"},
							URL:          "mock.task",
						},
						Task: &plugin.MockTask{
							Meta: &plugin.TaskMeta{
								Name:         "hal",
								Dependencies: []string{"mock.task.hello"},
								URL:          "mock.task",
							},
						},
					},
					"mock.task.hello": &plugin.TaskItem{
						Meta: &plugin.TaskMeta{
							Name:         "hello",
							URL:          "mock.task",
							Dependencies: []string{},
						},
						Task: &plugin.MockTask{
							Meta: &plugin.TaskMeta{
								Name:         "hello",
								URL:          "mock.task",
								Dependencies: []string{},
							},
						},
					},
					"mock.task.foo": &plugin.TaskItem{
						Meta: &plugin.TaskMeta{
							Name:         "foo",
							URL:          "mock.task",
							Dependencies: []string{},
						},
						Task: &plugin.MockTask{
							Meta: &plugin.TaskMeta{
								Name:         "foo",
								URL:          "mock.task",
								Dependencies: []string{},
							},
						},
					},
				},
			},
		},
	}

	for _, data := range tests {
		_, err := data.config.Graph()
		if err != nil {
			t.Fatal(err)
		}
		data.config.Traverse(func(t *plugin.TaskItem) {
			fmt.Println(t)
		})
	}
}
