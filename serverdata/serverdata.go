package serverdata

import (
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
	Engine                           *xorm.Engine
	serverDataTableNameToInterface   = map[string]interface{}{}
	serverDataTableNameToInitializer = map[string]Initializer{}
	overwrite                        bool
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
	exist, err := Engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)

	if !exist {
		fmt.Println("Creating new table:", tableName)
		err = Engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
		return true
	} else {
		if !overwrite {
			return false
		}
		fmt.Println("Overwrite existing table:", tableName)
		err := Engine.DropTables(tableName)
		utils.CheckErr(err)
		err = Engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
		return true
	}
}

func isTableEmpty(tableName string) bool {
	total, err := Engine.Table(tableName).Count()
	utils.CheckErr(err)
	return total == 0
}

func InitTables() {
	initializers := []Initializer{}
	for tableName := range serverDataTableNameToInterface {
		newOrEmpty := createTable(tableName, serverDataTableNameToInterface[tableName], overwrite)
		newOrEmpty = newOrEmpty || isTableEmpty(tableName)
		if newOrEmpty {
			initializers = append(initializers, serverDataTableNameToInitializer[tableName])
		}
	}
	session := Engine.NewSession()
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
