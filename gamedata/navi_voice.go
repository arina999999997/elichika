package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type NaviVoice struct {
	Id       int32 `xorm:"pk 'id'"`
	ListType int32 `xorm:"'list_type'" enum:"navi_voice_list_type"`
}

func loadNaviVoice(gamedata *Gamedata) {
	fmt.Println("Loading NaviVoice")
	gamedata.NaviVoice = make(map[int32]*NaviVoice)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_navi_voice").Find(&gamedata.NaviVoice)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadNaviVoice)
}
