package preflight

// Provisioner ...
type Provisioner interface {
	Validate(*Task) ([]string, []error)

	// Create(*Task) error
	// Delete(*Task) error
	Exists(*Task) (bool, error)
	// Read(*Task) error
}
