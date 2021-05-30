package main

import (
	"io"
	"os"
)

type MyWriter struct{ w io.Writer }

func (mw *MyWriter) WriteString(s string) (n int, err error) {
	return io.WriteString(mw.w, s)
}

func main() {
	var sw io.StringWriter = &MyWriter{w: os.Stdout}
	sw.WriteString("hello, world")
}
