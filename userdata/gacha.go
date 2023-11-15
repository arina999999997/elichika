package userdata

import (
	"elichika/gamedata"
	"elichika/model"
)

func (session *Session) GetGachaList() []model.Gacha {
	gachaList := []model.Gacha{}
	// the code is like this because gacha might also contain personal data
	// it's not handled for now
	for _, gacha := range session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).GachaList {
		gachaList = append(gachaList, *gacha)
	}
	return gachaList
}

func (session *Session) GetGacha(gachaMasterID int) model.Gacha {
	return *session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).Gacha[gachaMasterID]
}
