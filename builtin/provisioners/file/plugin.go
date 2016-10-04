package file

import (
	"github.com/ChrisMcKenzie/preflight/config"
	"github.com/ChrisMcKenzie/preflight/preflight"
)

type provisioner struct{}

// Provisioner ...
func Provisioner() preflight.Provisioner {
	return &provisioner{}
}

// Validate ...
func (*provisioner) Validate(t *config.Task) ([]string, []error) {
	return []string{}, []error{}
}

// Exists ...
func (*provisioner) Exists(t *config.Task) (bool, error) {

	return false, nil
}
