package response

import (
	"elichika/client"
	"elichika/generic"
)

type MissionReceiveResponse struct {
	MissionMasterIdList  generic.List[int32]          `json:"mission_master_id_list"`
	UserModel            *client.UserModel            `json:"user_model"`
	ReceivedPresentItems generic.List[client.Content] `json:"received_present_items"`
	LimitExceeded        bool                         `json:"limit_exceeded"`
}
