package preflight

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

// CheckList ...
type CheckList struct {
	Tasks []*Task `hcl:"package"`
}

// Plan ...
func (cl *CheckList) Plan() {
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	for _, task := range cl.Tasks {
		b, _ := json.MarshalIndent(task.Config, "", "  ")
		fmt.Printf("===== TASK: (%s) %s =====\n\n%s\n\n", green(task.Type), task.Name, yellow(string(b)))
	}
}
