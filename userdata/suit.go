package userdata

import (
	"elichika/model"
	"elichika/utils"
)

// suit are inserted when the function is called as suit is unique and doesn't change

func (session *Session) InsertUserSuits(suits []model.UserSuit) {
	session.UserModel.UserSuitBySuitID.Objects = append(session.UserModel.UserSuitBySuitID.Objects, suits...)
}

func (session *Session) InsertUserSuit(suitMasterID int) {
	suit := model.UserSuit{
		UserID:       session.UserStatus.UserID,
		SuitMasterID: suitMasterID,
		IsNew:        true,
	}
	session.UserModel.UserSuitBySuitID.PushBack(suit)
}

func suitFinalizer(session *Session) {
	for _, suit := range session.UserModel.UserSuitBySuitID.Objects {
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
	addGenericTableFieldPopulator("u_suit", "UserSuitBySuitID")
}
