package response

import (
	"elichika/client"
)

type FetchOtherUserCardResponse struct {
	OtherUserCard client.OtherUserCard `json:"other_user_card"`
}
