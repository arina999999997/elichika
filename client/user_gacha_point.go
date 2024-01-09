package client

import (
	"elichika/enum"

	"fmt"
)

type UserGachaPoint struct {
	PointMasterId int32 `json:"point_master_id"`
	Amount        int32 `json:"amount"`
}

func (ugp *UserGachaPoint) Id() int64 {
	return int64(ugp.PointMasterId)
}
func (ugp *UserGachaPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGachaPoint { // 5
		panic(fmt.Sprintln("Wrong content for GachaPoint: ", content))
	}
	ugp.PointMasterId = content.ContentId
	ugp.Amount = content.ContentAmount
}
func (ugp *UserGachaPoint) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeGachaPoint,
		ContentId:     ugp.PointMasterId,
		ContentAmount: ugp.Amount,
	}
}
