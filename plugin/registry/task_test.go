package registry

import (
	"testing"

	"github.com/ChrisMcKenzie/preflight/plugin"
)

func TestGetTask(t *testing.T) {
	meta := plugin.Meta{URL: "bash.script", Name: "test"}
	_, err := GetTask(meta)

	if err != nil {
		t.Fatal(err)
	}
}
