package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/gamedata/drop"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

func loadLiveDropContentGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveDropContentGroup")
	gamedata.LiveDropContentGroup = make(map[int32]*drop.DropList)

	type LiveDropContentGroup struct {
		Id                 int32 `xorm:"pk"`
		DropContentGroupId int32
		ContentType        int32
		ContentId          int32
		Amount             int32
	}

	groups := []LiveDropContentGroup{}
	err := masterdata_db.Table("m_live_drop_content").Find(&groups)
	utils.CheckErr(err)

	for _, item := range groups {
		_, exist := gamedata.LiveDropContentGroup[item.DropContentGroupId]
		if !exist {
			gamedata.LiveDropContentGroup[item.DropContentGroupId] = new(drop.DropList)
		}
		gamedata.LiveDropContentGroup[item.DropContentGroupId].AddContent(client.Content{
			ContentType:   item.ContentType,
			ContentId:     item.ContentId,
			ContentAmount: item.Amount,
		})
	}
}

func init() {
	addLoadFunc(loadLiveDropContentGroup)
}
