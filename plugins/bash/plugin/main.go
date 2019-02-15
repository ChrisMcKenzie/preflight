package main

import (
	"github.com/ChrisMcKenzie/preflight/plugins/bash"
	"github.com/ChrisMcKenzie/preflight/task"
)

var (
	Version string
)

const _ = task.TaskPackageIsVersion1

func RegisterTasks() map[string]task.TaskFunc {
	return map[string]task.TaskFunc{
		"bash.script": bash.NewScriptTask,
	}
}
