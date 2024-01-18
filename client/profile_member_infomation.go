package client

import (
	"elichika/generic"
)

type ProfileMemberInfomation struct {
	UserMembers generic.Array[ProfileUserMember] `json:"user_members"`
}
