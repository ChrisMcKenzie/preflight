package config

import (
	fmt "fmt"
	strings "strings"

	"github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/ChrisMcKenzie/preflight/plugin/registry"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

type Marshaller interface {
	MarshalHCL(*ast.ObjectList) error
}

type Variable struct {
	Name    string
	Type    string
	Default interface{}
}

func parse(filename string, f *ast.File) (*Config, error) {
	// Top-level item should be the object list
	list, ok := f.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: file does not contain root node object")
	}

	cfg := new(Config)

	if tasks := list.Filter("task"); len(tasks.Items) > 0 {
		var err error
		cfg.Tasks, err = loadTasks(tasks)
		if err != nil {
			return nil, err
		}
	}

	if vars := list.Filter("var"); len(vars.Items) > 0 {
		var err error
		cfg.Variables, err = loadVariables(vars)
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func loadTasks(list *ast.ObjectList) ([]plugin.Task, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []plugin.Task

	for _, item := range list.Items {
		taskURL := strings.Replace(item.Keys[0].Token.Value().(string), "_", ".", -1)
		name := item.Keys[1].Token.Value().(string)

		meta := plugin.Meta{URL: taskURL, Name: name}

		ot, ok := item.Val.(*ast.ObjectType)
		if !ok {
			return nil, fmt.Errorf("task '%s.%s': should be an object", taskURL, name)
		}

		if o := ot.List.Filter("url"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&meta.URL, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing url for %s: %s", meta.Name, err)
			}
		}

		if o := ot.List.Filter("version"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&meta.Version, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing version for %s: %s", meta.Name, err)
			}
		}

		t, err := registry.GetTask(meta)
		if err != nil {
			return nil, fmt.Errorf("unable to get task named %s: %s", taskURL, err)
		}

		if m, ok := t.(Marshaller); ok {
			if err := m.MarshalHCL(ot.List); err != nil {
				return nil, err
			}
		}
		result = append(result, t)
	}

	return result, nil
}

func loadVariables(list *ast.ObjectList) ([]Variable, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []Variable
	for _, item := range list.Items {
		name := item.Keys[0].Token.Value().(string)

		ot, ok := item.Val.(*ast.ObjectType)
		if !ok {
			return nil, fmt.Errorf("variable '%s': should be an object", name)
		}

		var typ string
		if o := ot.List.Filter("type"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&typ, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing type for %s: %s", name, err)
			}
		}

		var def interface{}
		if o := ot.List.Filter("default"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&def, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing type for %s: %s", name, err)
			}
		}

		if typ == "" {
			typ = fmt.Sprintf("%T", def)
		}

		result = append(result, Variable{
			Name:    name,
			Type:    typ,
			Default: def,
		})
	}

	return result, nil
}
