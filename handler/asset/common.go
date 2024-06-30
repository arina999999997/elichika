package asset

import (
	"elichika/utils"

	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func sendRange(ctx *gin.Context, fileName string, start, size int) {
	ctx.Header("Content-Length", fmt.Sprint(size))
	ctx.Header("Content-Type", "application/octet-stream")

	buffer := make([]byte, 1024)
	f, err := os.Open(fileName)
	utils.CheckErr(err)
	_, err = f.Seek(int64(start), io.SeekStart)
	utils.CheckErr(err)
	for ; size > 0; size -= 1024 {
		count, err := f.Read(buffer)
		utils.CheckErr(err)
		if count > size {
			count = size
		} else if (count < 1024) && (count < size) {
			panic("wrong requested range")
		}
		buffer = buffer[:count]
		ctx.Writer.Write(buffer)
	}
}
