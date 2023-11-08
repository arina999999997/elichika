package userdata

import (
	"elichika/gamedata"
	"elichika/model"
)

func (session *Session) GetGachaList() []model.Gacha {
	gachaList := []model.Gacha{}
	for _, gacha := range session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).Gacha {
		gachaList = append(gachaList, *gacha)
	}
	return gachaList
}

func (session *Session) GetGacha(gachaMasterID int) model.Gacha {
	return *session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).Gacha[gachaMasterID]
}
