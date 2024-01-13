package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchCommunicationMemberDetailResponse struct {
	MemberLovePanels generic.Array[client.MemberLovePanel] `json:"member_love_panels"`
	WeekdayState     client.WeekdayState                   `json:"weekday_state"`
}
