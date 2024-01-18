package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchMissionResponse struct {
	MissionMasterIdList generic.List[int32] `json:"mission_master_id_list"`
	UserModel           *client.UserModel   `json:"user_model"`
}
