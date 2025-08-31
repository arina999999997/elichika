package log

import (
	"io"
	goLog "log"
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
