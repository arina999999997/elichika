package client

import (
	"elichika/generic"
)

type BootstrapBanner struct {
	Banners generic.Array[Banner1] `json:"banners"`
}
