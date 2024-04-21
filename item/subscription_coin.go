package item

import (
	"elichika/client"
	"elichika/enum"
)

var (
	MemberCoin = client.Content{
		ContentType:   enum.ContentTypeSubscriptionCoin,
		ContentId:     0,
		ContentAmount: 1,
	}
)
