package model

type UserLoginBonus struct {
	UserId             int `xorm:"pk"`
	LoginBonusId       int `xorm:"pk"`
	LastReceivedReward int
	LastReceivedAt     int64
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_login_bonus"] = UserLoginBonus{}
}
