package preflight

type File struct {
	Path    string
	Source  string
	Content string

	Attrs struct {
		Owner       string
		Group       string
		Permissions int
	}
}
