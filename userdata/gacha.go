package userdata

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
)

func (session *Session) GetGachaList() []model.Gacha {
	gachaList := []model.Gacha{}
	// the code is like this because gacha might also contain personal data
	// it's not handled for now
	for _, gacha := range session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).GachaList {
		// skip the tutorial gacha if already done with tutorial
		if gacha.GachaMasterId == 999999 && session.UserStatus.TutorialPhase != enum.TutorialPhaseGacha {
			continue
		}
		gachaList = append(gachaList, *gacha)
	}
	return gachaList
}

func (session *Session) GetGacha(gachaMasterId int) model.Gacha {
	return *session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).Gacha[gachaMasterId]
}
