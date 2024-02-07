package user_live_party

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserLiveParty(session *userdata.Session, partyId int) client.UserLiveParty {
	liveParty := client.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND party_id = ?", session.UserId, partyId).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		panic("Party doesn't exist")
	}
	return liveParty
}
