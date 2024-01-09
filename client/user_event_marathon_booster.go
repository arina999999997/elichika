package client

import (
	"elichika/enum"

	"fmt"
)

type UserEventMarathonBooster struct {
	EventItemId int32 `json:"event_item_id"`
	Amount      int32 `json:"amount"`
}

func (uemb *UserEventMarathonBooster) Id() int64 {
	return int64(uemb.EventItemId)
}
func (uemb *UserEventMarathonBooster) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeEventMarathonBooster { // 27
		panic(fmt.Sprintln("Wrong content for EventMarathonBooster: ", content))
	}
	uemb.EventItemId = content.ContentId
	uemb.Amount = content.ContentAmount
}
func (uemb *UserEventMarathonBooster) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeEventMarathonBooster,
		ContentId:     uemb.EventItemId,
		ContentAmount: uemb.Amount,
	}
}
