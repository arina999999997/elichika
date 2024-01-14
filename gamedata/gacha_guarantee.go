package gamedata

import (
	"elichika/dictionary"
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type GachaGuarantee = serverdata.GachaGuarantee

func populateGachaGuarantee(gachaGuarantee *GachaGuarantee, gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	if len(gachaGuarantee.CardSetSQL) > 0 {
		cards := []int32{}
		err := masterdata_db.Table("m_card").Where(gachaGuarantee.CardSetSQL).Cols("id").Find(&cards)
		utils.CheckErr(err)
		gachaGuarantee.GuaranteedCardSet = make(map[int32]bool)
		for _, card := range cards {
			gachaGuarantee.GuaranteedCardSet[card] = true
		}
	}
}

func loadGachaGuarantee(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.GachaGuarantee = make(map[int32]*GachaGuarantee)
	err := serverdata_db.Table("s_gacha_guarantee").Find(&gamedata.GachaGuarantee)
	utils.CheckErr(err)
	for _, gachaGuarantee := range gamedata.GachaGuarantee {
		populateGachaGuarantee(gachaGuarantee, gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadGachaGuarantee)
}
