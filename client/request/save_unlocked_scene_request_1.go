package request

import (
	"elichika/generic"
)

type SaveUnlockedSceneRequest1 struct {
	UnlockSceneTypes generic.Array[int32] `json:"unlock_scene_types" enum:"UnlockSceneType"`
}
