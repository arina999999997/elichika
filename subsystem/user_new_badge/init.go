package user_new_badge

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata/database"
)

func init() {
	database.AddTable("u_new_badge", generic.UserIdWrapper[client.BootstrapNewBadge]{})
}
