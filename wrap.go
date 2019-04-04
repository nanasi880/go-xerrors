package xerrors

import (
	"fmt"
	"runtime"

	"golang.org/x/xerrors"
)

type wrapError struct {
	wrap    error
	message string
	caller  runtime.Frame
}

// Error is implements of error
func (e *wrapError) Error() string {
	return fmt.Sprint(e)
}

// Format is implements of fmt.Formatter
func (e *wrapError) Format(s fmt.State, verb rune) {
	FormatError(e, s, verb)
}

// Format is implements of xerrors.Formatter
func (e *wrapError) FormatError(p xerrors.Printer) error {
	p.Print(e.message)
	printFrame(e.caller, p)
	return e.wrap
}

// Unwrap is implements of xerrors.Wrapper
func (e *wrapError) Unwrap() error {
	return e.wrap
}
