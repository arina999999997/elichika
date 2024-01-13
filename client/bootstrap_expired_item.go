package client

import (
	"elichika/generic"
)

type BootstrapExpiredItem struct {
	ExpiredItems generic.Array[Content] `json:"expired_items"`
}
