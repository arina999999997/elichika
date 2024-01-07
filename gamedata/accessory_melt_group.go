package gamedata

import (
	"elichika/dictionary"
	// "elichika/model"
	"elichika/client"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryMeltGroup struct {
	Id     int            `xorm:"pk 'id'"`
	Reward client.Content `xorm:"extends"`
}

func loadAccessoryMeltGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryMeltGroup = make(map[int]*AccessoryMeltGroup)
	err := masterdata_db.Table("m_accessory_melt_group").Find(&gamedata.AccessoryMeltGroup)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryMeltGroup)
}
