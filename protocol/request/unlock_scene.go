package request

type SaveUnlockedSceneRequest struct {
	UnlockSceneTypes []int `json:"unlock_scene_types"`
}
