package response

import (
	"elichika/client"
)

type UpdateCardNewFlagResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
