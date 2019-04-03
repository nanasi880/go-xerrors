package xerrors

import "runtime"

// Builder is error builder
type Builder struct {
	wrap    error
	message string
	caller  runtime.Frame
}

// Wrap is wrapping base error.
// builder can have one error object.
func (e *Builder) Wrap(err error) *Builder {
	e.wrap = err
	return e
}

// Build is building error.
func (e *Builder) Build() error {
	return &wrapError{
		wrap:    e.wrap,
		message: e.message,
		caller:  e.caller,
	}
}
