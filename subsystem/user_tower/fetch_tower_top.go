package user_tower

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

// there's actually no error response for now, as towers are eternal
func FetchTowerTop(session *userdata.Session, towerId int32) (*response.FetchTowerTopResponse, *response.RecoverableExceptionResponse) {
	resp := response.FetchTowerTopResponse{
		TowerCardUsedCountRows: GetUserTowerCardUsedList(session, towerId),
		UserModelDiff:          &session.UserModel,
		// other fields are for DLP with voltage ranking
	}

	userTower := GetUserTower(session, towerId)
	tower := session.Gamedata.Tower[towerId]
	if userTower.ClearedFloor == userTower.ReadFloor {
		tower := session.Gamedata.Tower[towerId]
		if userTower.ReadFloor < tower.FloorCount {
			userTower.ReadFloor += 1
			resp.IsShowUnlockEffect = true
			// unlock all the bonus live at once
			for ; userTower.ReadFloor < tower.FloorCount; userTower.ReadFloor++ {
				if tower.Floor[userTower.ReadFloor].TowerCellType != enum.TowerCellTypeBonusLive {
					break
				}
			}
		}
	}
	UpdateUserTower(session, userTower)

	// if tower with voltage ranking, then we have to prepare that
	if tower.IsVoltageRanked {
		// EachBonusLiveVoltage should be filled with zero for everything, then fill in the score

		resp.EachBonusLiveVoltage.Slice = make([]int32, tower.FloorCount)
		myRank, hasRank := GetRankingByTowerId(session, towerId).RankOf(session.UserId)
		if hasRank {
			resp.Order = generic.NewNullable(int32(myRank))
		} else {
			// setting this allow the ranking button to actually show up, as the client use this to check for ranking
			resp.Order = generic.NewNullable(int32(0))
		}
		// fetch the score
		scores := GetUserTowerVoltageRankingScores(session, towerId)
		for _, score := range scores {
			resp.EachBonusLiveVoltage.Slice[score.FloorNo-1] = score.Voltage
		}
	}
	return &resp, nil
}
