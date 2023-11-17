package userdata

import (
	"elichika/model"
	"elichika/utils"

	"fmt"
)

func (session *Session) GetAllSuits() []model.UserSuit {
	suits := []model.UserSuit{}
	err := session.Db.Table("u_suit").Where("user_id = ?", session.UserStatus.UserID).Find(&suits)
	utils.CheckErr(err)
	return suits
}

func (session *Session) InsertUserSuits(suits []model.UserSuit) {
	if len(suits) == 0 {
		return
	}
	count, err := session.Db.Table("u_suit").AllCols().Insert(suits)
	utils.CheckErr(err)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted ", count, "suits")
}

func (session *Session) InsertUserSuit(suitMasterID int) {
	session.UserSuitDiffs = append(session.UserSuitDiffs,
		model.UserSuit{
			UserID:       session.UserStatus.UserID,
			SuitMasterID: suitMasterID,
			IsNew:        true,
		})
}

func (session *Session) FinalizeUserSuitDiffs() []any {
	session.InsertUserSuits(session.UserSuitDiffs)
	diffs := []any{}
	for _, suit := range session.UserSuitDiffs {
		session.UserModel.UserSuitBySuitID.PushBack(suit)
		diffs = append(diffs, suit.SuitMasterID)
		diffs = append(diffs, suit)
	}
	return diffs
}

func init() {
	addGenericTableFieldPopulator("u_suit", "UserSuitBySuitID")
}
