package log

import (
	"io"
	goLog "log"
	"runtime/debug"
)

func Println(v ...any) {
	goLog.Println(v...)
}

func Printf(v ...any) {
	goLog.Printf(v[0].(string), v[1:]...)
}

func SetOutput(w io.Writer) {
	goLog.SetOutput(w)
}

// capturing panic is annoying, so we will just replace panic with our own
func Panic(errAny any) {
	err := errAny.(error)
	Println(err)
	Println(string(debug.Stack()))
	panic(err)
}
