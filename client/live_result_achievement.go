package client

type LiveResultAchievement struct {
	Position            int32 `json:"position"`
	IsAlreadyAchieved   bool  `json:"is_already_achieved"`
	IsCurrentlyAchieved bool  `json:"is_currently_achieved"`
}
