package user_beginner_challenge

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func UpdateChallengeCell(session *userdata.Session, cell client.ChallengeCell) {
	affected, err := session.Db.Table("u_beginner_challenge_cell").Where("user_id = ? AND cell_id = ?", session.UserId, cell.CellId).
		AllCols().Update(cell)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_beginner_challenge_cell", cell)
	}
}
