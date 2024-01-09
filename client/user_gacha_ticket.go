package client

import (
	"elichika/enum"

	"fmt"
)

// normally this would need its own table for the specific amounts
// but we just combine everything into normal amount because there's no use for other amount anyway
type UserGachaTicket struct {
	TicketMasterId int32 `json:"ticket_master_id"`
	NormalAmount   int32 `json:"normal_amount"`
	AppleAmount    int32 `json:"apple_amount"`
	GoogleAmount   int32 `json:"google_amount"`
}

func (ugt *UserGachaTicket) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGachaTicket { // 9
		panic(fmt.Sprintln("Wrong content for GachaTicket: ", content))
	}
	ugt.TicketMasterId = content.ContentId
	ugt.NormalAmount = content.ContentAmount
	ugt.AppleAmount = 0
	ugt.GoogleAmount = 0
}
func (ugt *UserGachaTicket) Id() int64 {
	return int64(ugt.TicketMasterId)
}
func (ugt *UserGachaTicket) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     ugt.TicketMasterId,
		ContentAmount: ugt.NormalAmount + ugt.AppleAmount + ugt.GoogleAmount,
	}
}
