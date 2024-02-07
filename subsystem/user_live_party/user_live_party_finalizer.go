package user_live_party

import (
	"elichika/userdata"
	"elichika/utils"
)

func userLivePartyFinalizer(session *userdata.Session) {
	for _, party := range session.UserModel.UserLivePartyById.Map {
		affected, err := session.Db.Table("u_live_party").
			Where("user_id = ? AND party_id = ?", session.UserId, party.PartyId).AllCols().
			Update(*party)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_live_party", *party)
		}
	}
}
func init() {
	userdata.AddFinalizer(userLivePartyFinalizer)
}
