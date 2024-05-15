package user_social

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

func GetNullableOtherUser(session *userdata.Session, otherUserId int32) generic.Nullable[client.OtherUser] {
	if !UserExist(session, otherUserId) {
		return generic.Nullable[client.OtherUser]{}
	} else {
		return generic.NewNullable(GetOtherUser(session, otherUserId))
	}
}
