package client

import (
	"elichika/generic"
)

type LoginBonusRewards struct {
	Day                int32                             `json:"day"`
	Status             int32                             `json:"status" enum:"LoginBonusReceiveStatus"`
	ContentGrade       generic.Nullable[int32]           `xorm:"json 'content_grade'" json:"content_grade" enum:"LoginBonusContentGrade"` // can be 0
	LoginBonusContents generic.Array[LoginBonusContents] `json:"login_bonus_contents"`
}
