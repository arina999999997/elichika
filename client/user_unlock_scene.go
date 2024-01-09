package client

type UserUnlockScene struct {
	UnlockSceneType int32 `xorm:"pk 'unlock_scene_type'" json:"unlock_scene_type"`
	Status          int32 `xorm:"'status'" json:"status" enum:"UnlockSceneStatus"`
}

func (uus *UserUnlockScene) Id() int64 {
	return int64(uus.UnlockSceneType)
}
