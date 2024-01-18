package response

import (
	"elichika/client"
)

type UserProfileResponse struct {
	ProfileInfo client.ProfileInfomation       `json:"profile_info"`
	GuestInfo   client.ProfileGuestConfig      `json:"guest_info"`
	PlayInfo    client.ProfilePlayHistory      `json:"play_info"`
	MemberInfo  client.ProfileMemberInfomation `json:"member_info"`
}
