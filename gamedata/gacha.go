package gamedata

import (
	"elichika/dictionary"
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type Gacha = serverdata.ServerGacha

func populateGacha(gacha *Gacha, gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
}

func loadGacha(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.Gacha = make(map[int32]*Gacha)
	err := serverdata_db.Table("s_gacha").OrderBy("gacha_master_id").Find(&gamedata.GachaList)
	utils.CheckErr(err)
	for _, gacha := range gamedata.GachaList {
		gamedata.Gacha[gacha.GachaMasterId] = gacha
		populateGacha(gacha, gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadGacha)
}
