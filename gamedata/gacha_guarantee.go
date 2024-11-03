package gamedata

import (
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type GachaGuarantee = serverdata.GachaGuarantee

func populateGachaGuarantee(gachaGuarantee *GachaGuarantee, gamedata *Gamedata) {
	if len(gachaGuarantee.CardSetSQL) > 0 {
		cards := []int32{}
		var err error
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_card").Where(gachaGuarantee.CardSetSQL).Cols("id").Find(&cards)
		})
		utils.CheckErr(err)
		gachaGuarantee.GuaranteedCardSet = make(map[int32]bool)
		for _, card := range cards {
			gachaGuarantee.GuaranteedCardSet[card] = true
		}
	}
}

func loadGachaGuarantee(gamedata *Gamedata) {
	gamedata.GachaGuarantee = make(map[int32]*GachaGuarantee)
	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_gacha_guarantee").Find(&gamedata.GachaGuarantee)
	})
	utils.CheckErr(err)
	for _, gachaGuarantee := range gamedata.GachaGuarantee {
		populateGachaGuarantee(gachaGuarantee, gamedata)
	}
}

func init() {
	addLoadFunc(loadGachaGuarantee)
}
