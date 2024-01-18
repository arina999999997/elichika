package response

import (
	"elichika/client"
	"elichika/generic"
)

type OpenMemberLovePanelResponse struct {
	MemberLovePanels generic.Array[client.MemberLovePanel] `json:"member_love_panels"`
	UserModel        *client.UserModel                     `json:"user_model"`
}
