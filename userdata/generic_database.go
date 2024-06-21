package userdata

import (
	"elichika/utils"

	"reflect"
)

func hasUserId[T any](item T) bool {
	field := reflect.ValueOf(item).FieldByName("UserId")
	return field.IsValid()
}

func GenericDatabaseInsert[T any](session *Session, table string, item T) {
	if hasUserId(item) {
		_, err := session.Db.Table(table).Insert(item)
		utils.CheckErr(err)
	} else {
		type UserIdWrapper struct {
			UserId int32 `xorm:"pk 'user_id'"`
			Item   *T    `xorm:"extends"`
		}
		_, err := session.Db.Table(table).Insert(UserIdWrapper{
			UserId: session.UserId,
			Item:   &item,
		})
		utils.CheckErr(err)
	}
}

func GenericDatabaseExist[T any](session *Session, table string, item T) bool {
	if hasUserId(item) {
		exist, err := session.Db.Table(table).Exist(&item)
		utils.CheckErr(err)
		return exist
	} else {
		type UserIdWrapper struct {
			UserId int32 `xorm:"pk 'user_id'"`
			Item   *T    `xorm:"extends"`
		}
		exist, err := session.Db.Table(table).Exist(&UserIdWrapper{
			UserId: session.UserId,
			Item:   &item,
		})
		utils.CheckErr(err)
		return exist
	}
}
