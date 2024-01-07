package userdata

import (
	"elichika/enum"
	"elichika/model"
	"elichika/utils"
)

func (session *Session) InsertUserStoryMain(storyMainMasterId int) bool {
	userStoryMain := model.UserStoryMain{
		StoryMainMasterId: storyMainMasterId,
	}
	has, err := session.Db.Table("u_story_main").Exist(&userStoryMain)
	utils.CheckErr(err)
	if has { // already have it
		return false
	}

	session.UserModel.UserStoryMainByStoryMainId.PushBack(userStoryMain)
	// also handle unlocking scene (feature)
	// use m_scene_unlock_hint for guide as this seems to be entirely server sided
	// Id is from m_story_main_cell, so maybe load it instead of hard coding
	switch storyMainMasterId {
	case 1007: // k.m_lesson_menu_select_unlock_hint
		session.UnlockScene(enum.UnlockSceneTypeLesson, enum.UnlockSceneStatusOpen)
	case 1009: // k.m_live_music_select_unlock_hint
		session.UnlockScene(enum.UnlockSceneTypeFreeLive, enum.UnlockSceneStatusOpen)
	case 1018: // k.m_accessory_list_unlock_hint
		session.UnlockScene(enum.UnlockSceneTypeAccessory, enum.UnlockSceneStatusOpen)
		session.UnlockScene(enum.UnlockSceneTypeReferenceBookSelect, enum.UnlockSceneStatusOpen)
	default:
	}
	return true
}

func storyMainFinalizer(session *Session) {
	for _, userStoryMain := range session.UserModel.UserStoryMainByStoryMainId.Objects {
		exist, err := session.Db.Table("u_story_main").Exist(&userStoryMain)
		utils.CheckErr(err)
		if !exist {
			genericDatabaseInsert(session, "u_story_main", userStoryMain)
		}
	}
}

func (session *Session) UpdateUserStoryMainSelected(storyMainCellId, selectedId int) {
	userStoryMainSelected := model.UserStoryMainSelected{
		StoryMainCellId: storyMainCellId,
		SelectedId:      selectedId,
	}
	session.UserModel.UserStoryMainSelectedByStoryMainCellId.PushBack(userStoryMainSelected)
}

func storyMainSelectedFinalizer(session *Session) {
	for _, userStoryMainSelected := range session.UserModel.UserStoryMainSelectedByStoryMainCellId.Objects {
		affected, err := session.Db.Table("u_story_main_selected").Where("user_id = ? AND story_main_cell_id = ?",
			session.UserId, userStoryMainSelected.StoryMainCellId).AllCols().Update(userStoryMainSelected)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_story_main_selected", userStoryMainSelected)
		}
	}
}

func (session *Session) InsertUserStoryMainPartDigestMovie(partId int) {
	userStoryMainPartDigestMovie := model.UserStoryMainPartDigestMovie{
		StoryMainPartMasterId: partId,
	}
	session.UserModel.UserStoryMainPartDigestMovieById.PushBack(userStoryMainPartDigestMovie)
}

func storyMainPartDigestMovieFinalizer(session *Session) {
	for _, userStoryMainPartDigestMovie := range session.UserModel.UserStoryMainPartDigestMovieById.Objects {
		exist, err := session.Db.Table("u_story_main_part_digest_movie").Exist(&userStoryMainPartDigestMovie)
		utils.CheckErr(err)
		if !exist {
			genericDatabaseInsert(session, "u_story_main_part_digest_movie", userStoryMainPartDigestMovie)
		}
	}
}

func (session *Session) InsertUserStoryLinkage(storyLinkageCellMasterId int) {
	userStoryLinkage := model.UserStoryLinkage{
		StoryLinkageCellMasterId: storyLinkageCellMasterId,
	}
	has, err := session.Db.Table("u_story_linkage").Exist(&userStoryLinkage)
	utils.CheckErr(err)
	if has {
		return
	}
	session.UserModel.UserStoryLinkageById.PushBack(userStoryLinkage)
}

func storyLinkageFinalizer(session *Session) {
	for _, userStoryLinkage := range session.UserModel.UserStoryLinkageById.Objects {
		exist, err := session.Db.Table("u_story_linkage").Exist(&userStoryLinkage)
		utils.CheckErr(err)
		if !exist {
			genericDatabaseInsert(session, "u_story_linkage", userStoryLinkage)
		}
	}
}

func init() {
	addFinalizer(storyMainFinalizer)
	addFinalizer(storyMainSelectedFinalizer)
	addFinalizer(storyMainPartDigestMovieFinalizer)
	addFinalizer(storyLinkageFinalizer)
	addGenericTableFieldPopulator("u_story_main", "UserStoryMainByStoryMainId")
	addGenericTableFieldPopulator("u_story_main_selected", "UserStoryMainSelectedByStoryMainCellId")
	addGenericTableFieldPopulator("u_story_main_part_digest_movie", "UserStoryMainPartDigestMovieById")
	addGenericTableFieldPopulator("u_story_linkage", "UserStoryLinkageById")
}
