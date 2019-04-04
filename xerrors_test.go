package xerrors_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"go.nanasi880.dev/xerrors"
)

func TestNew(t *testing.T) {

	l := line()
	err := xerrors.New("test error").Build()
	errString := fmt.Sprintf("%+v", err)
	errStrings := strings.Split(errString, "\n")

	if len(errStrings) != 3 {
		t.Fatalf("invalid error messages length want:3 got:%d", len(errStrings))
	}
	if errStrings[0] != "test error:" {
		t.Fatal(errStrings)
	}
	if !strings.HasSuffix(errStrings[1], "go.nanasi880.dev/xerrors_test.TestNew") {
		t.Fatal(errStrings)
	}
	if !strings.HasSuffix(errStrings[2], fmt.Sprintf("xerrors_test.go:%d", l+1)) {
		t.Fatal(errStrings)
	}
}

func line() int {
	_, _, l, _ := runtime.Caller(1)
	return l
}
