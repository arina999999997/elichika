package gamedata

import (
	"elichika/model"
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryMeltGroup struct {
	ID int `xorm:"pk 'id'"`
	Reward model.Content `xorm:"extends"`
}

func loadAccessoryMeltGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryMeltGroup = make(map[int]*AccessoryMeltGroup)
	err := masterdata_db.Table("m_accessory_melt_group").Find(&gamedata.AccessoryMeltGroup)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryMeltGroup)
}