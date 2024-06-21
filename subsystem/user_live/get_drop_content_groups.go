package user_live

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/userdata"
)

func getDropContentGroups(session *userdata.Session, liveScore *client.LiveScore, liveDifficulty *gamedata.LiveDifficulty) (*gamedata.LiveDropGroup, *gamedata.LiveDropGroup) {
	if liveDifficulty.UnlockPattern != enum.LiveUnlockPatternDaily {
		// not a daily song
		return liveDifficulty.DropContentGroup, liveDifficulty.RareDropContentGroup
	} else {
		// daily song can have their drop group overwritten depending on the day
		// this create
		var timeStamp int64
		if liveScore != nil {
			timeStamp = liveScore.StartInfo.CreatedAt
		} else { // skip ticket
			timeStamp = session.Time.Unix()
		}
		liveDailyMasterId := GetLiveDailyMasterIdAtTime(session, liveDifficulty.Live.LiveId, timeStamp)
		liveDaily := session.Gamedata.LiveDaily[*liveDailyMasterId]
		return liveDaily.DropContentGroup, liveDaily.RareDropContentGroup
	}
}
