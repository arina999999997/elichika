package userdata

import (
	"elichika/enum"
	"elichika/model"
	"elichika/utils"
)

func (session *Session) InsertUserStoryMain(storyMainMasterID int) bool {
	userStoryMain := model.UserStoryMain{
		UserID:            session.UserStatus.UserID,
		StoryMainMasterID: storyMainMasterID,
	}
	has, err := session.Db.Table("u_story_main").Exist(&userStoryMain)
	utils.CheckErr(err)
	if has { // already have it
		return false
	}

	_, err = session.Db.Table("u_story_main").Insert(userStoryMain)
	utils.CheckErr(err)
	session.UserModel.UserStoryMainByStoryMainID.PushBack(userStoryMain)
	// also handle unlocking scene (feature)
	// use m_scene_unlock_hint for guide as this seems to be entirely server sided
	// ID is from m_story_main_cell, so maybe load it instead of hard coding
	switch storyMainMasterID {
	case 1007: // k.m_lesson_menu_select_unlock_hint
		session.UnlockScene(enum.UnlockSceneTypeLessonMenuSelect, enum.UnlockSceneStatusInitial)
	case 1009: // k.m_live_music_select_unlock_hint
		session.UnlockScene(enum.UnlockSceneTypeLiveMusicSelect, enum.UnlockSceneStatusInitial)
	case 1018: // k.m_accessory_list_unlock_hint
		session.UnlockScene(enum.UnlockSceneTypeAccessoryList, enum.UnlockSceneStatusInitial)
		session.UnlockScene(enum.UnlockSceneTypeReferenceBookSelect, enum.UnlockSceneStatusInitial)
	default:
	}
	return true
}

func (session *Session) UpdateUserStoryMainSelected(storyMainCellID, selectedID int) {
	userStoryMainSelected := model.UserStoryMainSelected{
		UserID:          session.UserStatus.UserID,
		StoryMainCellID: storyMainCellID,
		SelectedID:      selectedID,
	}
	affected, err := session.Db.Table("u_story_main_selected").Where("user_id = ? AND story_main_cell_id = ?",
		userStoryMainSelected.UserID, userStoryMainSelected.StoryMainCellID).AllCols().Update(userStoryMainSelected)
	utils.CheckErr(err)
	if affected == 0 {
		_, err := session.Db.Table("u_story_main_selected").Insert(userStoryMainSelected)
		utils.CheckErr(err)
	}
	session.UserModel.UserStoryMainSelectedByStoryMainCellID.PushBack(userStoryMainSelected)

}

func (session *Session) InsertUserStoryMainPartDigestMovie(partID int) {
	userStoryMainPartDigestMovie := model.UserStoryMainPartDigestMovie{
		UserID:                session.UserStatus.UserID,
		StoryMainPartMasterID: partID,
	}
	has, err := session.Db.Table("u_story_main_part_digest_movie").Exist(&userStoryMainPartDigestMovie)
	utils.CheckErr(err)
	if has { // already have it
		return
	}
	_, err = session.Db.Table("u_story_main_part_digest_movie").Insert(userStoryMainPartDigestMovie)
	utils.CheckErr(err)
	session.UserModel.UserStoryMainPartDigestMovieByID.PushBack(userStoryMainPartDigestMovie)
}

func (session *Session) InsertUserStoryLinkage(storyLinkageCellMasterID int) {
	userStoryLinkage := model.UserStoryLinkage{
		UserID:                   session.UserStatus.UserID,
		StoryLinkageCellMasterID: storyLinkageCellMasterID,
	}
	has, err := session.Db.Table("u_story_linkage").Exist(&userStoryLinkage)
	utils.CheckErr(err)
	if has {
		return
	}

	_, err = session.Db.Table("u_story_linkage").Insert(userStoryLinkage)
	utils.CheckErr(err)
	session.UserModel.UserStoryLinkageByID.PushBack(userStoryLinkage)
}

func init() {
	addGenericTableFieldPopulator("u_story_main", "UserStoryMainByStoryMainID")
	addGenericTableFieldPopulator("u_story_main_selected", "UserStoryMainSelectedByStoryMainCellID")
	addGenericTableFieldPopulator("u_story_main_part_digest_movie", "UserStoryMainPartDigestMovieByID")
	addGenericTableFieldPopulator("u_story_linkage", "UserStoryLinkageByID")
}
