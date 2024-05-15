package response

import (
	"elichika/client"
	"elichika/generic"
)

type UserSearchResponse struct {
	UserSearchList generic.Array[client.OtherUser] `json:"user_search_list"`
}
