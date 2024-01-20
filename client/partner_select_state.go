package client

import (
	"elichika/generic"
)

type PartnerSelectState struct {
	LivePartners generic.Array[LivePartner] `json:"live_partners"`
	FriendCount  int32                      `json:"friend_count"`
}
