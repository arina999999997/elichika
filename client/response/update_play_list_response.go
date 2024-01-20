package response

import (
	"elichika/client"
)

type UpdatePlayListResponse struct {
	IsSuccess     bool              `json:"is_success"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
