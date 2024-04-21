package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type NaviVoice struct {
	Id       int32 `xorm:"pk 'id'"`
	ListType int32 `xorm:"'list_type'" enum:"navi_voice_list_type"`
}

func loadNaviVoice(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading NaviVoice")
	gamedata.NaviVoice = make(map[int32]*NaviVoice)
	err := masterdata_db.Table("m_navi_voice").Find(&gamedata.NaviVoice)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadNaviVoice)
}
