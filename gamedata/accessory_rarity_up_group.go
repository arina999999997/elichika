package gamedata

import (
	"elichika/client"

	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryRarityUpGroup struct {
	// from m_accessory_rarity_up_group
	Id       int32          `xorm:"pk 'id'"`
	Resource client.Content `xorm:"extends"`
}

func loadAccessoryRarityUpGroup(gamedata *Gamedata) {
	gamedata.AccessoryRarityUpGroup = make(map[int32]*AccessoryRarityUpGroup)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_accessory_rarity_up_group").Find(&gamedata.AccessoryRarityUpGroup)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryRarityUpGroup)
}
