package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	TrainingTicket = client.Content{
		ContentType:   enum.ContentTypeRecoveryAp,
		ContentId:     2200,
		ContentAmount: 1,
	}
)
