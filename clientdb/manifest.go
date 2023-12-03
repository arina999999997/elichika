package clientdb

import (
	"elichika/utils"

	"crypto/sha1"
	"encoding/hex"
	"os"
)

type manifest struct {
	Version string
	Lang    string
	Files   []FileReference
}

func readBytes(file *os.File, size int) []byte {
	bytes := make([]byte, size)
	file.Read(bytes)
	return bytes
}

func readPrefixString(file *os.File) string {
	return string(readBytes(file, readUbyte(file)))
}

func readUbyte(file *os.File) int {
	b := readBytes(file, 1)
	return int(b[0])
}

func readUint(file *os.File) int {
	b := readBytes(file, 4)
	f := int(b[3]) << 8
	f = (f + int(b[2])) << 8
	f = (f + int(b[1])) << 7
	f = (f + int(b[0])) // cursed af
	return f
}

func serializeUbyte(b byte) []byte {
	bytes := make([]byte, 1)
	bytes[0] = b
	return bytes
}

func serializeUint(i int) []byte {
	b := make([]byte, 4)
	b[0] = byte(i & 0xff)
	b[0] |= 0x80
	i -= int(b[0])
	b[1] = byte((i >> 7) & 0xff)
	i = (i >> 7) - int(b[1])
	b[2] = byte((i >> 8) & 0xff)
	b[3] = byte((i >> 16) & 0xff)
	return b
}

func serializePrefixString(s string) []byte {
	b := []byte{}
	b = append(b, byte(len(s)))
	b = append(b, []byte(s)...)
	return b
}

func (m *manifest) serialize() []byte {
	buf := []byte{}

	buf = append(buf, serializePrefixString(m.Version)...)
	buf = append(buf, serializePrefixString(m.Lang)...)
	buf = append(buf, serializeUbyte(byte(len(m.Files)))...)

	for _, f := range m.Files {
		buf = append(buf, serializePrefixString(f.Name)...)
		buf = append(buf, serializePrefixString(f.SHA)...)
	}
	for _, f := range m.Files {
		bytes, err := hex.DecodeString(f.EncryptedSHA)
		utils.CheckErr(err)
		buf = append(buf, bytes...)
		buf = append(buf, serializeUint(f.Size)...)
	}
	sha1 := sha1.Sum(buf)
	buf = append(sha1[:], buf...)
	return buf
}

func newManifest(file *os.File) *manifest {
	m := new(manifest)
	_ = readBytes(file, 20) // sha1hash
	m.Version = readPrefixString(file)
	m.Lang = readPrefixString(file)

	m.Files = []FileReference{}

	fileCount := readUbyte(file)
	for i := 0; i < fileCount; i++ {
		name := readPrefixString(file)
		sha := readPrefixString(file)
		m.Files = append(m.Files, makeFileReference(m.Version, name, sha))
	}

	for i := range m.Files {
		m.Files[i].EncryptedSHA = hex.EncodeToString(readBytes(file, 20))
		m.Files[i].Size = readUint(file)
	}
	return m
}
