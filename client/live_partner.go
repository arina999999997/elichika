package client

import (
	"elichika/generic"
)

type LivePartner struct {
	UserId              int32                                    `json:"user_id"`
	Name                LocalizedText                            `json:"name"`
	Rank                int32                                    `json:"rank"`
	LastLoginAt         int64                                    `json:"last_login_at"`
	CardByCategory      generic.Dictionary[int32, OtherUserCard] `xorm:"-" json:"card_by_category"`
	EmblemId            int32                                    `json:"emblem_id"`
	IsFriend            bool                                     `xorm:"-" json:"is_friend"`
	IntroductionMessage LocalizedText                            `xorm:"'message'" json:"introduction_message"`
}
