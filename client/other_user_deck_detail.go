package client

import (
	"elichika/generic"
)

type OtherUserDeckDetail struct {
	Deck             OtherUserDeck                    `json:"deck"`
	MemberLoveLevels generic.Dictionary[int32, int32] `json:"member_love_levels"`
}
