package gamedata

import (
	"elichika/client"

	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryMeltGroup struct {
	Id     int32          `xorm:"pk 'id'"`
	Reward client.Content `xorm:"extends"`
}

func loadAccessoryMeltGroup(gamedata *Gamedata) {
	gamedata.AccessoryMeltGroup = make(map[int32]*AccessoryMeltGroup)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_accessory_melt_group").Find(&gamedata.AccessoryMeltGroup)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryMeltGroup)
}
