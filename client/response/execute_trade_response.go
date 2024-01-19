package response

import (
	"elichika/client"
	"elichika/generic"
)

type ExecuteTradeResponse struct {
	Trades           generic.Array[client.Trade] `json:"trades"`
	IsSendPresentBox bool                        `json:"is_send_present_box"`
	UserModelDiff    *client.UserModel           `json:"user_model_diff"`
}
