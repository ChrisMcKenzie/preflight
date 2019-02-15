package bash

import (
	fmt "fmt"

	"github.com/ChrisMcKenzie/preflight/task"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

func NewScriptTask(m task.Meta) task.Task {
	return &Script{Meta: &m}
}

func (s *Script) Name() string {
	return fmt.Sprintf("%s.%s", s.Meta.URL, s.Meta.Name)
}

func (s *Script) MarshalHCL(l *ast.ObjectList) error {
	if o := l.Filter("script"); len(o.Items) > 0 {
		err := hcl.DecodeObject(&s.Script, o.Items[0].Val)
		if err != nil {
			return fmt.Errorf(
				"Error parsing script for %s: %s", s.Meta.Name, err)
		}
	}

	return nil
}

func (s *Script) Evaluate() {

}

func (s *Script) Apply() {

}
