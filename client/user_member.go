package client

type UserMember struct {
	MemberMasterId           int32 `xorm:"pk 'member_master_id'" json:"member_master_id"`
	CustomBackgroundMasterId int32 `xorm:"'custom_background_master_id'" json:"custom_background_master_id"`
	SuitMasterId             int32 `xorm:"'suit_master_id'" json:"suit_master_id"`
	LovePoint                int32 `json:"love_point"`
	LovePointLimit           int32 `json:"love_point_limit"`
	LoveLevel                int32 `json:"love_level"`
	ViewStatus               int32 `json:"view_status"`
	IsNew                    bool  `json:"is_new"`
}
