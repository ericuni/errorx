package errs

import (
	"fmt"
	"runtime"
)

func getLocation(level int) string {
	// 0 for the current function, 1 for the caller, and 2 for the caller of the caller
	pc, _, line, _ := runtime.Caller(2)
	fun := runtime.FuncForPC(pc)
	return fmt.Sprintf("[%v:%v]", fun.Name(), line)
}
