package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type CustomBackground struct {
	// from m_custom_background
	Id int32 `xorm:"pk 'id'"`
	// ...
}

func loadCustomBackground(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading CustomBackground")
	gamedata.CustomBackground = make(map[int32]*CustomBackground)
	err := masterdata_db.Table("m_custom_background").Find(&gamedata.CustomBackground)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadCustomBackground)
}
