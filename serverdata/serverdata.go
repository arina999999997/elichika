package serverdata

import (
	"elichika/db"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

// the serverdata system work as follow:
// - each table has a defined structure and an initializer, which can be null
// - if a table is new or empty, the initializer is called
// - howerver, all tables are created before any intializer is called, so one initializer can initalize multiple tables

type Initializer = func(*xorm.Session)

var (
	engine                           *xorm.Engine
	Database                         *db.DatabaseSync
	serverDataTableNameToInterface   = map[string]interface{}{}
	serverDataTableNameToInitializer = map[string]Initializer{}

	// whether to rebuild the assets
	// setting this to true should only update the assets to the newest version, and if the version are the same, it should not change anything
	rebuildAsset bool

	// whether to reset the server state
	// this do not delete any user data, it only reset the server to the initial state
	// can be used if the server sided tasks are somehow in a bad state
	// this will almost certainly disrupt on-going events
	resetServer bool
)

func addTable(tableName string, structure interface{}, initializer Initializer) {
	_, exist := serverDataTableNameToInterface[tableName]
	if exist {
		panic("table already exist: " + tableName)
	}
	serverDataTableNameToInterface[tableName] = structure
	serverDataTableNameToInitializer[tableName] = initializer
}

func createTable(tableName string, structure interface{}, overwrite bool) bool {
	exist, err := engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)

	if !exist {
		fmt.Println("Creating new table:", tableName)
		err = engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
		return true
	} else {
		if !overwrite {
			return false
		}
		fmt.Println("Overwrite existing table:", tableName)
		err := engine.DropTables(tableName)
		utils.CheckErr(err)
		err = engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
		return true
	}
}

func isTableEmpty(tableName string) bool {
	total, err := engine.Table(tableName).Count()
	utils.CheckErr(err)
	return total == 0
}

func InitTables() {
	isServerState := map[string]bool{}
	isServerState["s_scheduled_task"] = true
	isServerState["s_event_active"] = true
	initializers := []Initializer{}
	for tableName := range serverDataTableNameToInterface {
		overwrite := rebuildAsset
		if isServerState[tableName] {
			overwrite = resetServer
		}
		newOrEmpty := createTable(tableName, serverDataTableNameToInterface[tableName], overwrite)
		newOrEmpty = newOrEmpty || isTableEmpty(tableName)
		if newOrEmpty {
			initializers = append(initializers, serverDataTableNameToInitializer[tableName])
		}
	}
	session := engine.NewSession()
	defer session.Close()
	session.Begin()
	for _, initializer := range initializers {
		if initializer == nil {
			continue
		}
		initializer(session)
	}
	session.Commit()
}
