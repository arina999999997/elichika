package user_present

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata/database"
)

// Present require an unique id that is of type int32, and the network data from official server indicate that they store stat and just count up
type UserPresentStat struct {
	UserId       int32 `xorm:"pk"`
	PresentCount int32
}

func init() {
	database.AddTable("u_present", generic.UserIdWrapper[client.PresentItem]{})
	database.AddTable("u_present_stat", UserPresentStat{})
}
