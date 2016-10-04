package plugin

import (
	"encoding/gob"

	"github.com/hashicorp/go-plugin"
)

func init() {
	gob.Register(make([]map[string]interface{}, 0))
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"provisioner": new(ProvisionerPlugin),
}
