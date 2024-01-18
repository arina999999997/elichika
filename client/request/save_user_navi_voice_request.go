package request

import (
	"elichika/generic"
)

type SaveUserNaviVoiceRequest struct {
	NaviVoiceMasterIds generic.Array[int32] `json:"navi_voice_master_ids"`
}
