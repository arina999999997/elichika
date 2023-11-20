package userdata

import (
	"elichika/model"
	"elichika/utils"

	"fmt"
)

// suit are inserted when the function is called as suit is unique and doesn't change

func (session *Session) InsertUserSuits(suits []model.UserSuit) {
	count, err := session.Db.Table("u_suit").AllCols().Insert(suits)
	utils.CheckErr(err)
	fmt.Println("Inserted ", count, "suit(s)")
}

func (session *Session) InsertUserSuit(suitMasterID int) {
	suit := model.UserSuit{
		UserID:       session.UserStatus.UserID,
		SuitMasterID: suitMasterID,
		IsNew:        true,
	}
	session.UserModel.UserSuitBySuitID.PushBack(suit)
	count, err := session.Db.Table("u_suit").AllCols().Insert(suit)
	utils.CheckErr(err)
	fmt.Println("Inserted ", count, "suit(s)")

}

func init() {
	addGenericTableFieldPopulator("u_suit", "UserSuitBySuitID")
}
