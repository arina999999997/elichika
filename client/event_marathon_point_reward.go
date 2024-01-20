package client

type EventMarathonPointReward struct {
	RewardContent     Content `json:"reward_content"`
	RequiredPoint     int32   `json:"required_point"`
	IsStartLoopReward bool    `json:"is_start_loop_reward"`
}
