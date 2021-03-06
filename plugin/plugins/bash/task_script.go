package bash

import (
	"context"
	fmt "fmt"

	plugin "github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

func NewScriptTask(pm *plugin.Meta) func(plugin.TaskMeta) plugin.Task {
	return func(tm plugin.TaskMeta) plugin.Task {
		tm.PluginMeta = pm
		return &Script{Meta: &tm}
	}
}

func (s *Script) Name() string {
	return fmt.Sprintf("%s.%s", s.Meta.URL, s.Meta.Name)
}

func (s *Script) MarshalHCL(l *ast.ObjectList) error {
	if o := l.Filter("evaluate"); len(o.Items) > 0 {
		err := hcl.DecodeObject(&s.EvaluateScript, o.Items[0].Val)
		if err != nil {
			return fmt.Errorf(
				"Error parsing script for %s: %s", s.Meta.Name, err)
		}
	}

	if o := l.Filter("apply"); len(o.Items) > 0 {
		err := hcl.DecodeObject(&s.ApplyScript, o.Items[0].Val)
		if err != nil {
			return fmt.Errorf(
				"Error parsing script for %s: %s", s.Meta.Name, err)
		}
	}

	return nil
}

func (s *Script) Evaluate(ctx context.Context) (*plugin.ResultChange, error) {
	return nil, nil
}

func (s *Script) Apply(ctx context.Context) (*plugin.ResultChange, error) {
	return nil, nil
}
