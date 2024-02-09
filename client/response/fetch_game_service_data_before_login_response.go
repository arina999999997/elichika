package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchGameServiceDataBeforeLoginResponse struct {
	Data generic.Nullable[client.UserLinkDataBeforeLogin] `json:"data"`
}
