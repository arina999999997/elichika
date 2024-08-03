package serverdata

import (
	"elichika/config"
	"elichika/utils"

	"os"

	"xorm.io/xorm"
)

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", config.ServerdataPath)
	utils.CheckErr(err)
	Engine.SetMaxOpenConns(50)
	Engine.SetMaxIdleConns(10)
	overwrite = (len(os.Args) == 2) && (os.Args[1] == "reinit")
	InitTables()
}
