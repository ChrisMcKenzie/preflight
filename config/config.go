package config

import (
	fmt "fmt"
	strings "strings"

	"github.com/ChrisMcKenzie/preflight/plugin"
)

type Config struct {
	Tasks     []plugin.Task
	Variables []Variable
}

func (c *Config) GoString() string {
	str := "Config {\n  Tasks: {\n"
	for _, task := range c.Tasks {
		result := fmt.Sprintf("%+v", task)
		s, ok := task.(fmt.GoStringer)
		if ok {
			result = s.GoString()
		}
		str += strings.Replace(result, "\n", "    \n", -1) + "\n"
	}

	str += "  }\n}"

	return str
}
