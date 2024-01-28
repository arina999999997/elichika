package user_suit

import (
	"elichika/client"
	"elichika/userdata"
)

// suit are inserted when the function is called as suit is unique and doesn't change

func InsertUserSuits(session *userdata.Session, suits []client.UserSuit) {
	for _, suit := range suits {
		session.UserModel.UserSuitBySuitId.Set(suit.SuitMasterId, suit)
	}
}

func InsertUserSuit(session *userdata.Session, suitMasterId int32) {
	suit := client.UserSuit{
		SuitMasterId: suitMasterId,
		IsNew:        true,
	}
	session.UserModel.UserSuitBySuitId.Set(suit.SuitMasterId, suit)
}
