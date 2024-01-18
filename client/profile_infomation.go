package client

import (
	"elichika/generic"
)

type ProfileInfomation struct {
	BasicInfo                 OtherUser                      `json:"basic_info"`
	TotalLovePoint            int32                          `json:"total_love_point"`
	LoveMembers               generic.Array[MemberLovePoint] `json:"love_members"`
	MemberGuildMemberMasterId generic.Nullable[int32]        `json:"member_guild_member_master_id"`
}
