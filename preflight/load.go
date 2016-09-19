package preflight

import (
	"fmt"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

// LoadHcl ...
func LoadHcl(f *ast.File) (*CheckList, error) {
	// Top-level item should be the object list
	list, ok := f.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: file does not contain root node object")
	}

	cl := new(CheckList)

	if packages := list.Filter("package"); len(packages.Items) > 0 {
		var err error
		cl.Packages, err = loadPackageHcl(packages)
		if err != nil {
			return nil, err
		}
	}

	if packages := list.Filter("file"); len(packages.Items) > 0 {
		var err error
		cl.Files, err = loadFileHcl(packages)
		if err != nil {
			return nil, err
		}
	}

	return cl, nil
}

func loadFileHcl(list *ast.ObjectList) ([]*File, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []*File

	for _, item := range list.Items {
		action := item.Keys[0].Token.Value().(string)
		name := item.Keys[1].Token.Value().(string)

		var listVal *ast.ObjectList
		if ot, ok := item.Val.(*ast.ObjectType); ok {
			listVal = ot.List
		} else {
			return nil, fmt.Errorf("module '%s': should be an object", name)
		}

		var config map[string]interface{}
		if err := hcl.DecodeObject(&config, item.Val); err != nil {
			return nil, fmt.Errorf(
				"Error reading config for %s: %s",
				name,
				err)
		}

		delete(config, "path")
		delete(config, "source")
		delete(config, "attrs")

		var path string
		if o := listVal.Filter("path"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&path, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing source for %s: %s",
					name,
					err)
			}
		}

		var source string
		if o := listVal.Filter("source"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&source, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing source for %s: %s",
					name,
					err)
			}
		}

		var attrs struct {
			Owner       string
			Group       string
			Permissions int
		}
		if o := listVal.Filter("attrs"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&attrs, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing source for %s: %s",
					name,
					err)
			}
		}

		result = append(result, &File{
			Name:      name,
			Action:    action,
			Path:      path,
			Source:    source,
			Attrs:     attrs,
			rawConfig: config,
		})
	}

	return result, nil
}

func loadPackageHcl(list *ast.ObjectList) ([]*Package, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []*Package

	for _, item := range list.Items {
		action := item.Keys[0].Token.Value().(string)
		name := item.Keys[1].Token.Value().(string)

		var listVal *ast.ObjectList
		if ot, ok := item.Val.(*ast.ObjectType); ok {
			listVal = ot.List
		} else {
			return nil, fmt.Errorf("module '%s': should be an object", name)
		}

		var config map[string]interface{}
		if err := hcl.DecodeObject(&config, item.Val); err != nil {
			return nil, fmt.Errorf(
				"Error reading config for %s: %s",
				name,
				err)
		}

		delete(config, "provider")

		var provider string
		if o := listVal.Filter("provider"); len(o.Items) > 0 {
			err := hcl.DecodeObject(&provider, o.Items[0].Val)
			if err != nil {
				return nil, fmt.Errorf(
					"Error parsing source for %s: %s",
					name,
					err)
			}
		}

		result = append(result, &Package{
			Name:      name,
			Action:    action,
			Provider:  provider,
			rawConfig: config,
		})
	}

	return result, nil
}
