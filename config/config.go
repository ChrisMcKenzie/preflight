package config

import (
	fmt "fmt"
	strings "strings"

	"github.com/ChrisMcKenzie/preflight/plugin"
)

type Config struct {
	Tasks     map[string]plugin.Task
	Variables []Variable
}

func (c *Config) GoString() string {
	str := "Config {\n  Tasks: {\n"
	for k, task := range c.Tasks {
		result := fmt.Sprintf("%s: %+v", k, task)
		s, ok := task.(fmt.GoStringer)
		if ok {
			result = s.GoString()
		}
		str += strings.Replace(result, "\n", "    \n", -1) + "\n"
	}

	str += "  }\n}"

	return str
}

func (c *Config) Graph() *ItemGraph {
	return c.renderGraph()
}

func (c *Config) renderGraph() *ItemGraph {
	return nil
}
