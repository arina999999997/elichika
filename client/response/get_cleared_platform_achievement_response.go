package response

import (
	"elichika/generic"
)

type GetClearedPlatformAchievementResponse struct {
	ClearedIds generic.Array[string] `json:"cleared_ids"`
}
