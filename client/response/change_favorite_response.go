package response

import (
	"elichika/client"
)

type ChangeFavoriteResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
