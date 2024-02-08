package user_story_member

import (
	"elichika/userdata"
	"elichika/utils"
)

func userStoryMemberFinalizer(session *userdata.Session) {
	for _, userStoryMember := range session.UserModel.UserStoryMemberById.Map {
		affected, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
			session.UserId, userStoryMember.StoryMemberMasterId).AllCols().Update(*userStoryMember)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_story_member", *userStoryMember)
		}
	}
}

func init() {
	userdata.AddFinalizer(userStoryMemberFinalizer)
}
