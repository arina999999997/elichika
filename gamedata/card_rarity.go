package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type CardRarity struct {
	CardRarityType int32 `xorm:"pk"`
	// MaxLevel int32
	PlusLevel int32
}

func loadCardRarity(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.CardRarity = make(map[int32]*CardRarity)
	err := masterdata_db.Table("m_card_rarity").Find(&gamedata.CardRarity)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadCardRarity)
}
