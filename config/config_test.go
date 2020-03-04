package config

import (
	fmt "fmt"
	"testing"

	"github.com/ChrisMcKenzie/preflight/plugin"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		config   *Config
		expected *ItemGraph
	}{
		{
			config: &Config{
				Tasks: map[string]plugin.Task{
					"mock.task.hello": &plugin.MockTask{
						Meta: &plugin.Meta{
							Dependencies: []string{""},
						},
					},
				},
			},
		},
	}

	for _, t := range tests {
		g := t.config.Graph()
		fmt.Println(g)
	}
}
