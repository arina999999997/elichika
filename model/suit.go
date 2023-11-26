package model

type UserSuit struct {
	UserID       int  `xorm:"pk 'user_id'" json:"-"`
	SuitMasterID int  `xorm:"pk 'suit_master_id'" json:"suit_master_id"`
	IsNew        bool `xorm:"'is_new'" json:"is_new"`
}

func (us *UserSuit) ID() int64 {
	return int64(us.SuitMasterID)
}
func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_suit"] = UserSuit{}
}
