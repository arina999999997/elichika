package user_live_difficulty

import (
	"elichika/userdata"
	"elichika/utils"
)

func liveDifficultyFinalizer(session *userdata.Session) {
	for _, userLiveDifficulty := range session.UserModel.UserLiveDifficultyByDifficultyId.Map {
		updated, err := session.Db.Table("u_live_difficulty").
			Where("user_id = ? AND live_difficulty_id = ?", session.UserId, userLiveDifficulty.LiveDifficultyId).
			AllCols().Update(*userLiveDifficulty)
		utils.CheckErr(err)
		if updated == 0 {
			userdata.GenericDatabaseInsert(session, "u_live_difficulty", *userLiveDifficulty)
		}
	}

}

func init() {
	userdata.AddFinalizer(liveDifficultyFinalizer)
}
