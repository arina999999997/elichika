package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type CustomBackground struct {
	// from m_custom_background
	Id int32 `xorm:"pk 'id'"`
	// ...
}

func loadCustomBackground(gamedata *Gamedata) {
	fmt.Println("Loading CustomBackground")
	gamedata.CustomBackground = make(map[int32]*CustomBackground)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_custom_background").Find(&gamedata.CustomBackground)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadCustomBackground)
}
