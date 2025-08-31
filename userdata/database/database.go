package database

import (
	"elichika/log"
	"elichika/utils"

	"xorm.io/xorm"
)

// This package handle the userdata database, from setting them up to updating them.
// Eventually it should be named userdata / userdatabase, while the current package should be named usermodel or something like that

var (
	Engine                       *xorm.Engine
	UserDataTableNameToInterface = map[string]interface{}{}
)

func AddTable(tableName string, structure interface{}) {
	_, exist := UserDataTableNameToInterface[tableName]
	if exist {
		log.Panic("table name already used: " + tableName)
	}
	UserDataTableNameToInterface[tableName] = structure
}

func InitTable(session *xorm.Session, tableName string, structure interface{}) {
	exist, err := session.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)

	if !exist {
		log.Println("Creating new table:", tableName)
		err = session.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
	}
}

func InitTables(engine *xorm.Engine) {
	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	for tableName, structure := range UserDataTableNameToInterface {
		InitTable(session, tableName, structure)
	}
	session.Commit()
}
