package response

import (
	"elichika/client"
	"elichika/generic"
)

type MissionReceiveErrorResponse struct {
	MissionMasterIdList generic.List[int32] `json:"mission_master_id_list"`
	UserModel           *client.UserModel   `json:"user_model"`
	ExpiredMissionIds   generic.List[int32] `json:"expired_mission_ids"`
}
