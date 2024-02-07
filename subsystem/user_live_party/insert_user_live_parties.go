package user_live_party

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertUserLiveParties(session *userdata.Session, parties []client.UserLiveParty) {
	for _, party := range parties {
		UpdateUserLiveParty(session, party)
	}
}
