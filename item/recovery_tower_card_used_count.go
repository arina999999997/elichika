package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	PerformanceDrink = client.Content{
		ContentType:   enum.ContentTypeRecoveryTowerCardUsedCount,
		ContentId:     24001,
		ContentAmount: 1,
	}
)
