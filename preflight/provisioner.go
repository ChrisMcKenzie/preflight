package preflight

import "github.com/ChrisMcKenzie/preflight/config"

// Provisioner ...
type Provisioner interface {
	Validate(*config.Task) ([]string, []error)
	// Create(*Task) error
	// Delete(*Task) error
	Exists(*config.Task) (bool, error)
	// Read(*Task) error
}
