package user_live

import (
	"elichika/userdata"
)

// return the daily master id based on the live id of today
// if this is not available today, return nil
func GetLiveDailyMasterId(session *userdata.Session, liveId int32) *int32 {
	weekday := int32(session.Time.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	for _, liveDaily := range session.Gamedata.Live[liveId].LiveDailies {
		if liveDaily.Weekday == weekday {
			return &liveDaily.Id
		}
	}
	return nil
}
