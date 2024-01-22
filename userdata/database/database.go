package database

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

// This package handle the userdata database, from setting them up to updating them.
// Eventually it should be named userdata / userdatabase, while the current package should be named usermodel or something like that

var (
	Engine                       *xorm.Engine
	userDataTableNameToInterface = map[string]interface{}{}
)

func AddTable(tableName string, structure interface{}) {
	_, exist := userDataTableNameToInterface[tableName]
	if exist {
		panic("table name already used: " + tableName)
	}
	userDataTableNameToInterface[tableName] = structure
}

func InitTable(tableName string, structure interface{}, overwrite bool) {
	exist, err := Engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)

	if !exist {
		fmt.Println("Creating new table:", tableName)
		err = Engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
	} else {
		if !overwrite {
			return
		}
		fmt.Println("Overwrite existing table:", tableName)
		err := Engine.DropTables(tableName)
		utils.CheckErr(err)
		err = Engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
	}
}

func InitTables(overwrite bool) {
	for tableName, structure := range userDataTableNameToInterface {
		InitTable(tableName, structure, overwrite)
	}
}
