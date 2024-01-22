package client

import (
	"elichika/enum"

	"fmt"
)

type UserStoryEventUnlockItem struct {
	StoryEventUnlockItemMasterId int32 `json:"story_event_unlock_item_master_id"`
	Amount                       int32 `json:"amount"`
}

func (useui *UserStoryEventUnlockItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeStoryEventUnlock { // 30
		panic(fmt.Sprintln("Wrong content for StoryEventUnlockItem: ", content))
	}
	useui.StoryEventUnlockItemMasterId = content.ContentId
	useui.Amount = content.ContentAmount
}
func (useui *UserStoryEventUnlockItem) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeStoryEventUnlock,
		ContentId:     contentId,
		ContentAmount: useui.Amount,
	}
}
