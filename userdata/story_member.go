package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/utils"
)

func (session *Session) InsertMemberStory(storyMemberMasterId int32) {
	// this is correct, but it is obsolete since the client unlock all the bond episode when
	// unlock scene type 4 is set
	// setting UnlockSceneStatusOpen also works but there's no fancy animation so might as well save 1 request
	session.UnlockScene(enum.UnlockSceneTypeStoryMember, enum.UnlockSceneStatusOpened)
	userStoryMember := client.UserStoryMember{
		StoryMemberMasterId: storyMemberMasterId,
		IsNew:               true,
		AcquiredAt:          session.Time.Unix(),
	}
	session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)
}

func memberStoryFinalizer(session *Session) {
	for _, userStoryMember := range session.UserModel.UserStoryMemberById.Map {
		affected, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
			session.UserId, userStoryMember.StoryMemberMasterId).AllCols().Update(*userStoryMember)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_story_member", *userStoryMember)
		}
	}
}

// return true if this is first clear
// insert the story if necessary
func (session *Session) FinishStoryMember(storyMemberMasterId int32) bool {
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

func init() {
	addFinalizer(memberStoryFinalizer)
	addGenericTableFieldPopulator("u_story_member", "UserStoryMemberById")
}
