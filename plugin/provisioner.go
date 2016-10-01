package plugin

import (
	"net/rpc"

	"github.com/ChrisMcKenzie/preflight/preflight"
	"github.com/hashicorp/go-plugin"
)

// ProvisionerPlugin ...
type ProvisionerPlugin struct {
	F func() preflight.Provisioner
}

// Server ...
func (p ProvisionerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ProvisionerServer{p.F()}, nil
}

// Client ...
func (ProvisionerPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &Provisioner{client: c}, nil
}

// Provisioner ...
type Provisioner struct{ client *rpc.Client }

// Validate ...
func (pc *Provisioner) Validate(t *preflight.Task) []string {
	var resp ProvisionerValidateResponse
	err := pc.client.Call("Plugin.Validate", t, &resp)
	if err != nil {
		return []string{err.Error()}
	}

	return resp.Errors
}

// ProvisionerServer ...
type ProvisionerServer struct {
	Provisioner preflight.Provisioner
}

// Validate ...
func (ps ProvisionerServer) Validate(
	args *preflight.Task,
	resp *ProvisionerValidateResponse) error {
	errs := ps.Provisioner.Validate(args)
	*resp = ProvisionerValidateResponse{
		Errors: errs,
	}
	return nil
}

// ProvisionerValidateResponse ...
type ProvisionerValidateResponse struct {
	Errors []string
}
