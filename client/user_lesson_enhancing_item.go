package client

import (
	"elichika/enum"
	"elichika/log"

	"fmt"
)

type UserLessonEnhancingItem struct {
	EnhancingItemId int32 `json:"enhancing_item_id"`
	Amount          int32 `json:"amount"`
}

func (ulei *UserLessonEnhancingItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeLessonEnhancingItem { // 6
		log.Panic(fmt.Sprintln("Wrong content for LessonEnhancingItem: ", content))
	}
	ulei.EnhancingItemId = content.ContentId
	ulei.Amount = content.ContentAmount
}
func (ulei *UserLessonEnhancingItem) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeLessonEnhancingItem,
		ContentId:     contentId,
		ContentAmount: ulei.Amount,
	}
}
