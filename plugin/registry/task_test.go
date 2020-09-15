package registry

import (
	"os"
	"testing"

	"github.com/ChrisMcKenzie/preflight/plugin"
)

func TestGetTask(t *testing.T) {
	defer func() {
		os.RemoveAll("./.preflight")
	}()

	meta := plugin.TaskMeta{URL: "bash.script", Name: "test"}
	_, err := GetTask(meta)

	if err != nil {
		t.Fatal(err)
	}
}
