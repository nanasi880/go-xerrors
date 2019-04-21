# xerrors
xerrorsはtesting.Tのようなスタックフレーム管理ができるエラーハンドリングパッケージです  
xerrorsではHelper関数を使うことで、様々なポイントから呼び出されるサブルーチンを避けてスタックフレームを記録することが可能です

## example

```go
package main

import (
	"fmt"

	"go.nanasi880.dev/xerrors"
)

func main() {
	err := newError()

	fmt.Printf("%v\n", err)
	// Output:
	// fail
	
	fmt.Printf("%+v\n", err)
	// Output:
	// fail
	//     mypackage/main.newError
	//         /Users/myname/project/main.go:23
}

func newError() error {
	return xerrors.New("fail").Build()
}
```

```go
package main

import (
	"fmt"

	"go.nanasi880.dev/xerrors"
)

func main() {
	err := newError()

	fmt.Printf("%v\n", err)
	// Output:
	// fail
	
	fmt.Printf("%+v\n", err)
	// Output:
	// fail
	//     mypackage/main.main
	//         /Users/myname/project/main.go:10
}

func newError() error {
	xerrors.Helper() // this function is Helper. do not logging stacktrace.
	return xerrors.New("fail").Build()
}
```
