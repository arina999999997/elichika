package client

import (
	"elichika/enum"
	"elichika/log"

	"fmt"
)

type UserGradeUpItem struct {
	ItemMasterId int32 `json:"item_master_id"`
	Amount       int32 `json:"amount"`
}

func (ugui *UserGradeUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGradeUpper { // 13
		log.Panic(fmt.Sprintln("Wrong content for GradeUpItem: ", content))
	}
	ugui.ItemMasterId = content.ContentId
	ugui.Amount = content.ContentAmount
}
func (ugui *UserGradeUpItem) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeGradeUpper,
		ContentId:     contentId,
		ContentAmount: ugui.Amount,
	}
}
