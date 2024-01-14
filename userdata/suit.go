package userdata

import (
	"elichika/client"
	"elichika/utils"
)

// suit are inserted when the function is called as suit is unique and doesn't change

func (session *Session) InsertUserSuits(suits []client.UserSuit) {
	for _, suit := range suits {
		session.UserModel.UserSuitBySuitId.Set(suit.SuitMasterId, suit)
	}
}

func (session *Session) InsertUserSuit(suitMasterId int32) {
	suit := client.UserSuit{
		SuitMasterId: suitMasterId,
		IsNew:        true,
	}
	session.UserModel.UserSuitBySuitId.Set(suit.SuitMasterId, suit)
}

func suitFinalizer(session *Session) {
	for _, suit := range session.UserModel.UserSuitBySuitId.Map {
		affected, err := session.Db.Table("u_suit").
			Where("user_id = ? AND suit_master_id = ?", session.UserId, suit.SuitMasterId).
			Update(*suit)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_suit", *suit)
		}
	}
}

func init() {
	addFinalizer(suitFinalizer)
}
