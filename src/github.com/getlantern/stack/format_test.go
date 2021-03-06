// +build go1.2

package stack_test

import (
	"deng/src/github.com/getlantern/stack"
	"fmt"
)

func Example_callFormat() {
	logCaller("%+s")
	logCaller("%v   %[1]n()")
	// Output:
	// github.com/getlantern/stack/format_test.go
	// format_test.go:13   Example_callFormat()
}

func logCaller(format string) {
	fmt.Printf(format+"\n", stack.Caller(1))
}
