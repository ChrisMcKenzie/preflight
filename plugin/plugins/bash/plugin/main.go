package main

import (
	"github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/ChrisMcKenzie/preflight/plugin/plugins/bash"
)

var (
	Version string
)

const _ = plugin.TaskPackageIsVersion1

func RegisterTasks() map[string]plugin.TaskFunc {
	return map[string]plugin.TaskFunc{
		"bash.script": bash.NewScriptTask,
	}
}
