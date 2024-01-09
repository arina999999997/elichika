package userdata

import (
	"elichika/client"
	"elichika/model"
	"elichika/protocol/response"
	"elichika/utils"
)

func (session *Session) GetUserTowerCardUsed(towerId, cardMasterId int32) model.UserTowerCardUsedCount {
	cardUsed := model.UserTowerCardUsedCount{}
	exist, err := session.Db.Table("u_tower_card_used").
		Where("user_id = ? AND tower_id = ? AND card_master_id = ?", session.UserId, towerId, cardMasterId).Get(&cardUsed)
	utils.CheckErr(err)
	if !exist {
		cardUsed = model.UserTowerCardUsedCount{
			TowerId:        towerId,
			CardMasterId:   cardMasterId,
			UsedCount:      0,
			RecoveredCount: 0,
			LastUsedAt:     0,
		}
	}
	return cardUsed
}

func (session *Session) UpdateUserTowerCardUsed(card model.UserTowerCardUsedCount) {
	affected, err := session.Db.Table("u_tower_card_used").
		Where("user_id = ? AND tower_id = ? AND card_master_id = ?", session.UserId, card.TowerId, card.CardMasterId).
		AllCols().Update(card)
	utils.CheckErr(err)
	if affected == 0 {
		genericDatabaseInsert(session, "u_tower_card_used", card)
	}
}

func (session *Session) GetUserTowerCardUsedList(towerId int32) []model.UserTowerCardUsedCount {
	list := []model.UserTowerCardUsedCount{}
	err := session.Db.Table("u_tower_card_used").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Find(&list)
	utils.CheckErr(err)
	return list
}

func (session *Session) GetUserTower(towerId int32) client.UserTower {
	pos, exist := session.UserTowerMapping.SetList(&session.UserModel.UserTowerByTowerId).Map[int64(towerId)]
	if exist {
		return session.UserModel.UserTowerByTowerId.Objects[pos]
	}
	tower := client.UserTower{}
	exist, err := session.Db.Table("u_tower").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Get(&tower)
	utils.CheckErr(err)
	if !exist {
		tower = client.UserTower{
			TowerId:                     towerId,
			ClearedFloor:                0,
			ReadFloor:                   0,
			Voltage:                     0,
			RecoveryPointFullAt:         session.Time.Unix() + 86400,
			RecoveryPointLastConsumedAt: session.Time.Unix(),
		}
	} else {

		// this is to make sure user can always mass recover LP
		// it's the only way to do this since these things are decided by the timestamps alone (+ plus the limit from the database)
		tower.RecoveryPointFullAt = session.Time.Unix() + 86400
		tower.RecoveryPointLastConsumedAt = session.Time.Unix()
	}
	return tower
}

func (session *Session) UpdateUserTower(tower client.UserTower) {
	session.UserTowerMapping.SetList(&session.UserModel.UserTowerByTowerId).Update(tower)
}

func (session *Session) GetUserTowerVoltageRankingScores(towerId int32) []model.UserTowerVoltageRankingScore {
	scores := []model.UserTowerVoltageRankingScore{}
	err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Find(&scores)
	utils.CheckErr(err)
	return scores
}

func (session *Session) GetUserTowerVoltageRankingScore(towerId, floorNo int32) model.UserTowerVoltageRankingScore {
	score := model.UserTowerVoltageRankingScore{}
	exists, err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ? AND floor_no = ?", session.UserId, towerId, floorNo).Get(&score)
	utils.CheckErr(err)
	if !exists {
		score = model.UserTowerVoltageRankingScore{
			TowerId: towerId,
			FloorNo: floorNo,
			Voltage: 0,
		}
	}
	return score
}

func (session *Session) UpdateUserTowerVoltageRankingScore(score model.UserTowerVoltageRankingScore) {
	affected, err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ? AND floor_no = ?", session.UserId, score.TowerId, score.FloorNo).AllCols().
		Update(score)
	utils.CheckErr(err)
	if affected == 0 {
		genericDatabaseInsert(session, "u_tower_voltage_ranking_score", score)
	}
}

func (session *Session) GetTowerRankingCell(towerId int32) response.TowerRankingCell {
	scores := session.GetUserTowerVoltageRankingScores(towerId)
	cell := response.TowerRankingCell{
		Order:            1,
		SumVoltage:       0,
		TowerRankingUser: session.GetRankingUser(),
	}
	for _, score := range scores {
		cell.SumVoltage += score.Voltage
	}
	return cell
}

func towerFinalizer(session *Session) {
	for _, userTower := range session.UserModel.UserTowerByTowerId.Objects {
		affected, err := session.Db.Table("u_tower").
			Where("user_id = ? AND tower_id = ?", session.UserId, userTower.TowerId).
			AllCols().Update(userTower)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_tower", userTower)
		}
	}
}

func init() {
	addFinalizer(towerFinalizer)
	addGenericTableFieldPopulator("u_tower", "UserTowerByTowerId")
}
