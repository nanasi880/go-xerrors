package xerrors

import (
	"fmt"

	"golang.org/x/xerrors"
)

// https://godoc.org/golang.org/x/xerrors#Unwrap
func Unwrap(err error) error {
	return xerrors.Unwrap(err)
}

// https://godoc.org/golang.org/x/xerrors#Is
func Is(err error, target error) bool {
	return xerrors.Is(err, target)
}

// https://godoc.org/golang.org/x/xerrors#As
func As(err error, target interface{}) bool {
	return xerrors.As(err, target)
}

// https://godoc.org/golang.org/x/xerrors#Opaque
func Opaque(err error) error {
	return xerrors.Opaque(err)
}

// https://godoc.org/golang.org/x/xerrors#FormatError
func FormatError(f xerrors.Formatter, s fmt.State, verb rune) {
	xerrors.FormatError(f, s, verb)
}
