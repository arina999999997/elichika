package userdata

import (
	"elichika/userdata/database"
	"elichika/utils"

	"fmt"
	"os"
	"reflect"

	"xorm.io/xorm"
)

// export to .db one user

func (session *Session) ExportDb() []byte {
	fileName := fmt.Sprintf("backup_%d.db", session.UserId)
	os.Remove(fileName) // remove dirty data if any
	engine, err := xorm.NewEngine("sqlite", fileName)
	defer engine.Close()
	utils.CheckErr(err)
	defer func() {
		err := os.Remove(fileName)
		utils.CheckErr(err)
	}()
	engine.SetMaxOpenConns(1)
	engine.SetMaxIdleConns(1)
	database.InitTables(engine)

	writeSession := engine.NewSession()
	defer writeSession.Close()
	err = writeSession.Begin()
	utils.CheckErr(err)
	for table, inter := range database.UserDataTableNameToInterface {
		if table == "u_friend_status" {
			// do not extract this info as it is server side only, and importing friend connection lead to bad data
			continue
		}
		rows := reflect.New(reflect.SliceOf(reflect.TypeOf(inter)))
		err := session.Db.Table(table).Where("user_id = ?", session.UserId).Find(rows.Interface())
		utils.CheckErr(err)
		// we have to insert one by one, otherwise we mgiht get too many variable
		sz := reflect.Indirect(rows).Len()
		for i := 0; i < sz; i++ {
			_, err := writeSession.Table(table).Insert(reflect.Indirect(rows).Index(i).Interface())
			utils.CheckErr(err)
		}
	}
	writeSession.Commit()
	writeSession.Close()
	engine.Close()
	bytes, err := os.ReadFile(fileName)
	utils.CheckErr(err)
	return bytes
}
