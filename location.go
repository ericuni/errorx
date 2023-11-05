package errorx

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func getLocation() string {
	// 0 for the current function, 1 for the caller, and 2 for the caller of the caller
	// returned file is absolute path, e.g.: /home/ericuni/git/errorx/error_test.go
	pc, file, line, _ := runtime.Caller(2)
	fun := runtime.FuncForPC(pc)

	fileName := path.Base(file)

	// fun.Name() is like
	// github.com/ericuni/errorx.TestNew
	// github.com/ericuni/errorx.TestTrace.func2.1 if has goroutine or subtest
	fullFunName := fun.Name()

	parts := strings.Split(path.Base(fullFunName), ".")
	if len(parts) < 2 {
		return fmt.Sprintf("[%s.%s:%d]", fullFunName, fileName, line)
	}

	pkg := path.Join(path.Dir(fullFunName), parts[0])
	funName := parts[1]

	return fmt.Sprintf("[%s/%s:%s:%d]", pkg, fileName, funName, line)
}
