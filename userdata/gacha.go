package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
)

func (session *Session) GetGachaList() generic.List[client.Gacha] {
	gachaList := generic.List[client.Gacha]{}
	// the code is like this because gacha might also contain personal data
	// it's not handled for now though
	for _, gacha := range session.Ctx.MustGet("gamedata").(*gamedata.Gamedata).GachaList {
		// skip the tutorial gacha if already done with tutorial
		if gacha.GachaMasterId == 999999 && session.UserStatus.TutorialPhase != enum.TutorialPhaseGacha {
			continue
		}
		gachaList.Append(gacha.ClientGacha)
	}
	return gachaList
}
