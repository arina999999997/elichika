package request

import (
	"elichika/generic"
)

type FetchBootstrapRequest struct {
	BootstrapFetchTypes generic.Array[int32] `json:"bootstrap_fetch_types" enum:"BootstrapFetchType"`
	DeviceToken         string               `json:"device_token"`
	DeviceName          string               `json:"device_name"`
}
