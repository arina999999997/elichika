package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryRarityUpGroup struct {
	// from m_accessory_rarity_up_group
	Id       int           `xorm:"pk 'id'"`
	Resource model.Content `xorm:"extends"`
}

func loadAccessoryRarityUpGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryRarityUpGroup = make(map[int]*AccessoryRarityUpGroup)
	err := masterdata_db.Table("m_accessory_rarity_up_group").Find(&gamedata.AccessoryRarityUpGroup)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryRarityUpGroup)
}
