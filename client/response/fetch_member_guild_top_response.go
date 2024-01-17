package response

import (
	"elichika/client"
)

type FetchMemberGuildTopResponse struct {
	MemberGuildTopStatus client.MemberGuildTopStatus `json:"member_guild_top_status"`
	UserModelDiff        *client.UserModel           `json:"user_model_diff"`
}
