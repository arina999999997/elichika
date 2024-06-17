package user_live

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserLiveDaily(session *userdata.Session, liveDailyMasterId int32) client.LiveDaily {
	result := client.LiveDaily{}

	_, err := session.Db.Table("u_live_daily").Where("user_id = ? AND live_daily_master_id = ?",
		session.UserId, liveDailyMasterId).Get(&result)
	utils.CheckErr(err)
	liveDailySetting := session.Gamedata.LiveDaily[liveDailyMasterId]
	if result.EndAt <= session.Time.Unix() { // reset or not exist in database
		return client.LiveDaily{
			LiveDailyMasterId:      liveDailyMasterId,
			LiveMasterId:           liveDailySetting.LiveId,
			EndAt:                  utils.BeginOfNextDay(session.Time).Unix(),
			RemainingPlayCount:     liveDailySetting.LimitCount,
			RemainingRecoveryCount: generic.NewNullable(liveDailySetting.MaxLimitCountRecover),
		}
	} else {
		result.LiveMasterId = liveDailySetting.LiveId
	}
	return result
}
