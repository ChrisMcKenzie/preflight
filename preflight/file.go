package preflight

// File ...
type File struct {
	Name      string
	Action    string
	Path      string
	Source    string
	Content   string
	rawConfig map[string]interface{}

	Attrs struct {
		Owner       string
		Group       string
		Permissions int
	}
}
