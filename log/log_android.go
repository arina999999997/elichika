// go:build embedded
package log

/*
#cgo LDFLAGS: -llog

#include <android/log.h>
#include <stdlib.h>

void android_log(const char* data) {
    __android_log_write(ANDROID_LOG_INFO, "elichika", data);
}
*/
import "C"

// it's important for the import C to be right after the implementation
import (
	"unsafe"

	"github.com/gin-gonic/gin"
)

type AndroidLogger struct {
}

func (w AndroidLogger) Write(data []byte) (n int, err error) {
	dataStr := C.CString(string(data))
	defer C.free(unsafe.Pointer(dataStr))
	C.android_log(dataStr)
	return len(data), nil
}

func init() {
	// capture standard log
	SetOutput(AndroidLogger{})
	// capture gin's log
	gin.DefaultWriter = AndroidLogger{}
	Printf("Set android logger")
}
