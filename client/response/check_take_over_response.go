package response

import (
	"elichika/client"
)

type CheckTakeOverResponse struct {
	LinkedData    client.UserLinkData    `json:"linked_data"`
	CurrentData   client.CurrentUserData `json:"current_data"`
	IsNotTakeOver bool                   `json:"is_not_take_over"`
}
