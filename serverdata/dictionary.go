package serverdata

import (
	"elichika/config"
	"elichika/utils"

	"encoding/json"

	"xorm.io/xorm"
)

type DictionaryItem struct {
	Id      string `xorm:"pk 'id'"`
	Message string `xorm:"'message'"`
}

func initDictionary(session *xorm.Session) {
	var err error

	text := utils.ReadAllText(config.AssetPath + "dictionary.json")
	type DictionaryValue struct {
		Ja *string `json:"ja"`
		En *string `json:"en"`
		Ko *string `json:"ko"`
		Zh *string `json:"zh"`
	}
	values := map[string]DictionaryValue{}
	err = json.Unmarshal([]byte(text), &values)
	utils.CheckErr(err)
	items := map[string][]DictionaryItem{}
	for id, value := range values {
		items["ja"] = append(items["ja"], DictionaryItem{
			Id:      id,
			Message: *value.Ja,
		})
		items["en"] = append(items["en"], DictionaryItem{
			Id:      id,
			Message: *value.En,
		})
		items["ko"] = append(items["ko"], DictionaryItem{
			Id:      id,
			Message: *value.Ko,
		})
		items["zh"] = append(items["zh"], DictionaryItem{
			Id:      id,
			Message: *value.Zh,
		})
	}
	for language, values := range items {
		_, err = session.Table("s_dictionary_" + language).Insert(values)
		utils.CheckErr(err)
	}
}

func init() {
	addTable("s_dictionary_ja", DictionaryItem{}, initDictionary)
	addTable("s_dictionary_en", DictionaryItem{}, nil)
	addTable("s_dictionary_ko", DictionaryItem{}, nil)
	addTable("s_dictionary_zh", DictionaryItem{}, nil)
}
