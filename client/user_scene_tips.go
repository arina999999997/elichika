package client

type UserSceneTips struct {
	SceneTipsType int32 `xorm:"pk 'scene_tips_type'" json:"scene_tips_type" enum:"SceneTipsType"`
}
