package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	SkipTicket = client.Content{
		ContentType:   enum.ContentTypeLiveSkipTicket,
		ContentId:     16001,
		ContentAmount: 1,
	}
)
