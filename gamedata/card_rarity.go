package gamedata

import (
	"elichika/utils"

	"xorm.io/xorm"
)

type CardRarity struct {
	CardRarityType int32 `xorm:"pk"`
	MaxLevel       int32
	PlusLevel      int32
}

func loadCardRarity(gamedata *Gamedata) {
	gamedata.CardRarity = make(map[int32]*CardRarity)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_card_rarity").Find(&gamedata.CardRarity)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadCardRarity)
}
