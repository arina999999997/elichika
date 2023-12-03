package clientdb

import (
	"elichika/config"
	"elichika/utils"

	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
)

var (
	keySpec = hwdecrypt.HwdKeyset{
		Key1: 0x49e66da3,
		Key2: 0x59e1e89a,
		Key3: 0x24ebb207,
	}
)

func versionString(shaInHex string) string {
	// take the top xor with the bottom
	bytes, _ := hex.DecodeString(shaInHex)
	length := len(bytes) / 2
	return hex.EncodeToString(utils.Xor(bytes[:length], bytes[length:]))
}

func fileSha1(fileName string) []byte {

	file, err := os.ReadFile(fileName)
	if err != nil {
		file = []byte{}
	}
	sha1 := sha1.Sum(file)
	return sha1[:]
}

func updateClientDb(baseDir string, masterdataRefs []string) {
	versionBytes := []byte{}
	oldVersion := ""
	newFileVersion := map[string]string{}
	mMap := map[string]*manifest{}
	for _, refPath := range masterdataRefs {
		fmt.Println(baseDir + refPath)
		f, err := os.Open(baseDir + refPath)
		utils.CheckErr(err)
		defer f.Close()
		mMap[refPath] = newManifest(f)
		if oldVersion == "" {
			oldVersion = mMap[refPath].Version
		} else if oldVersion != mMap[refPath].Version {
			panic("manifest files must come from the same version")
		}
		for _, file := range mMap[refPath].Files {
			actualFileVersion := hex.EncodeToString(fileSha1(baseDir + file.Name))
			if actualFileVersion != file.SHA {
				newFileVersion[file.Name] = actualFileVersion
			}
			versionBytes = append(versionBytes, []byte(actualFileVersion)...)
		}
	}
	version := md5.Sum(versionBytes) // use md5 because it generate a 16 bytes string, which we can fold into 16 hex char
	newVersion := versionString(hex.EncodeToString(version[:]))
	os.MkdirAll(config.StaticDataPath+newVersion+"/", 0755)
	if newVersion == oldVersion {
		// check for integrity  of the data
		fmt.Println("Perform integrity check for version: ", oldVersion)
		updated := false
		for _, m := range mMap {
			for i := range m.Files {
				inputFile := baseDir + m.Files[i].Name
				existingFile := config.StaticDataPath + oldVersion + "/" + m.Files[i].Name
				if hex.EncodeToString(fileSha1(existingFile)) != m.Files[i].EncryptedSHA {
					fmt.Println("Integrity error detected for file: ", m.Files[i].Name, ", rekeying")
					// remove the file first, if it exists
					if fileExist(existingFile) {
						err := os.Remove(existingFile)
						utils.CheckErr(err)
					}
					rekey(inputFile, existingFile, &m.Files[i], keySpec)
					updated = true
				}
			}
		}
		if updated { // update all the manifest
			for refPath, m := range mMap {
				serialized := m.serialize()
				err := os.WriteFile(config.StaticDataPath+oldVersion+"/"+refPath, serialized, 0777)
				utils.CheckErr(err)
				err = os.WriteFile(baseDir+refPath, serialized, 0777)
				utils.CheckErr(err)
			}
		}
	} else {
		fmt.Println("Difference detected, new version: ", newVersion)
		for refPath, m := range mMap {
			m.Version = newVersion
			for i := range m.Files {
				ver, exists := newFileVersion[m.Files[i].Name]
				inputFile := baseDir + m.Files[i].Name
				outputFile := config.StaticDataPath + newVersion + "/" + m.Files[i].Name
				if !exists {
					// no need to change this file, but we have to copy it over from the existing version
					err := copy(config.StaticDataPath+oldVersion+"/"+m.Files[i].Name, outputFile)
					// if there is an error then something is wrong, here we assume the old version doesn't already exists, so we just key this file anyway
					if err == nil {
						continue
					}
					ver = m.Files[i].SHA
				}
				fmt.Println("Rekeying file: ", m.Files[i].Name)
				m.Files[i].SHA = ver
				rekey(inputFile, outputFile, &m.Files[i], keySpec)
			}
			serialized := m.serialize()
			err := os.WriteFile(config.StaticDataPath+newVersion+"/"+refPath, serialized, 0777)
			utils.CheckErr(err)
			err = os.WriteFile(baseDir+refPath, serialized, 0777)
			utils.CheckErr(err)
		}
	}
}

func init() {
	// outDir := config.StaticDataPath

	// Gl
	masterdataRefGls := []string{}
	for _, p := range config.Platforms {
		for _, l := range config.GlLanguages {
			masterdataRefGls = append(masterdataRefGls, "masterdata_"+p+"_"+l)
		}
	}
	updateClientDb(config.GlDatabasePath, masterdataRefGls)
	// Jp
	masterdataRefJps := []string{}
	for _, p := range config.Platforms {
		for _, l := range config.JpLanguages {
			masterdataRefJps = append(masterdataRefJps, "masterdata_"+p+"_"+l)
		}
	}
	updateClientDb(config.JpDatabasePath, masterdataRefJps)

	config.MasterVersionGl = readMasterdataManinest(config.GlDatabasePath + "masterdata_a_en")
	config.MasterVersionJp = readMasterdataManinest(config.JpDatabasePath + "masterdata_a_ja")

	fmt.Println("gl master version:", config.MasterVersionGl)
	fmt.Println("jp master version:", config.MasterVersionJp)
}

func readMasterdataManinest(path string) string {
	file, err := os.Open(path)
	utils.CheckErr(err)
	buff := make([]byte, 1024)
	count, err := file.Read(buff)
	utils.CheckErr(err)
	if count < 21 {
		panic("file too short")
	}
	length := int(buff[20])
	version := string(buff[21 : 21+length])
	return version
}

func copy(in, out string) error {
	file, err := os.ReadFile(in)
	if err != nil {
		return err
	}
	err = os.WriteFile(out, file, 0777)
	utils.CheckErr(err)
	return nil
}
