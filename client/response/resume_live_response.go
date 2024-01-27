package response

import (
	"elichika/client"
)

type ResumeLiveResponse struct {
	Live          client.Live                `json:"live"`
	PartnerUserId int32                      `json:"partner_user_id"`
	IsAutoPlay    bool                       `json:"is_auto_play"`
	PickupInfo    client.BootstrapPickupInfo `json:"pickup_info"`
	WeekdayState  client.WeekdayState        `json:"weekday_state"`
}
