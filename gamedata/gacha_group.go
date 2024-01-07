package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"xorm.io/xorm"
)

type GachaGroup struct {
	// s_gacha_group
	GroupMasterId int   `xorm:"pk 'group_master_id'"`
	GroupWeight   int64 `xorm:"'group_weight'"`
	// s_gacha_card
	Cards []model.GachaCard `xorm:"-"`
}

func (gachaGroup *GachaGroup) PickRandomCard(randOutput int64) int {
	return gachaGroup.Cards[int(randOutput%int64(len(gachaGroup.Cards)))].CardMasterId
}

func (gachaGroup *GachaGroup) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := serverdata_db.Table("s_gacha_card").Where("group_master_id = ?", gachaGroup.GroupMasterId).Find(&gachaGroup.Cards)
	utils.CheckErr(err)
}

func loadGachaGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.GachaGroup = make(map[int]*GachaGroup)
	err := serverdata_db.Table("s_gacha_group").Find(&gamedata.GachaGroup)
	utils.CheckErr(err)
	for _, gachaGroup := range gamedata.GachaGroup {
		gachaGroup.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadGachaGroup)
}
