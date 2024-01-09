package client

import (
	"elichika/enum"

	"fmt"
)

type UserGradeUpItem struct {
	ItemMasterId int32 `json:"item_master_id"`
	Amount       int32 `json:"amount"`
}

func (ugui *UserGradeUpItem) Id() int64 {
	return int64(ugui.ItemMasterId)
}
func (ugui *UserGradeUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGradeUpper { // 13
		panic(fmt.Sprintln("Wrong content for GradeUpItem: ", content))
	}
	ugui.ItemMasterId = content.ContentId
	ugui.Amount = content.ContentAmount
}
func (ugui *UserGradeUpItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeGradeUpper,
		ContentId:     ugui.ItemMasterId,
		ContentAmount: ugui.Amount,
	}
}
