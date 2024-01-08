package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryRarityUpGroup struct {
	// from m_accessory_rarity_up_group
	Id       int32          `xorm:"pk 'id'"`
	Resource client.Content `xorm:"extends"`
}

func loadAccessoryRarityUpGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryRarityUpGroup = make(map[int32]*AccessoryRarityUpGroup)
	err := masterdata_db.Table("m_accessory_rarity_up_group").Find(&gamedata.AccessoryRarityUpGroup)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryRarityUpGroup)
}
