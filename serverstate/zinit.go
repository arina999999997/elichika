package serverstate

import (
	"elichika/config"
	"elichika/db"
	"elichika/utils"

	"os"

	"xorm.io/xorm"
)

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite", config.ServerstatePath)
	utils.CheckErr(err)
	engine.SetMaxOpenConns(50)
	engine.SetMaxIdleConns(10)
	resetServer = (len(os.Args) >= 2) && (os.Args[1] == "reset_server")
	InitTables()
	engine.Close()
	Database, err = db.NewDatabase(config.ServerstatePath)
	utils.CheckErr(err)
}
