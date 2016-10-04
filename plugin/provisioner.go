package plugin

import (
	"net/rpc"

	"github.com/ChrisMcKenzie/preflight/config"
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
func (pc *Provisioner) Validate(t *config.Task) ([]string, []error) {
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

// Exists ...
func (pc *Provisioner) Exists(t *config.Task) (bool, error) {
	var resp ProvisionerExistsResponse
	err := pc.client.Call("Plugin.Exists", t, &resp)
	if err != nil {
		return false, err
	}

	if resp.Error != nil {
		return false, resp.Error
	}

	return resp.Exists, nil
}

// ProvisionerServer ...
type ProvisionerServer struct {
	Provisioner preflight.Provisioner
}

// Validate ...
func (ps ProvisionerServer) Validate(
	args *config.Task,
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

// Exists ...
func (ps ProvisionerServer) Exists(
	args *config.Task,
	resp *ProvisionerExistsResponse) error {
	exists, err := ps.Provisioner.Exists(args)
	berr := plugin.NewBasicError(err)
	*resp = ProvisionerExistsResponse{
		Exists: exists,
		Error:  berr,
	}

	return nil
}

// ProvisionerValidateResponse ...
type ProvisionerValidateResponse struct {
	Warnings []string
	Errors   []*plugin.BasicError
}

// ProvisionerExistsResponse ...
type ProvisionerExistsResponse struct {
	Exists bool
	Error  *plugin.BasicError
}
