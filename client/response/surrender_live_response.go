package response

import (
	"elichika/client"
	"elichika/generic"
)

type SurrenderLiveResponse struct {
	LpDiff        generic.Nullable[int32] `json:"lp_diff"`
	UserModelDiff *client.UserModel       `json:"user_model_diff"`
}
