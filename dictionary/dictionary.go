// dictionary db
package dictionary

import (
	"elichika/client"
	_ "elichika/clientdb"
	"elichika/generic"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"
	"strings"

	"xorm.io/xorm"
)

// TODO(dictionary): just move this inside gamedata or something
type Dictionary struct {
	Language         string
	Dictionaries     map[string]*xorm.Engine
	ClientDictionary map[string]string // there is no need to load the whole client dictionary, but we might use some keys a lot
	ServerDictionary map[string]string
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

	var err error
	for _, dictType := range dictionaryTypes {
		dictionary.Dictionaries[dictType], err = xorm.NewEngine("sqlite", path+"dictionary_"+language+"_"+dictType+".db")
		utils.CheckErr(err)
		dictionary.Dictionaries[dictType].SetMaxOpenConns(50)
		dictionary.Dictionaries[dictType].SetMaxIdleConns(10)
	}

	type Pair struct {
		Id      string `xorm:"pk 'id'"`
		Message string `xorm:"'message'"`
	}

	dictionary.ServerDictionary = map[string]string{}
	dictionary.ClientDictionary = map[string]string{}
	pairs := []Pair{}
	serverdata.Database.Do(func(session *xorm.Session) {
		err = session.Table("s_dictionary_" + language).Find(&pairs)
	})
	utils.CheckErr(err)
	for _, p := range pairs {
		dictionary.ServerDictionary[p.Id] = p.Message
	}
}

func (dictionary *Dictionary) Resolve(id string) string {
	result, exist := dictionary.ClientDictionary[id]
	if exist {
		return result
	}
	keys := strings.Split(id, ".")
	if dictionary.Dictionaries[keys[0]] == nil {
		return id
	}
	result = ""
	exist, err := dictionary.Dictionaries[keys[0]].Table("m_dictionary").Where("id = ?", keys[1]).Cols("message").Get(&result)
	utils.CheckErr(err)
	if !exist {
		fmt.Printf("Warning: client dictionary key doesn't exist: %s\n", id)
	}
	dictionary.ClientDictionary[id] = result
	return result
}

func (dictionary *Dictionary) ServerResolve(id string) string {
	message, exist := dictionary.ServerDictionary[id]
	if !exist {
		panic("server dictionary key doesn't exist: " + id)
	}
	return message
}

// try to resolve id using server / client database
// if not possible, resolve to id itself
func (dictionary *Dictionary) UniversalResolve(id string) string {
	message, exist := dictionary.ServerDictionary[id]
	if exist {
		return message
	}
	return dictionary.Resolve(id)
}

func (dictionary *Dictionary) ResolveServerLocalizedText(item generic.Nullable[client.LocalizedText]) generic.Nullable[client.LocalizedText] {
	if item.HasValue {
		// server localisation rule:
		// - split the value by " "
		// - then try to key each of the items, then glue them all together, without the space
		// - this design force space into the word, because different language handle spacing differently
		words := strings.Split(item.Value.DotUnderText, " ")
		item.Value.DotUnderText = ""
		for _, word := range words {
			item.Value.DotUnderText += dictionary.UniversalResolve(word)
		}
	}
	return item
}
