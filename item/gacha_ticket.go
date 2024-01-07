package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	NormalScoutingTicket = client.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     9000,
		ContentAmount: 1,
	}
	SRScoutingTicket = client.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     9002,
		ContentAmount: 1,
	}
	URScoutingTicket = client.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     9015,
		ContentAmount: 1,
	}
)
