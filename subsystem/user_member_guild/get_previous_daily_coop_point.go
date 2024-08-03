package user_member_guild

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

// if the data exist then return it, otherwise calculate
// the calculation is only correct if we assume that nothing has been calculated before
// TODO(database): This do not auto delete old data
func GetPreviousDailyCoopPoint(session *userdata.Session, resetAt int64) int32 {
	memberMasterId := session.UserStatus.MemberGuildMemberMasterId.Value
	totalPoint := int32(0)
	exist, err := session.Db.Table("u_member_guild_daily_coop_point").Where("member_master_id = ? AND reset_at = ?",
		memberMasterId, resetAt).Cols("total_point").Get(&totalPoint)
	utils.CheckErr(err)
	if exist {
		return totalPoint
	}

	var totals []int64
	totals, err = session.Db.Table("u_member_guild").
		Where("member_master_id = ? AND daily_support_point_reset_at = ?",
			memberMasterId, resetAt).SumsInt(&client.UserMemberGuild{}, "daily_support_point", "daily_love_point")
	utils.CheckErr(err)
	totalPoint = int32(totals[0] + totals[1])
	_, err = session.Db.Table("u_member_guild_daily_coop_point").Insert(database.UserMemberGuildDailyCoopPoint{
		MemberMasterId: memberMasterId,
		ResetAt:        resetAt,
		TotalPoint:     totalPoint,
	})
	utils.CheckErr(err)

	return totalPoint
}
