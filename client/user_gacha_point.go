package client

import (
	"elichika/enum"

	"fmt"
)

type UserGachaPoint struct {
	PointMasterId int32 `json:"point_master_id"`
	Amount        int32 `json:"amount"`
}

func (ugp *UserGachaPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGachaPoint { // 5
		panic(fmt.Sprintln("Wrong content for GachaPoint: ", content))
	}
	ugp.PointMasterId = content.ContentId
	ugp.Amount = content.ContentAmount
}
func (ugp *UserGachaPoint) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeGachaPoint,
		ContentId:     contentId,
		ContentAmount: ugp.Amount,
	}
}
