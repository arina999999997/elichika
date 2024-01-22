package database

import (
	"elichika/generic"
)

type UserLoginBonus struct {
	LoginBonusId       int32 `xorm:"pk"`
	LastReceivedReward int
	LastReceivedAt     int64
}

func init() {
	AddTable("u_login_bonus", generic.UserIdWrapper[UserLoginBonus]{})
}
