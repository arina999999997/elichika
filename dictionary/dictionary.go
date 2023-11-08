// dictionary db
package dictionary

import (
	"elichika/utils"

	"strings"

	"xorm.io/xorm"
)

type Dictionary struct {
	Language     string
	Dictionaries map[string]*xorm.Engine
}

func (dictionary *Dictionary) Init(path string, language string) {
	dictionaryTypes := []string{
		"v",
		"android",
		"dummy",
		"inline_image",
		"ios",
		"k",
		"m",
		"petag",
		// "s", // s has different structure
	}
	dictionary.Dictionaries = make(map[string]*xorm.Engine)

	for _, dictType := range dictionaryTypes {
		var err error
		dictionary.Dictionaries[dictType], err = xorm.NewEngine("sqlite", path+"dictionary_"+language+"_"+dictType+".db")
		utils.CheckErr(err)
		dictionary.Dictionaries[dictType].SetMaxOpenConns(50)
		dictionary.Dictionaries[dictType].SetMaxIdleConns(10)
	}
}

func (dictionary *Dictionary) Resolve(id string) string {
	keys := strings.Split(id, ".")
	res := ""
	exists, err := dictionary.Dictionaries[keys[0]].Table("m_dictionary").Where("id = ?", keys[1]).Cols("message").Get(&res)
	utils.CheckErrMustExist(err, exists)
	return res
}
