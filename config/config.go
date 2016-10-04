package config

// Config ...
type Config struct {
	Tasks []*Task `hcl:"package"`
}
