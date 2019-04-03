// Package xerrors implements like "golang.org/x/errors" package.
//
// This package have Helper function, and if calling Helper, got a stacktrace will skipping it function.
//
// This package difference to "golang.org/x/errors" package is error wrapping style.
// Original package is using printf style ( e.g. ": %w" ) however this package using Wrap function.
//
// Example:
//     newErr := xerrors.New("Oops").Wrap(baseErr).Build()
package xerrors // import "go.nanasi880.dev/xerrors"
