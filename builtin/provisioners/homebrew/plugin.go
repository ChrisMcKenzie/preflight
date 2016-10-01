package homebrew

import "github.com/ChrisMcKenzie/preflight/preflight"

// Greeter ...
type greeter struct{}

// Greeter ...
func Greeter() preflight.Greeter {
	return &greeter{}
}

// Greet ...
func (g *greeter) Greet() string {
	return "Hello, World"
}
