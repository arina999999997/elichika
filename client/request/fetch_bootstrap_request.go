package request

type FetchBootstrapRequest struct {
	BootstrapFetchTypes []int32 `json:"bootstrap_fetch_types" enum:"BootstrapFetchType"`
	DeviceToken         string  `json:"device_token"`
	DeviceName          string  `json:"device_name"`
}
