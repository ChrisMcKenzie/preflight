package main

import (
	"github.com/ChrisMcKenzie/preflight/builtin/provisioners/homebrew"
	"github.com/ChrisMcKenzie/preflight/plugin"
)

func main() {
	// We're a plugin! Serve the plugin. We set the handshake config
	// so that the host and our plugin can verify they can talk to each other.
	// Then we set the plugin map to say what plugins we're serving.
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: homebrew.Provisioner,
	})
}
