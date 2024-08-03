package user_member_guild

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetDailyCoopPoint(session *userdata.Session) int32 {
	memberMasterId := session.UserStatus.MemberGuildMemberMasterId.Value
	var totals []int64
	var err error
	totals, err = session.Db.Table("u_member_guild").
		Where("member_master_id = ? AND daily_support_point_reset_at = ?",
			memberMasterId, utils.NextMidDay(session.Time).Unix()).SumsInt(&client.UserMemberGuild{}, "daily_support_point", "daily_love_point")
	utils.CheckErr(err)
	return int32(totals[0] + totals[1])
}
