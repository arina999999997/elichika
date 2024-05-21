package serverdata

import (
	"elichika/config"
	"elichika/utils"

	"encoding/json"

	"xorm.io/xorm"
)

type NgWord struct {
	Word string `xorm:"pk"`
}

func InitialiseNgWord(session *xorm.Session) {
	files := []string{config.ServerInitJsons + "wordlist_gl.json", config.ServerInitJsons + "wordlist_jp.json"}
	for _, file := range files {
		wordsJson := utils.ReadAllText(file)
		words := []string{}
		err := json.Unmarshal([]byte(wordsJson), &words)
		utils.CheckErr(err)
		for _, word := range words {
			ngWord := NgWord{
				Word: word,
			}
			exist, err := session.Table("s_ng_word").Exist(&ngWord)
			utils.CheckErr(err)
			if !exist {
				_, err = session.Table("s_ng_word").Insert(&ngWord)
				utils.CheckErr(err)
			}
		}
	}
}
