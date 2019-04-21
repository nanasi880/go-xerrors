// Package xerrors implements like "golang.org/x/errors" package.
//
// This package have Helper function, and if calling Helper, got a stacktrace will skipping it function.
//
// This package difference to "golang.org/x/errors" package is error wrapping style.
// Original package is using printf style ( e.g. ": %w" ) however this package using Wrap function.
//
// Simply Example:
//     package main
//
//     import (
//	       "fmt"
//
//	       "go.nanasi880.dev/xerrors"
//     )
//
//     func main() {
//	       err := newError()
//
//	       fmt.Printf("%v\n", err)
//	       // Output:
//	       // fail
//
//	       fmt.Printf("%+v\n", err)
//	       // Output:
//	       // fail
//	       //     mypackage/main.newError
//	       //         /Users/myname/project/main.go:23
//     }
//
//     func newError() error {
//	       return xerrors.New("fail").Build()
//     }
//
//
// Helper Example:
//     package main
//
//     import (
//          "fmt"
//
//           "go.nanasi880.dev/xerrors"
//      )
//
//      func main() {
//           err := newError()
//
//           fmt.Printf("%v\n", err)
//           // Output:
//           // fail
//
//           fmt.Printf("%+v\n", err)
//           // Output:
//           // fail
//           //     mypackage/main.main
//           //         /Users/myname/project/main.go:10
//      }
//
//      func newError() error {
//           xerrors.Helper() // this function is Helper. do not logging stacktrace.
//           return xerrors.New("fail").Build()
//      }
package xerrors // import "go.nanasi880.dev/xerrors"
