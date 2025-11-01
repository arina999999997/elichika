package mining

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/generic/ranking"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"

	"xorm.io/xorm"
)

// this is how many score to accept for a song
// the amount of song is fixed to 3, but this amount can change
// it was initially 5 but later on it was changed to 3
// support value from 3 to 5
var voltageSelectionAmount int32 = 1

// not sure how the ranking actually work, we use the following model:
// - the score is divided by 9 first, then that result is ranked
//   - in practice this is doneat the get score level
//
// - it is possible that 2 players have the same divided score, but different sum:
//   - in this case, this implementation will give them the same rank
//   - probably too rare to ever happen though
type UserVoltageRankingSummary struct {
	BestScores [3][5]database.UserEventMiningMusicBestScore
}

func NewUserVoltageRankingSummary(eventId int32) UserVoltageRankingSummary {
	if eventId < 0 {
		eventId = gamedata.Instance.EventActive.GetEventMining().EventId
	}
	eventMining := gamedata.Instance.EventMining[eventId]
	if voltageSelectionAmount != eventMining.TopStatus.SelectionAmount {
		voltageSelectionAmount = eventMining.TopStatus.SelectionAmount
	}
	res := UserVoltageRankingSummary{}
	for songPos := 0; songPos < 3; songPos++ {
		for frameNo := int32(0); frameNo < voltageSelectionAmount; frameNo++ {
			res.BestScores[songPos][frameNo] = database.UserEventMiningMusicBestScore{
				LiveId:           eventMining.TopStatus.EventMiningCompetitionMasterRows.Slice[songPos].LiveId,
				LiveDifficultyId: 0,
				Score:            0,
			}
		}
	}
	return res
}

// add a new score and return the old replaced score, and return whether the score was added
func (s *UserVoltageRankingSummary) AddRecord(liveId, liveDifficultyId, score int32) (replaced database.UserEventMiningMusicBestScore, added bool) {
	added = false
	// it's fine to do this here
	songPos := 0
	for ; (songPos < 3) && (s.BestScores[songPos][0].LiveId != liveId); songPos++ {
	}
	if songPos == 3 { // not a song that is ranking
		return
	}
	if s.BestScores[songPos][voltageSelectionAmount-1].Score >= score { // new score is not high enough
		return
	}
	added = true
	replaced = s.BestScores[songPos][voltageSelectionAmount-1] // always replace the weakest score
	s.BestScores[songPos][voltageSelectionAmount-1].LiveDifficultyId = liveDifficultyId
	s.BestScores[songPos][voltageSelectionAmount-1].Score = score
	// keep the array sorted by decreasing Score
	for i := voltageSelectionAmount - 1; i >= 1; i-- {
		if s.BestScores[songPos][i-1].Score >= s.BestScores[songPos][i].Score {
			continue
		}
		s.BestScores[songPos][i-1], s.BestScores[songPos][i] = s.BestScores[songPos][i], s.BestScores[songPos][i-1]
	}
	return
}

func (s UserVoltageRankingSummary) TotalScore() int32 {
	total := int32(0)
	for songPos := 0; songPos < 3; songPos++ {
		for frameNo := int32(0); frameNo < voltageSelectionAmount; frameNo++ {
			total += s.BestScores[songPos][frameNo].Score
		}
	}
	return total / (3 * voltageSelectionAmount)
}

func (s UserVoltageRankingSummary) GetMusicBestScoreList() generic.List[client.EventMiningMusicBestScore] {
	result := generic.List[client.EventMiningMusicBestScore]{}
	for songPos := 0; songPos < 3; songPos++ {
		for frameNo := int32(0); frameNo < voltageSelectionAmount; frameNo++ {
			result.Append(client.EventMiningMusicBestScore{
				LiveId:           s.BestScores[songPos][frameNo].LiveId,
				FrameNo:          frameNo,
				LiveDifficultyId: s.BestScores[songPos][frameNo].LiveDifficultyId,
				Score:            s.BestScores[songPos][frameNo].Score,
			})
		}
	}
	return result

}

func GetUserVoltageRankingSummary(session *userdata.Session, eventId int32) UserVoltageRankingSummary {
	scores := []database.UserEventMiningMusicBestScore{}
	err := session.Db.Table("u_event_mining_music_best_score").Where("user_id = ?", session.UserId).Find(&scores)
	utils.CheckErr(err)
	summary := NewUserVoltageRankingSummary(eventId)
	for _, score := range scores {
		summary.AddRecord(score.LiveId, score.LiveDifficultyId, score.Score)
	}
	return summary
}

// TODO(threading): There is no lock here, because it's implicitly using the database lock
type VoltageRankingType = ranking.Ranking[int32, int32]

// this get invalidated everytime a new event start
// failure to setup the schedule properly will return in bad data
var eventMiningVoltageRanking *VoltageRankingType = nil

func GetVoltageRanking(userdata_db *xorm.Session, eventId int32) *VoltageRankingType {
	if eventId < 0 {
		eventId = gamedata.Instance.EventActive.GetEventMining().EventId
	}
	if eventMiningVoltageRanking != nil {
		return eventMiningVoltageRanking

	}
	eventMiningVoltageRanking = ranking.NewRanking[int32, int32]()
	type UserIdScore struct {
		UserId int32                                  `xorm:"'user_id'"`
		Score  database.UserEventMiningMusicBestScore `xorm:"extends"`
	}
	records := []UserIdScore{}
	err := userdata_db.Table("u_event_mining_music_best_score").OrderBy("user_id").Find(&records)
	utils.CheckErr(err)
	var currentUserSummary UserVoltageRankingSummary
	var currentUserId int32
	for i := range records {
		if (i == 0) || (records[i].UserId != records[i-1].UserId) {
			if currentUserSummary.TotalScore() > 0 {
				eventMiningVoltageRanking.Update(currentUserId, currentUserSummary.TotalScore())
			}
			currentUserId = records[i].UserId
			currentUserSummary = NewUserVoltageRankingSummary(eventId)
		}
		currentUserSummary.AddRecord(records[i].Score.LiveId, records[i].Score.LiveDifficultyId, records[i].Score.Score)
	}
	if currentUserSummary.TotalScore() > 0 {
		eventMiningVoltageRanking.Update(currentUserId, currentUserSummary.TotalScore())
	}
	return eventMiningVoltageRanking
}

func ResetVoltageRanking() {
	eventMiningVoltageRanking = nil
}
