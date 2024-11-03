package gamedata

import (
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type GachaGroup struct {
	// s_gacha_group
	GroupMasterId int32 `xorm:"pk 'group_master_id'"`
	GroupWeight   int64 `xorm:"'group_weight'"`
	// s_gacha_card
	Cards []serverdata.GachaCard `xorm:"-"`
}

func (gachaGroup *GachaGroup) PickRandomCard(randOutput int64) int32 {
	return gachaGroup.Cards[int(randOutput%int64(len(gachaGroup.Cards)))].CardMasterId
}

func (gachaGroup *GachaGroup) populate(gamedata *Gamedata) {
	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_gacha_card").Where("group_master_id = ?", gachaGroup.GroupMasterId).Find(&gachaGroup.Cards)
	})
	utils.CheckErr(err)
}

func loadGachaGroup(gamedata *Gamedata) {
	gamedata.GachaGroup = make(map[int32]*GachaGroup)
	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_gacha_group").Find(&gamedata.GachaGroup)
	})
	utils.CheckErr(err)
	for _, gachaGroup := range gamedata.GachaGroup {
		gachaGroup.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadGachaGroup)
}
