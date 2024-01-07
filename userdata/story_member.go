package userdata

import (
	"elichika/enum"
	"elichika/model"
	"elichika/utils"
)

func (session *Session) InsertMemberStory(storyMemberMasterId int) {
	// this is correct, but it is obsolete since the client unlock all the bond episode when
	// unlock scene type 4 is set
	// setting UnlockSceneStatusOpen also works but there's no fancy animation so might as well save 1 request
	session.UnlockScene(enum.UnlockSceneTypeStoryMember, enum.UnlockSceneStatusOpened)
	userStoryMember := model.UserStoryMember{
		UserId:              session.UserStatus.UserId,
		StoryMemberMasterId: storyMemberMasterId,
		IsNew:               true,
		AcquiredAt:          session.Time.Unix(),
	}
	session.UserModel.UserStoryMemberById.PushBack(userStoryMember)
}

func memberStoryFinalizer(session *Session) {
	for _, userStoryMember := range session.UserModel.UserStoryMemberById.Objects {
		affected, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
			userStoryMember.UserId, userStoryMember.StoryMemberMasterId).AllCols().Update(userStoryMember)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_story_member").Insert(userStoryMember)
			utils.CheckErr(err)
		}
	}
}

// return true if this is first clear
// insert the story if necessary
func (session *Session) FinishStoryMember(storyMemberMasterId int) bool {
	userStoryMember := model.UserStoryMember{}
	exist, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
		session.UserStatus.UserId, storyMemberMasterId).Get(&userStoryMember)
	utils.CheckErr(err)
	if !exist {
		userStoryMember = model.UserStoryMember{
			UserId:              session.UserStatus.UserId,
			StoryMemberMasterId: storyMemberMasterId,
			IsNew:               false,
			AcquiredAt:          session.Time.Unix(),
		}
		session.UserModel.UserStoryMemberById.PushBack(userStoryMember)
	}
	if !userStoryMember.IsNew {
		return false
	}
	userStoryMember.IsNew = false
	session.UserModel.UserStoryMemberById.PushBack(userStoryMember)
	return true
}

func init() {
	addFinalizer(memberStoryFinalizer)
	addGenericTableFieldPopulator("u_story_member", "UserStoryMemberById")
}
