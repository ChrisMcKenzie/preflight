package preflight

// Provisioner ...
type Provisioner interface {
	Validate(*Task) []string
}
