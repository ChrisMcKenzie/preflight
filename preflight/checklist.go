package preflight

import "log"

type CheckList struct {
	Packages []*Package `hcl:"package"`
	Files    []*File    `hcl:"file"`
}

func (cl *CheckList) Resolve() error {
	for _, pkg := range cl.Packages {
		log.Printf("Installing package %s ...", pkg.Name)
		if err := pkg.Install(); err != nil {
			log.Println(err)
		}
	}

	return nil
}
