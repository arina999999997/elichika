package database

import (
	"elichika/generic"
)

type UserCardPlayCountStat struct {
	CardMasterId         int32 `xorm:"pk"`
	LiveJoinCount        int32
	ActiveSkillPlayCount int32
}

func init() {
	AddTable("u_card_play_count_stat", generic.UserIdWrapper[UserCardPlayCountStat]{})
}
