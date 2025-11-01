package mining

import (
	"elichika/gamedata"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func TrackClearVoltage(session *userdata.Session, liveDifficulty *gamedata.LiveDifficulty, voltage int32) {
	eventMining := session.Gamedata.EventActive.GetEventMining()
	good := false
	for _, live := range eventMining.TopStatus.EventMiningCompetitionMasterRows.Slice {
		if live.LiveId == *liveDifficulty.LiveId {
			good = true
			break
		}
	}
	if !good {
		return
	}
	summary := GetUserVoltageRankingSummary(session, eventMining.EventId)
	replacedScore, added := summary.AddRecord(*liveDifficulty.LiveId, liveDifficulty.LiveDifficultyId, voltage)
	if !added {
		return
	}
	GetVoltageRanking(session.Db, eventMining.EventId).Update(session.UserId, summary.TotalScore())
	// added, replace the score
	if replacedScore.Score > 0 {
		_, err := session.Db.Table("u_event_mining_music_best_score").Where("user_id = ? AND live_id = ? AND live_difficulty_id = ? AND score = ?",
			session.UserId, replacedScore.LiveId, replacedScore.LiveDifficultyId, replacedScore.Score).Limit(1).Cols("live_id", "live_difficulty_id", "score").
			Update(database.UserEventMiningMusicBestScore{
				UserId:           session.UserId,
				LiveId:           *liveDifficulty.LiveId,
				LiveDifficultyId: liveDifficulty.LiveDifficultyId,
				Score:            voltage,
			})
		utils.CheckErr(err)
	} else { // added, new
		_, err := session.Db.Table("u_event_mining_music_best_score").Insert(database.UserEventMiningMusicBestScore{
			UserId:           session.UserId,
			LiveId:           *liveDifficulty.LiveId,
			LiveDifficultyId: liveDifficulty.LiveDifficultyId,
			Score:            voltage,
		})
		utils.CheckErr(err)
	}
}
