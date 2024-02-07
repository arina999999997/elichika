package user_member

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertMembers(session *userdata.Session, members []client.UserMember) {
	for _, member := range members {
		UpdateMember(session, member)
	}
}
