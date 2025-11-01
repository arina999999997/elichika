package client

type EventMiningRewardMasterRow struct {
	RewardGroupId int32   `xorm:"'reward_group_id'" json:"reward_group_id"`
	RewardContent Content `xorm:"extends" json:"reward_content"`
	DisplayOrder  int32   `xorm:"'display_order'" json:"display_order"`
}
