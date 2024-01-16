package client

import (
	"elichika/enum"

	"fmt"
)

// TODO(refactor): Remove the id field once we rewrite the map
// also pretty sure this is broken
type UserExchangeEventPoint struct {
	PointId int32 `json:"-"`
	Amount  int32 `json:"amount"`
}

func (ueep *UserExchangeEventPoint) Id() int64 {
	return int64(ueep.PointId)
}
func (ueep *UserExchangeEventPoint) SetId(id int64) {
	ueep.PointId = int32(id)
}
func (ueep *UserExchangeEventPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeExchangeEventPoint { // 21
		panic(fmt.Sprintln("Wrong content for ExchangeEventPoint: ", content))
	}
	ueep.PointId = content.ContentId
	ueep.Amount = content.ContentAmount
}
func (ueep *UserExchangeEventPoint) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeExchangeEventPoint,
		ContentId:     ueep.PointId,
		ContentAmount: ueep.Amount,
	}
}
