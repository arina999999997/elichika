package client

import (
	"elichika/enum"

	"fmt"
)

type UserLiveSkipTicket struct {
	TicketMasterId int32 `json:"ticket_master_id"`
	Amount         int32 `json:"amount"`
}

func (ulst *UserLiveSkipTicket) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeLiveSkipTicket { // 28
		panic(fmt.Sprintln("Wrong content for LiveSkipTicket: ", content))
	}
	ulst.TicketMasterId = content.ContentId
	ulst.Amount = content.ContentAmount
}
func (ulst *UserLiveSkipTicket) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeLiveSkipTicket,
		ContentId:     contentId,
		ContentAmount: ulst.Amount,
	}
}
