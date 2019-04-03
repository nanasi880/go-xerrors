package xerrors

import (
	"runtime"
	"sync"
)

var (
	helpersMutex sync.RWMutex
	helpers      = make(map[string]struct{})
)

// Helper is marking the caller function as helper function.
// New and Errorf function is skipping the helper function.
func Helper() {
	frame := callerFrame(1)
	if frame.PC == 0 {
		return
	}

	helpersMutex.Lock()
	defer helpersMutex.Unlock()

	helpers[frame.Function] = struct{}{}
}

func callerFrame(skip int) runtime.Frame {

	var (
		pc         [50]uintptr
		firstFrame runtime.Frame
	)

	n := runtime.Callers(skip+2, pc[:]) // skip + runtime.Callers + callerFrame
	if n == 0 {
		return firstFrame
	}

	frames := runtime.CallersFrames(pc[:])

	helpersMutex.RLock()
	defer helpersMutex.RUnlock()

	for i := 0; i < n; i++ {

		f, more := frames.Next()
		if !more {
			return firstFrame
		}

		if firstFrame.PC == 0 {
			firstFrame = f
		}

		if _, ok := helpers[f.Function]; ok {
			continue
		}

		return f
	}

	return firstFrame
}
