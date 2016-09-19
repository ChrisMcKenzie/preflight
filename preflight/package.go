package preflight

// Package ...
type Package struct {
	Action    string
	Name      string
	Provider  string
	rawConfig map[string]interface{}
}

// PackageProvider ...
type PackageProvider interface {
	IsInstalled(*Package) (bool, error)
	Installed(*Package) error
	Uninstalled(*Package) error
}
