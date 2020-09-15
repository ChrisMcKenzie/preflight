package main

import (
	"github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/ChrisMcKenzie/preflight/plugin/plugins/bash"
)

var (
	// nolint: deadcode, unused
	Version string
)

const _ = plugin.TaskPackageIsVersion1

// nolint: deadcode, unused
func RegisterTasks() map[string]plugin.TaskFunc {
	m := &plugin.Meta{
		Version: Version,
	}

	return map[string]plugin.TaskFunc{
		"bash.script": bash.NewScriptTask(m),
	}
}
