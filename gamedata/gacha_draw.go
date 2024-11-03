package gamedata

import (
	"elichika/client"
)

type GachaDraw = client.GachaDraw

type GachaDrawGuarantee struct {
	GuaranteeIds []int32
}

func loadGachaDrawGuarantee(gamedata *Gamedata) {
	gamedata.GachaDraw = make(map[int32]*GachaDraw)
	gamedata.GachaDrawGuarantee = make(map[int32]*GachaDrawGuarantee)
	for _, gacha := range gamedata.Gacha {
		for i := range gacha.ClientGacha.GachaDraws.Slice {
			id := gacha.ClientGacha.GachaDraws.Slice[i].GachaDrawMasterId
			gamedata.GachaDraw[id] = &gacha.ClientGacha.GachaDraws.Slice[i]
			gamedata.GachaDrawGuarantee[id] = new(GachaDrawGuarantee)
			gamedata.GachaDrawGuarantee[id].GuaranteeIds = append(gamedata.GachaDrawGuarantee[id].GuaranteeIds, gacha.DrawGuarantees[i]...)
		}
	}
}

func init() {
	addLoadFunc(loadGachaDrawGuarantee)
	addPrequisite(loadGachaDrawGuarantee, loadGacha)
}
