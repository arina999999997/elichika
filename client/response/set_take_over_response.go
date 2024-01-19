package response

import (
	"elichika/client"
)

type SetTakeOverResponse struct {
	Data client.UserLinkData `json:"data"`
}
