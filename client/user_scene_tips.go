package client

type UserSceneTips struct {
	SceneTipsType int32 `xorm:"pk 'scene_tips_type'" json:"scene_tips_type" enum:"SceneTipsType"`
}

func (ust *UserSceneTips) Id() int64 {
	return int64(ust.SceneTipsType)
}
