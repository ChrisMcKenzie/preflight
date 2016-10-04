package config

// Task ...
type Task struct {
	Type  string
	Name  string
	State string

	RawConfig map[string]interface{}
}
