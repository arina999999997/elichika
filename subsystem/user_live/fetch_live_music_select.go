package user_live

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/generic"
	"elichika/userdata"

	"time"
)

func FetchLiveMusicSelect(session *userdata.Session) response.FetchLiveMusicSelectResponse {
	year, month, day := session.Time.Year(), session.Time.Month(), session.Time.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, session.Time.Location()).Unix()

	weekday := int32(session.Time.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	resp := response.FetchLiveMusicSelectResponse{
		WeekdayState: client.WeekdayState{
			Weekday:       weekday,
			NextWeekdayAt: tomorrow,
		},
		UserModelDiff: &session.UserModel,
	}

	// TODO(live): Keep track of daily song play(?)
	for _, liveDaily := range session.Gamedata.LiveDaily {
		if liveDaily.Weekday != weekday {
			continue
		}
		resp.LiveDailyList.Append(client.LiveDaily{
			LiveDailyMasterId:      liveDaily.Id,
			LiveMasterId:           liveDaily.LiveId,
			EndAt:                  tomorrow,
			RemainingPlayCount:     5,
			RemainingRecoveryCount: generic.NewNullable(int32(10)),
		})
	}

	return resp
}
