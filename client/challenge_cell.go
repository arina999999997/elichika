package client

type ChallengeCell struct {
	CellId           int32 `xorm:"pk" json:"cell_id"`
	IsRewardReceived bool  `json:"is_reward_received"`
	Progress         int32 `json:"progress"`
}
