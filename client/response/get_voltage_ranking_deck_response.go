package response

import (
	"elichika/client"
)

type GetVoltageRankingDeckResponse struct {
	User       client.OtherUser           `xorm:"-" json:"user"`
	DeckDetail client.OtherUserDeckDetail `xorm:"json" json:"deck_detail"`
}
