package response

import (
	"elichika/client"
)

type StartLiveMvResponse struct {
	UniqId        int64             `json:"uniq_id"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
