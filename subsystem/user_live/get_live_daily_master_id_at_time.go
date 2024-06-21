package user_live

import (
	"elichika/userdata"

	"time"
)

// return the daily master id based on the live id at the provided time
// it could be that the provided time doesn't have the daily master id, so we go back until we actually find something
// return nil if nothing can be found
// if this is not available then, return nil
func GetLiveDailyMasterIdAtTime(session *userdata.Session, liveId int32, unix int64) *int32 {
	weekday := int32(time.Unix(unix, 0).Weekday())
	if weekday == 0 {
		weekday = 7
	}
	for i := 0; i < 7; i++ { // retry at most 7 times
		for _, liveDaily := range session.Gamedata.Live[liveId].LiveDailies {
			if liveDaily.Weekday == weekday {
				return &liveDaily.Id
			}
		}
		weekday--
		if weekday == 0 {
			weekday = 7
		}
	}
	return nil
}
