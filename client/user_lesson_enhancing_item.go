package client

import (
	"elichika/enum"

	"fmt"
)

type UserLessonEnhancingItem struct {
	EnhancingItemId int32 `json:"enhancing_item_id"`
	Amount          int32 `json:"amount"`
}

func (ulei *UserLessonEnhancingItem) Id() int64 {
	return int64(ulei.EnhancingItemId)
}
func (ulei *UserLessonEnhancingItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeLessonEnhancingItem { // 6
		panic(fmt.Sprintln("Wrong content for LessonEnhancingItem: ", content))
	}
	ulei.EnhancingItemId = content.ContentId
	ulei.Amount = content.ContentAmount
}
func (ulei *UserLessonEnhancingItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeLessonEnhancingItem,
		ContentId:     ulei.EnhancingItemId,
		ContentAmount: ulei.Amount,
	}
}
