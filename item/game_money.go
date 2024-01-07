package item

import (
	"elichika/enum"
	"elichika/model"
)

var (
	Gold = model.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentId:     0,
		ContentAmount: 1,
	}
)
