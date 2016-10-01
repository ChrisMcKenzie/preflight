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
func (pc *Provisioner) Validate(t *preflight.Task) ([]string, []error) {
	var resp ProvisionerValidateResponse
	err := pc.client.Call("Plugin.Validate", t, &resp)
	if err != nil {
		return []string{}, []error{err}
	}

	var errs []error
	if len(resp.Errors) > 0 {
		errs = make([]error, len(resp.Errors))
		for i, err := range resp.Errors {
			errs[i] = err
		}
	}

	return resp.Warnings, errs
}

// ProvisionerServer ...
type ProvisionerServer struct {
	Provisioner preflight.Provisioner
}

// Validate ...
func (ps ProvisionerServer) Validate(
	args *preflight.Task,
	resp *ProvisionerValidateResponse) error {
	wars, errs := ps.Provisioner.Validate(args)
	berrs := make([]*plugin.BasicError, len(errs))
	for i, err := range errs {
		berrs[i] = plugin.NewBasicError(err)
	}
	*resp = ProvisionerValidateResponse{
		Warnings: wars,
		Errors:   berrs,
	}
	return nil
}

// ProvisionerValidateResponse ...
type ProvisionerValidateResponse struct {
	Warnings []string
	Errors   []*plugin.BasicError
}
