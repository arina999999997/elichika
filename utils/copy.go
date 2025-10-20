package utils

import (
	"io"
	"os"
)

func CopyFile(from, to string) {
	buf := make([]byte, 1024)
	source, err := os.Open(from)
	CheckErr(err)
	defer source.Close()
	dest, err := os.Create(to)
	CheckErr(err)
	defer dest.Close()
	for {
		read, err := source.Read(buf)
		if (err != nil) && (err != io.EOF) {
			CheckErr(err)
		}
		if read == 0 {
			break
		}
		_, err = dest.Write(buf[:read])
		CheckErr(err)
	}
}
