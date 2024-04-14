package userdata

import (
	"elichika/userdata/database"

	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine = database.Engine
)

func Init() {
	database.InitTables(Engine)
}
