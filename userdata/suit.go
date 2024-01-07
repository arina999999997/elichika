package userdata

import (
	"elichika/model"
	"elichika/utils"
)

// suit are inserted when the function is called as suit is unique and doesn't change

func (session *Session) InsertUserSuits(suits []model.UserSuit) {
	session.UserModel.UserSuitBySuitId.Objects = append(session.UserModel.UserSuitBySuitId.Objects, suits...)
}

func (session *Session) InsertUserSuit(suitMasterId int) {
	suit := model.UserSuit{
		SuitMasterId: suitMasterId,
		IsNew:        true,
	}
	session.UserModel.UserSuitBySuitId.PushBack(suit)
}

func suitFinalizer(session *Session) {
	for _, suit := range session.UserModel.UserSuitBySuitId.Objects {
		exist, err := session.Db.Table("u_suit").Exist(&suit)
		utils.CheckErr(err)
		if !exist {
			genericDatabaseInsert(session, "u_suit", suit)
		}
	}
}

func init() {
	addFinalizer(suitFinalizer)
	addGenericTableFieldPopulator("u_suit", "UserSuitBySuitId")
}
