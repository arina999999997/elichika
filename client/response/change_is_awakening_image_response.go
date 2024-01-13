package response

import (
	"elichika/client"
)

type ChangeIsAwakeningImageResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
