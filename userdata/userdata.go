package userdata

import (
	"elichika/config"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
)

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
	for tableName, interf := range model.TableNameToInterface {
		InitTable(tableName, interf, overwrite)
	}
	// TODO: redesign this to not store merged data, maybe

	InitTable("u_resource", UserResource{}, overwrite)
}

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", config.UserdataDbPath)
	utils.CheckErr(err)
	Engine.SetMaxOpenConns(50)
	Engine.SetMaxIdleConns(10)
	InitTables(false) // insert new tables if necessary
}
