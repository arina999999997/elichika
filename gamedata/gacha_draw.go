package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"xorm.io/xorm"
)

type GachaDraw = model.GachaDraw

func loadGachaDraw(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.GachaDraw = make(map[int]*GachaDraw)
	err := serverdata_db.Table("s_gacha_draw").Find(&gamedata.GachaDraw)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadGachaDraw)
	// addLoadFunc(loadGachaGuarantee)
}
