package gamedata

import (
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type Gacha = serverdata.ServerGacha

func populateGacha(gacha *Gacha, gamedata *Gamedata) {
}

func loadGacha(gamedata *Gamedata) {
	gamedata.Gacha = make(map[int32]*Gacha)
	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_gacha").OrderBy("gacha_master_id").Find(&gamedata.GachaList)
	})
	utils.CheckErr(err)
	for _, gacha := range gamedata.GachaList {
		gamedata.Gacha[gacha.GachaMasterId] = gacha
		populateGacha(gacha, gamedata)
	}
}

func init() {
	addLoadFunc(loadGacha)
}
