package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchEmblemResponse struct {
	EmblemIsNewDataList generic.List[client.EmblemIsNewData] `json:"emblem_is_new_data_list"`
	UserModel           *client.UserModel                    `json:"user_model"`
}
