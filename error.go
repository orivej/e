package e

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

var Output io.Writer = os.Stdout

func context() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	fn := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s (%s:%d)", fn.Name(), path.Base(file), line)
}

func Exit(e error) {
	if e != nil {
		fmt.Fprintf(Output, "%s: %s\n", context(), e)
		os.Exit(1)
	}
}

func Print(e error) bool {
	if e != nil {
		fmt.Fprintf(Output, "%s: %s\n", context(), e)
	}
	return e != nil
}

func Panic(e error) {
	if e != nil {
		fmt.Fprintf(Output, "%s: %s\n", context(), e)
		panic(e)
	}
}

func CloseOrExit(c io.Closer) {
	e := c.Close()
	if e != nil {
		fmt.Fprintf(Output, "%s: %s\n", context(), e)
		os.Exit(1)
	}
}

func CloseOrPrint(c io.Closer) bool {
	e := c.Close()
	if e != nil {
		fmt.Fprintf(Output, "%s: %s\n", context(), e)
	}
	return e != nil
}
