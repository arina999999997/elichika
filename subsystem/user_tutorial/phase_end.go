package user_tutorial

import (
	"elichika/enum"
	"elichika/subsystem/user_reference_book"
	"elichika/subsystem/user_unlock_scene"
	"elichika/userdata"
)

func PhaseEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseFinal {
		panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseTutorialEnd
	session.UserStatus.TutorialEndAt = session.Time.Unix()
	{ // unlock systems
		user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeLesson, enum.UnlockSceneStatusOpen)
		user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeFreeLive, enum.UnlockSceneStatusOpen)
		user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeAccessory, enum.UnlockSceneStatusOpen)
		user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeStoryMember, enum.UnlockSceneStatusOpen)
		user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeEvent, enum.UnlockSceneStatusOpen)
		user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeReferenceBookSelect, enum.UnlockSceneStatusOpen)
	}
	{
		// mark lesson as finished
		user_reference_book.InsertUserReferenceBook(session, 1001)
		user_reference_book.InsertUserReferenceBook(session, 1002)
		user_reference_book.InsertUserReferenceBook(session, 1003)
		user_reference_book.InsertUserReferenceBook(session, 1004)
	}
}
