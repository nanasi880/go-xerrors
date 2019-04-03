package xerrors

import (
	"fmt"
)

// New is create error builder with plain test message.
func New(s string) *Builder {
	return &Builder{
		message: s,
		caller:  callerFrame(1),
	}
}

// Errorf is create error builder with format string.
func Errorf(format string, args ...interface{}) *Builder {
	return &Builder{
		message: fmt.Sprintf(format, args...),
		caller:  callerFrame(1),
	}
}
