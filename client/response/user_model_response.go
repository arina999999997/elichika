package response

import (
	"elichika/client"
)

type UserModelResponse struct {
	UserModel *client.UserModel `json:"user_model"`
}
