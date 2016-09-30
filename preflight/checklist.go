package preflight

import "log"

// CheckList ...
type CheckList struct {
	Tasks []*Task `hcl:"package"`
}

// Resolve ...
func (cl *CheckList) Resolve() error {
	done := make(chan bool)
	if len(cl.Tasks) > 0 {
		// go cl.resolvePackages(done)
	}

	<-done
	return nil
}

func (cl *CheckList) resolvePackages(done chan bool) {
	for _, pkg := range cl.Tasks {
		log.Printf("%s.%s ...\n", pkg.Name, pkg.Type)
		switch pkg.Type {
		case "installed":
			// router.Installed(pkg)
		case "uninstalled":
			// router.Uninstalled(pkg)
		default:
			log.Printf("Invalid state %s for %s...\n", pkg.Type, pkg.Name)
		}
	}

	done <- true
}
