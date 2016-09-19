package preflight

import "log"

// CheckList ...
type CheckList struct {
	Packages []*Package `hcl:"package"`
	Files    []*File    `hcl:"file"`
}

// Resolve ...
func (cl *CheckList) Resolve() error {
	for _, pkg := range cl.Packages {
		log.Printf("Installing package %s ...\n", pkg.Name)
	}

	for _, file := range cl.Files {
		log.Printf("Handling file %s ...\n", file.Name)
	}

	return nil
}
