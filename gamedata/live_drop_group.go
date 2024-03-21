package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/generic/drop"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveDropGroup struct {
	Bad                             bool
	GroupId                         int32
	LiveDropContentGroupByDropColor map[int32]*drop.DropList[*drop.WeightedDropList[client.Content]] `xorm:"-"`
}

func (ldg *LiveDropGroup) Check() {
	if ldg == nil {
		return
	}
	if ldg.Bad {
		panic(fmt.Sprint("bad live drop group:", ldg.GroupId))
	}
}

func (ldg *LiveDropGroup) GetRandomItemByDropColor(dropColor int32) client.Content {
	_, exist := ldg.LiveDropContentGroupByDropColor[dropColor]
	if !exist {
		for fallback := range ldg.LiveDropContentGroupByDropColor {
			dropColor = fallback
			break
		}
	}
	return ldg.LiveDropContentGroupByDropColor[dropColor].GetRandomItem().GetRandomItem()
}

func loadLiveDropGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveDropGroup")
	gamedata.LiveDropGroup = make(map[int32]*LiveDropGroup)

	type LiveDropGroupRow struct {
		Id                 int32
		GroupId            int32
		DropColor          int32
		DropCount          int32
		DropContentGroupId int32
	}
	rows := []LiveDropGroupRow{}
	err := masterdata_db.Table("m_live_drop_group").Find(&rows)
	utils.CheckErr(err)

	for _, row := range rows {
		if gamedata.LiveDropContentGroup[row.DropContentGroupId] == nil {
			continue
		}

		_, exist := gamedata.LiveDropGroup[row.GroupId]
		if !exist {
			gamedata.LiveDropGroup[row.GroupId] = new(LiveDropGroup)
			gamedata.LiveDropGroup[row.GroupId].GroupId = row.GroupId
			gamedata.LiveDropGroup[row.GroupId].LiveDropContentGroupByDropColor = map[int32]*drop.DropList[*drop.WeightedDropList[client.Content]]{}
		}
		_, exist = gamedata.LiveDropGroup[row.GroupId].LiveDropContentGroupByDropColor[row.DropColor]
		if !exist {
			gamedata.LiveDropGroup[row.GroupId].LiveDropContentGroupByDropColor[row.DropColor] = new(drop.DropList[*drop.WeightedDropList[client.Content]])
		}
		gamedata.LiveDropGroup[row.GroupId].LiveDropContentGroupByDropColor[row.DropColor].AddItem(gamedata.LiveDropContentGroup[row.DropContentGroupId])
	}
}

func init() {
	addLoadFunc(loadLiveDropGroup)
	addPrequisite(loadLiveDropGroup, loadLiveDropContentGroup)

}
