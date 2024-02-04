package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata/database"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func (session *Session) GetUserTowerCardUsed(towerId, cardMasterId int32) client.TowerCardUsedCount {
	cardUsed := client.TowerCardUsedCount{}
	exist, err := session.Db.Table("u_tower_card_used_count").
		Where("user_id = ? AND tower_id = ? AND card_master_id = ?", session.UserId, towerId, cardMasterId).Get(&cardUsed)
	utils.CheckErr(err)
	if !exist {
		cardUsed = client.TowerCardUsedCount{
			CardMasterId:   cardMasterId,
			UsedCount:      0,
			RecoveredCount: 0,
			LastUsedAt:     0,
		}
	}
	return cardUsed
}

func (session *Session) UpdateUserTowerCardUsed(towerId int32, card client.TowerCardUsedCount) {
	affected, err := session.Db.Table("u_tower_card_used_count").
		Where("user_id = ? AND tower_id = ? AND card_master_id = ?", session.UserId, towerId, card.CardMasterId).
		AllCols().Update(card)
	utils.CheckErr(err)
	if affected == 0 {
		type Wrapper struct {
			Card    client.TowerCardUsedCount `xorm:"extends"`
			TowerId int32                     `xorm:"pk 'tower_id'"`
		}
		GenericDatabaseInsert(session, "u_tower_card_used_count", Wrapper{
			Card:    card,
			TowerId: towerId,
		})
	}
}

func (session *Session) GetUserTowerCardUsedList(towerId int32) generic.List[client.TowerCardUsedCount] {
	list := generic.List[client.TowerCardUsedCount]{}
	err := session.Db.Table("u_tower_card_used_count").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Find(&list.Slice)
	utils.CheckErr(err)
	return list
}

func (session *Session) GetUserTower(towerId int32) client.UserTower {
	ptr, exist := session.UserModel.UserTowerByTowerId.Get(towerId)
	if exist {
		return *ptr
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
	session.UserModel.UserTowerByTowerId.Set(tower.TowerId, tower)
}

func (session *Session) GetUserTowerVoltageRankingScores(towerId int32) []database.UserTowerVoltageRankingScore {
	scores := []database.UserTowerVoltageRankingScore{}
	err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Find(&scores)
	utils.CheckErr(err)
	return scores
}

func (session *Session) GetUserTowerVoltageRankingScore(towerId, floorNo int32) database.UserTowerVoltageRankingScore {
	score := database.UserTowerVoltageRankingScore{}
	exists, err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ? AND floor_no = ?", session.UserId, towerId, floorNo).Get(&score)
	utils.CheckErr(err)
	if !exists {
		score = database.UserTowerVoltageRankingScore{
			TowerId: towerId,
			FloorNo: floorNo,
			Voltage: 0,
		}
	}
	return score
}

func (session *Session) UpdateUserTowerVoltageRankingScore(score database.UserTowerVoltageRankingScore) {
	affected, err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ? AND floor_no = ?", session.UserId, score.TowerId, score.FloorNo).AllCols().
		Update(score)
	utils.CheckErr(err)
	if affected == 0 {
		GenericDatabaseInsert(session, "u_tower_voltage_ranking_score", score)
	}
}

func (session *Session) GetTowerRankingCell(towerId int32) client.TowerRankingCell {
	scores := session.GetUserTowerVoltageRankingScores(towerId)
	cell := client.TowerRankingCell{
		Order:            1,
		SumVoltage:       0,
		TowerRankingUser: session.GetTowerRankingUser(),
	}
	for _, score := range scores {
		cell.SumVoltage += score.Voltage
	}
	return cell
}

func towerFinalizer(session *Session) {
	for _, userTower := range session.UserModel.UserTowerByTowerId.Map {
		affected, err := session.Db.Table("u_tower").
			Where("user_id = ? AND tower_id = ?", session.UserId, userTower.TowerId).
			AllCols().Update(*userTower)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_tower", *userTower)
		}
	}
}

func init() {
	AddFinalizer(towerFinalizer)
}
