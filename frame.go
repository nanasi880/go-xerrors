package xerrors

import (
	"runtime"

	"golang.org/x/xerrors"
)

func printFrame(f runtime.Frame, p xerrors.Printer) {

	if p.Detail() {

		var (
			file     = f.File
			line     = f.Line
			function = f.Function
		)
		if file == "" {
			file = "???"
		}
		if function == "" {
			function = "???"
		}

		p.Printf("%s\n", function)
		p.Printf("    %s:%d\n", file, line)
	}

}
