package user_member_guild

import (
	"elichika/userdata"

	"fmt"
)

func IsMemberGuildRankingPeriod(session *userdata.Session) bool {
	period := session.Gamedata.MemberGuildPeriod
	currentPeriodTime := session.Time.Unix() - GetCurrentMemberGuildRankingPeriodStart(session)
	fmt.Println(currentPeriodTime)
	return (currentPeriodTime >= period.RankingStartSecs) && (currentPeriodTime < period.RankingEndSecs)
}
