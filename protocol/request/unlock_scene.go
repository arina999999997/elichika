package request

type SaveUnlockedSceneRequest struct {
	UnlockSceneTypes []int32 `json:"unlock_scene_types"`
}
