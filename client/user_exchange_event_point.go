package client

import (
	"elichika/enum"

	"fmt"
)

type UserExchangeEventPoint struct {
	Amount int32 `json:"amount"`
}

func (ueep *UserExchangeEventPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeExchangeEventPoint { // 21
		panic(fmt.Sprintln("Wrong content for ExchangeEventPoint: ", content))
	}
	ueep.Amount = content.ContentAmount
}
func (ueep *UserExchangeEventPoint) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeExchangeEventPoint,
		ContentId:     contentId,
		ContentAmount: ueep.Amount,
	}
}
