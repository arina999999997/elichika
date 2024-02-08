package user_story_member

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

// return true if this is first clear
// insert the story if necessary
func FinishStoryMember(session *userdata.Session, storyMemberMasterId int32) bool {
	userStoryMember := client.UserStoryMember{}
	exist, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
		session.UserId, storyMemberMasterId).Get(&userStoryMember)
	utils.CheckErr(err)
	if !exist {
		userStoryMember = client.UserStoryMember{
			StoryMemberMasterId: storyMemberMasterId,
			IsNew:               false,
			AcquiredAt:          session.Time.Unix(),
		}
		session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)
	}
	if !userStoryMember.IsNew {
		return false
	}
	userStoryMember.IsNew = false
	session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)
	return true
}
