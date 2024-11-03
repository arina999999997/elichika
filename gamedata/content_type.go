package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type ContentType struct {
	Id int32 `xorm:"pk"`
	// AmountText string
	IsUnique bool
	// DisplayOrder int32
}

func loadContentType(gamedata *Gamedata) {
	fmt.Println("Loading ContentType")
	gamedata.ContentType = make(map[int32]*ContentType)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_content_setting").Find(&gamedata.ContentType)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadContentType)
}
