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

	if packages := list.Filter("task"); len(packages.Items) > 0 {
		var err error
		cl.Tasks, err = loadTaskHcl(packages)
		if err != nil {
			return nil, err
		}
	}

	return cl, nil
}

func loadTaskHcl(list *ast.ObjectList) ([]*Task, error) {
	list = list.Children()
	if len(list.Items) == 0 {
		return nil, nil
	}

	var result []*Task

	for _, item := range list.Items {
		action := item.Keys[0].Token.Value().(string)
		name := item.Keys[1].Token.Value().(string)

		// var listVal *ast.ObjectList
		if _, ok := item.Val.(*ast.ObjectType); !ok {
			// 	listVal = ot.List
			// } else {
			return nil, fmt.Errorf("module '%s': should be an object", name)
		}

		var config map[string]interface{}
		if err := hcl.DecodeObject(&config, item.Val); err != nil {
			return nil, fmt.Errorf(
				"Error reading config for %s: %s",
				name,
				err)
		}

		// delete(config, "state")

		// var state string
		// if o := listVal.Filter("state"); len(o.Items) > 0 {
		// 	err := hcl.DecodeObject(&state, o.Items[0].Val)
		// 	if err != nil {
		// 		return nil, fmt.Errorf(
		// 			"Error parsing state for %s: %s",
		// 			name,
		// 			err)
		// 	}
		// }

		result = append(result, &Task{
			Name:   name,
			Type:   action,
			Config: config,
		})
	}

	return result, nil
}
