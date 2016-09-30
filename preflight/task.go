package preflight

// Task ...
type Task struct {
	Type   string
	Name   string
	State  string
	Config map[string]interface{}
}
