package user_story_main

import (
	"elichika/userdata"
	"elichika/utils"
)

func userStoryMainSelectedFinalizer(session *userdata.Session) {
	for _, userStoryMainSelected := range session.UserModel.UserStoryMainSelectedByStoryMainCellId.Map {
		affected, err := session.Db.Table("u_story_main_selected").
			Where("user_id = ? AND story_main_cell_id = ?", session.UserId, userStoryMainSelected.StoryMainCellId).
			AllCols().Update(*userStoryMainSelected)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_story_main_selected", *userStoryMainSelected)
		}
	}
}
func init() {
	userdata.AddFinalizer(userStoryMainSelectedFinalizer)
}
