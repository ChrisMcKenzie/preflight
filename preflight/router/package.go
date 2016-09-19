package router

import (
	"fmt"

	"github.com/ChrisMcKenzie/preflight/preflight/brew"
)

// PackageProviderRouter defines a sort of proxy provider that will read the
// provider set in the Package and execute the interface Method on that provider
type PackageProviderRouter struct{}

func (p *PackageProviderRouter) getProvider(provider string) (PackageProvider, error) {
	switch provider {
	case "brew":
		return brew.PackageProvider{}, nil
	default:
		return nil, fmt.Errorf("Provider %s not found...", provider)
	}
}

// Installed ...
func (p *PackageProviderRouter) Installed(pkg *Package) error {
	provider, err := p.getProvider(pkg.Provider)
	if err != nil {
		return err
	}

	return provider.Installed(pkg)
}

// IsInstalled ...
func (p *PackageProviderRouter) IsInstalled(pkg *Package) (bool, error) {
	provider, err := p.getProvider(pkg.Provider)
	if err != nil {
		return err
	}

	return provider.Installed(pkg)
}

// Uninstalled ...
func (p *PackageProviderRouter) Uninstalled(pkg *Package) error {
	provider, err := p.getProvider(pkg.Provider)
	if err != nil {
		return err
	}

	return provider.Uninstalled(pkg)
}
