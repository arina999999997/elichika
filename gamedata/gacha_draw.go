package gamedata

import (
	"elichika/client"
	"elichika/dictionary"

	"xorm.io/xorm"
)

type GachaDraw = client.GachaDraw

type GachaDrawGuarantee struct {
	GuaranteeIds []int32
}

func loadGachaDrawGuarantee(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.GachaDraw = make(map[int32]*GachaDraw)
	gamedata.GachaDrawGuarantee = make(map[int32]*GachaDrawGuarantee)
	for _, gacha := range gamedata.Gacha {
		for i := range gacha.ClientGacha.GachaDraws.Slice {
			id := gacha.ClientGacha.GachaDraws.Slice[i].GachaDrawMasterId
			gamedata.GachaDraw[id] = &gacha.ClientGacha.GachaDraws.Slice[i]
			gamedata.GachaDrawGuarantee[id] = new(GachaDrawGuarantee)
			for _, guaranteed := range gacha.DrawGuarantees[i] {
				gamedata.GachaDrawGuarantee[id].GuaranteeIds = append(gamedata.GachaDrawGuarantee[id].GuaranteeIds, guaranteed)
			}
		}
	}
}

func init() {
	addLoadFunc(loadGachaDrawGuarantee)
	addPrequisite(loadGachaDrawGuarantee, loadGacha)
}
