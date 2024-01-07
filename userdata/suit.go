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
		UserId:       session.UserStatus.UserId,
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
			_, err := session.Db.Table("u_suit").AllCols().Insert(suit)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(suitFinalizer)
	addGenericTableFieldPopulator("u_suit", "UserSuitBySuitId")
}
