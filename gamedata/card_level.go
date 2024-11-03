package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type CardLevel struct {
	// from m_card_level
	CardRarityType int `xorm:"pk 'card_rarity_type'"`
	// prefix sum so we can calculate total cost in O(1)
	ExpPrefixSum       []int32 `xorm:"-"`
	GameMoneyPrefixSum []int32 `xorm:"-"`
}

func (cardLevel *CardLevel) populate(gamedata *Gamedata) {
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_card_level").Where("card_rarity_type = ?", cardLevel.CardRarityType).OrderBy("level").Cols("exp").Find(&cardLevel.ExpPrefixSum)
	})
	utils.CheckErr(err)

	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_card_level").Where("card_rarity_type = ?", cardLevel.CardRarityType).OrderBy("level").Cols("game_money").Find(&cardLevel.GameMoneyPrefixSum)
	})
	utils.CheckErr(err)
	cardLevel.ExpPrefixSum = append([]int32{0}, cardLevel.ExpPrefixSum...)
	cardLevel.GameMoneyPrefixSum = append([]int32{0}, cardLevel.GameMoneyPrefixSum...)
	for i := 1; i < len(cardLevel.ExpPrefixSum); i++ {
		cardLevel.ExpPrefixSum[i] += cardLevel.ExpPrefixSum[i-1]
		cardLevel.GameMoneyPrefixSum[i] += cardLevel.GameMoneyPrefixSum[i-1]
	}
}

func loadCardLevel(gamedata *Gamedata) {
	fmt.Println("Loading CardLevel")
	gamedata.CardLevel = make(map[int32]*CardLevel)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_card_rarity").Find(&gamedata.CardLevel)
	})
	utils.CheckErr(err)
	for _, cardLevel := range gamedata.CardLevel {
		cardLevel.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadCardLevel)
}
