package plugin

import (
	"github.com/ChrisMcKenzie/preflight/preflight"
	"github.com/hashicorp/go-plugin"
)

// Handshake are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "PF_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "UEZfUExVR0lOX01BR0lDX0NPT0tJRQo",
}

// ProvisionerFunc ...
type ProvisionerFunc func() preflight.Provisioner

// ServeOpts ...
type ServeOpts struct {
	ProvisionerFunc ProvisionerFunc
}

// Serve serves a plugin. This function never returns and should be the final
// function called in the main function of the plugin.
func Serve(opts *ServeOpts) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins:         pluginMap(opts),
	})
}

// pluginMap returns the map[string]plugin.Plugin to use for configuring a plugin
// server or client.
func pluginMap(opts *ServeOpts) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"provisioner": &ProvisionerPlugin{opts.ProvisionerFunc},
	}
}
