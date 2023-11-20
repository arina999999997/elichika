package userdata

import (
	"elichika/enum"
	"elichika/model"
	"elichika/utils"

	"time"
)

func (session *Session) InsertMemberStory(storyMemberMasterID int) {
	// this is correct, but it is obsolete since the client unlock all the bond episode when
	// unlock scene type 4 is set
	// setting UnlockSceneStatusInitial also works but there's no fancy animation so might as well save 1 request
	session.UnlockScene(enum.UnlockSceneTypeStoryMember, enum.UnlockSceneStatusUnlocked)
	userStoryMember := model.UserStoryMember{
		UserID:              session.UserStatus.UserID,
		StoryMemberMasterID: storyMemberMasterID,
		IsNew:               true,
		AcquiredAt:          time.Now().Unix(),
	}
	has, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
		userStoryMember.UserID, userStoryMember.StoryMemberMasterID).Exist(&userStoryMember)
	utils.CheckErr(err)
	if has {
		return
	}
	_, err = session.Db.Table("u_story_member").Insert(userStoryMember)
	utils.CheckErr(err)
	session.UserModel.UserStoryMemberByID.PushBack(userStoryMember)
}

// return true if this is first clear
// insert the story if necessary
func (session *Session) FinishStoryMember(storyMemberMasterID int) bool {
	userStoryMember := model.UserStoryMember{}
	exists, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
		session.UserStatus.UserID, storyMemberMasterID).Get(&userStoryMember)
	utils.CheckErr(err)
	if !exists {
		userStoryMember = model.UserStoryMember{
			UserID:              session.UserStatus.UserID,
			StoryMemberMasterID: storyMemberMasterID,
			IsNew:               false,
			AcquiredAt:          time.Now().Unix(),
		}
		_, err = session.Db.Table("u_story_member").Insert(userStoryMember)
		utils.CheckErr(err)
		session.UserModel.UserStoryMemberByID.PushBack(userStoryMember)
		return true
	}
	if !userStoryMember.IsNew {
		return false
	}
	userStoryMember.IsNew = false
	_, err = session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
		userStoryMember.UserID, userStoryMember.StoryMemberMasterID).AllCols().Update(userStoryMember)
	utils.CheckErr(err)
	return true
}

func init() {
	addGenericTableFieldPopulator("u_story_member", "UserStoryMemberByID")
}
