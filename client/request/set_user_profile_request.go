package request

import (
	"elichika/generic"
)

type SetUserProfileRequest struct {
	Name        generic.Nullable[string] `json:"name"`         // pointer
	Nickname    generic.Nullable[string] `json:"nickname"`     // pointer
	Message     generic.Nullable[string] `json:"message"`      // pointer
	DeviceToken generic.Nullable[string] `json:"device_token"` // pointer
}
