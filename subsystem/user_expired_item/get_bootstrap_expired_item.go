package user_expired_item

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

// TODO(present_box): Handle expired items
func GetBootstrapExpiredItem(session *userdata.Session) generic.Nullable[client.BootstrapExpiredItem] {
	return generic.Nullable[client.BootstrapExpiredItem]{}
}
