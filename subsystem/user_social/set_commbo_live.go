package user_social

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func SetCommboLive(session *userdata.Session, liveDifficultyMasterId int32) {
	result, err := session.Db.Exec("UPDATE u_set_profile SET commbo_live_difficulty_id = ? WHERE user_id = ?",
		liveDifficultyMasterId, session.UserId)
	utils.CheckErr(err)
	affected, err := result.RowsAffected()
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_set_profile", client.UserSetProfile{
			CommboLiveDifficultyId: liveDifficultyMasterId,
		})
	}
}
