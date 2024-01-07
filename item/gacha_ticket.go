package item

import (
	"elichika/enum"
	"elichika/model"
)

var (
	NormalScoutingTicket = model.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     9000,
		ContentAmount: 1,
	}
	SRScoutingTicket = model.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     9002,
		ContentAmount: 1,
	}
	URScoutingTicket = model.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     9015,
		ContentAmount: 1,
	}
)
