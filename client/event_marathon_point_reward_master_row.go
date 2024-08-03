package client

type EventMarathonPointRewardMasterRow struct {
	RequiredPoint int32 `json:"required_point" xorm:"required_point"`
	RewardGroupId int32 `json:"reward_group_id" xorm:"reward_group_id"`
}
