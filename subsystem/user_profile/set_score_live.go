package user_profile

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func SetScoreLive(session *userdata.Session, liveDifficultyMasterId int32) {
	result, err := session.Db.Exec("UPDATE u_set_profile SET voltage_live_difficulty_id = ? WHERE user_id = ?",
		liveDifficultyMasterId, session.UserId)
	utils.CheckErr(err)
	affected, err := result.RowsAffected()
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_set_profile", client.UserSetProfile{
			VoltageLiveDifficultyId: liveDifficultyMasterId,
		})
	}
}
