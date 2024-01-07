package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	MemoryKey = client.Content{
		ContentType:   enum.ContentTypeStoryEventUnlock,
		ContentId:     17001,
		ContentAmount: 1,
	}
)
