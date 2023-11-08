package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"xorm.io/xorm"
)

// Gacha design is not very good, so use modelGacha for now
type Gacha = model.Gacha

func populateGacha(gacha *Gacha, gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := serverdata_db.Table("s_gacha_appeal").In("gacha_appeal_master_id", gacha.DbGachaAppeals).
		Find(&gacha.GachaAppeals)
	utils.CheckErr(err)
	err = serverdata_db.Table("s_gacha_draw").In("gacha_draw_master_id", gacha.DbGachaDraws).
		Find(&gacha.GachaDraws)
	utils.CheckErr(err)
}

func loadGacha(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.Gacha = make(map[int]*Gacha)
	err := serverdata_db.Table("s_gacha").Find(&gamedata.Gacha)
	utils.CheckErr(err)
	for _, gacha := range gamedata.Gacha {
		populateGacha(gacha, gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadGacha)
}
