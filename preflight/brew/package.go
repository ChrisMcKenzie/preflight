package brew

import (
	"os/exec"

	"github.com/ChrisMcKenzie/preflight/preflight"
)

var (
	packageManager       = "brew"
	packageInstallScript = "/usr/bin/ruby -e $(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
)

// PackageProvider ...
type PackageProvider struct{}

// Installed ...
func (p *PackageProvider) Installed(pkg *preflight.Package) error {
	if p.hasPackageManager() {
		p.installPackageManager()
	}

	return p.installPackage(pkg)
}

// IsInstalled ...
func (p *PackageProvider) IsInstalled(pkg *preflight.Package) (bool, error) {
	return false, nil
}

// Uninstalled ...
func (p *PackageProvider) Uninstalled(pkg *preflight.Package) error {
	return nil
}

func (p *PackageProvider) installPackageManager() error {
	cmd := exec.Command("bash", "-c", packageInstallScript)
	_, err := cmd.Output()
	return err
}

func (p *PackageProvider) hasPackageManager() bool {
	if _, err := exec.LookPath("brew"); err != nil {
		return false
	}

	return true
}

func (p *PackageProvider) installPackage(pkg *preflight.Package) error {
	// cmd := exec.Command("brew", "install", p.Name)
	// cmd.Stdout = os.Stdout
	// err := cmd.Run()
	// return err
	return nil
}
