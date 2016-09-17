package preflight

import (
	"os"
	"os/exec"
)

var (
	packageManager       = "brew"
	packageInstallScript = "/usr/bin/ruby -e $(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
)

type Package struct {
	Name   string
	Update bool
}

func (p *Package) Install() error {
	if p.hasPackageManager() {
		p.installPackageManager()
	}

	return p.installPackage()
}

func (p *Package) installPackageManager() error {
	cmd := exec.Command("bash", "-c", packageInstallScript)
	_, err := cmd.Output()
	return err
}

func (p *Package) hasPackageManager() bool {
	if _, err := exec.LookPath("brew"); err != nil {
		return false
	}

	return true
}

func (p *Package) installPackage() error {
	cmd := exec.Command("brew", "install", p.Name)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
