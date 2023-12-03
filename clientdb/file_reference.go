package clientdb

import (
	"elichika/utils"

	"strconv"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
)

type FileReference struct {
	Version      string
	Name         string
	SHA          string
	EncryptedSHA string
	Size         int
}

func makeFileReference(version, name, sha string) FileReference {
	file := FileReference{}
	file.Version = version
	file.Name = name
	file.SHA = sha
	return file
}

func (file *FileReference) getKey(initKey hwdecrypt.HwdKeyset) hwdecrypt.HwdKeyset {
	k1, err := strconv.ParseUint(file.SHA[:8], 16, 32)
	utils.CheckErr(err)
	k2, err := strconv.ParseUint(file.SHA[8:16], 16, 32)
	utils.CheckErr(err)
	k3, err := strconv.ParseUint(file.SHA[16:24], 16, 32)
	utils.CheckErr(err)
	return hwdecrypt.HwdKeyset{
		Key1: initKey.Key1 ^ uint32(k1),
		Key2: initKey.Key2 ^ uint32(k2),
		Key3: initKey.Key3 ^ uint32(k3),
	}
}
