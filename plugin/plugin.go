package plugin

import "github.com/hashicorp/go-plugin"

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"greeter": new(GreeterPlugin),
}
