package clientdb

import (
	"elichika/log"
	"elichika/utils"

	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"

	"os"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
)

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// rekey and update the reference
func rekey(inPath, outPath string, fromFile *FileReference, keySpec hwdecrypt.HwdKeyset) {
	// read the clear database file
	if fileExist(outPath) {
		// skip if this file already exists. Most of the time it will be masterdata.db.
		// note that if we have to correct a file, we would delete that file before calling this
		log.Println("Skipping already generated file: ", outPath)
		// still need to update the file reference
		var bytes []byte
		bytes, fromFile.Size = fileSha1AndSize(outPath)
		fromFile.EncryptedSHA = hex.EncodeToString(bytes)
		return
	}
	clear_buf, err := os.ReadFile(inPath)
	utils.CheckErr(err)
	if string(clear_buf[:16]) != "SQLite format 3\x00" {
		panic("Missing SQLite file signature. Is it already encrypted?")
	}
	var crypt_buf bytes.Buffer
	zlibWriter := zlib.NewWriter(&crypt_buf)
	zlibWriter.Write(clear_buf)
	zlibWriter.Close()
	keyset := fromFile.getKey(keySpec)
	outputBytes := crypt_buf.Bytes()
	// skip the header and checksum.
	// TODO(extra): maybe write our own zlib manually using flate to have the same wbits things as python
	outputBytes = outputBytes[2 : len(outputBytes)-4]
	hwdecrypt.DecryptBuffer(&keyset, outputBytes)
	err = os.WriteFile(outPath, outputBytes, 0777)
	utils.CheckErr(err)
	fromFile.Size = len(outputBytes)
	sha1 := sha1.Sum(outputBytes)
	fromFile.EncryptedSHA = hex.EncodeToString(sha1[:])
}
