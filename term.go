// Package term provides high level terminal i/o.
package term

import (
	"io"
	"os"
)

// Term provides high level terminal i/o.
type Term interface {
	// Args returns all command line arguments, excluding the program name itself.
	Args() []string
	// Stdout returns the stdout io.Writer.
	Stdout() io.Writer
	// Stderr returns the stderr io.Writer.
	Stderr() io.Writer
}

// DefaultTerm is the default terminal to use.
var DefaultTerm Term = &term{
	args:   os.Args[1:],
	stdout: os.Stdout,
	stderr: os.Stderr,
}

// term implements the Term interface.
type term struct {
	args   []string
	stdout io.Writer
	stderr io.Writer
}

// Args is part of the Term interface.
func (t *term) Args() []string {
	return t.args
}

// Stdout is part of the Term interface.
func (t *term) Stdout() io.Writer {
	return t.stdout
}

// Stderr is part of the Term interface.
func (t *term) Stderr() io.Writer {
	return t.stderr
}
