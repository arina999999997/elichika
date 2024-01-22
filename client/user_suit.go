package client

type UserSuit struct {
	SuitMasterId int32 `xorm:"pk 'suit_master_id'" json:"suit_master_id"`
	IsNew        bool  `xorm:"'is_new'" json:"is_new"`
}
