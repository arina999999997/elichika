package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/utils"
)

func (session *Session) InsertUserStoryMain(storyMainMasterId int32) bool {
	userStoryMain := client.UserStoryMain{
		StoryMainMasterId: storyMainMasterId,
	}
	if genericDatabaseExist(session, "u_story_main", userStoryMain) {
		return false
	}

	session.UserModel.UserStoryMainByStoryMainId.Set(storyMainMasterId, userStoryMain)
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
	for _, userStoryMain := range session.UserModel.UserStoryMainByStoryMainId.Map {
		if !genericDatabaseExist(session, "u_story_main", *userStoryMain) {
			GenericDatabaseInsert(session, "u_story_main", *userStoryMain)
		}
	}
}

func (session *Session) UpdateUserStoryMainSelected(storyMainCellId, selectedId int32) {
	userStoryMainSelected := client.UserStoryMainSelected{
		StoryMainCellId: storyMainCellId,
		SelectedId:      selectedId,
	}
	session.UserModel.UserStoryMainSelectedByStoryMainCellId.Set(storyMainCellId, userStoryMainSelected)
}

func storyMainSelectedFinalizer(session *Session) {
	for _, userStoryMainSelected := range session.UserModel.UserStoryMainSelectedByStoryMainCellId.Map {
		affected, err := session.Db.Table("u_story_main_selected").
			Where("user_id = ? AND story_main_cell_id = ?", session.UserId, userStoryMainSelected.StoryMainCellId).
			AllCols().Update(*userStoryMainSelected)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_story_main_selected", *userStoryMainSelected)
		}
	}
}

func (session *Session) InsertUserStoryMainPartDigestMovie(partId int32) {
	session.UserModel.UserStoryMainPartDigestMovieById.Set(partId, client.UserStoryMainPartDigestMovie{
		StoryMainPartMasterId: partId,
	})
}

func storyMainPartDigestMovieFinalizer(session *Session) {
	for _, userStoryMainPartDigestMovie := range session.UserModel.UserStoryMainPartDigestMovieById.Map {
		if !genericDatabaseExist(session, "u_story_main_part_digest_movie", *userStoryMainPartDigestMovie) {
			GenericDatabaseInsert(session, "u_story_main_part_digest_movie", *userStoryMainPartDigestMovie)
		}
	}
}

func (session *Session) InsertUserStoryLinkage(storyLinkageCellMasterId int32) {
	userStoryLinkage := client.UserStoryLinkage{
		StoryLinkageCellMasterId: storyLinkageCellMasterId,
	}
	if !genericDatabaseExist(session, "u_story_linkage", userStoryLinkage) {
		session.UserModel.UserStoryLinkageById.Set(storyLinkageCellMasterId, userStoryLinkage)
	}
}

func storyLinkageFinalizer(session *Session) {
	for _, userStoryLinkage := range session.UserModel.UserStoryLinkageById.Map {
		if !genericDatabaseExist(session, "u_story_linkage", *userStoryLinkage) {
			GenericDatabaseInsert(session, "u_story_linkage", *userStoryLinkage)
		}
	}
}

func init() {
	AddFinalizer(storyMainFinalizer)
	AddFinalizer(storyMainSelectedFinalizer)
	AddFinalizer(storyMainPartDigestMovieFinalizer)
	AddFinalizer(storyLinkageFinalizer)
}
