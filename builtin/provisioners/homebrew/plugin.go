package homebrew

import (
	"errors"

	"github.com/ChrisMcKenzie/preflight/preflight"
)

// provisioner ...
type provisioner struct{}

// Provisioner ...
func Provisioner() preflight.Provisioner {
	return &provisioner{}
}

// Validate ...
func (*provisioner) Validate(t *preflight.Task) ([]string, []error) {
	err := errors.New("emit macho dwarf: elf header corrupted")
	return []string{}, []error{err}
}
