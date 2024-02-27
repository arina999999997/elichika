package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	Gold = client.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentId:     1200,
		ContentAmount: 1,
	}
)
