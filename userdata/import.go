package userdata

import (
	"elichika/client/response"
	"elichika/userdata/database"
	"elichika/utils"

	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func (session *Session) ImportLoginData(ctx *gin.Context, loginData *response.LoginResponse) {
	session.SessionType = SessionTypeImportAccount
	if session.Ctx != ctx {
		panic("context changed")
	}
	session.UserModel = *loginData.UserModel
	session.UserStatus = &session.UserModel.UserStatus
	session.MemberLovePanels = loginData.MemberLovePanels.Slice
}

func (session *Session) ImportDatabaseData(ctx *gin.Context, bytes []byte) (*string, *string) {
	resp := new(string)
	session.SessionType = SessionTypeDirectDbWrite
	if session.Ctx != ctx {
		panic("context changed")
	}
	fileName := fmt.Sprintf("backup_%d.db", session.UserId)
	os.Remove(fileName) // remove dirty data if any
	err := os.WriteFile(fileName, bytes, 0644)
	utils.CheckErr(err)
	defer func() {
		os.Remove(fileName)
	}()
	engine, err := xorm.NewEngine("sqlite", fileName)
	utils.CheckErr(err)
	engine.SetMaxOpenConns(1)
	engine.SetMaxIdleConns(1)

	readSession := engine.NewSession()
	defer readSession.Close()
	err = readSession.Begin()
	utils.CheckErr(err)

	var userId *int32
	// allow importing from a separate user id to
	for table, inter := range database.UserDataTableNameToInterface {
		if table == "u_pass_word" || table == "u_authentication" || table == "u_friend_status" {
			continue
		}
		rows := reflect.New(reflect.SliceOf(reflect.TypeOf(inter)))
		err := readSession.Table(table).Find(rows.Interface())
		utils.CheckErr(err)
		// we have to insert one by one, otherwise we might get too many variable
		sz := reflect.Indirect(rows).Len()
		for i := 0; i < sz; i++ {
			item := reflect.Indirect(rows).Index(i)
			objectUserIdField := item.FieldByName("UserId")
			objectUserId := objectUserIdField.Interface().(int32)
			if userId == nil {
				userId = new(int32)
				*userId = objectUserId
			} else if objectUserId != *userId {
				*resp = "Error: uploaded database contain more than one user\nLoad it up with some other instance of elichika and extract only the user you want."
				return nil, resp
			}
			objectUserIdField.Set(reflect.ValueOf(session.UserId))

			_, err := session.Db.Table(table).Insert(item.Interface())
			utils.CheckErr(err)
		}
	}
	session.Finalize()
	*resp = fmt.Sprintf("Imported user, old id: %d, new id: %d.\nNote that the authentication is the same with existing data (logged into webui), not imported data.", *userId, session.UserId)
	return resp, nil
}
