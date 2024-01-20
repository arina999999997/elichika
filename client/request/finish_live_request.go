package request

import (
	"elichika/client"
)

type FinishLiveRequest struct {
	LiveId           int64                   `json:"live_id"`
	LiveFinishStatus int32                   `json:"live_finish_status" enum:"LiveFinishStatus"`
	LiveScore        client.LiveScore        `json:"live_score"`
	ResumeFinishInfo client.ResumeFinishInfo `json:"resume_finish_info"`
	RoomId           int32                   `json:"room_id"`
}
