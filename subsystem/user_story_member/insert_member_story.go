package user_story_member

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
)

func InsertMemberStory(session *userdata.Session, storyMemberMasterId int32) {
	// this is correct, but it is obsolete since the client unlock all the bond episode when
	// unlock scene type 4 is set
	// setting UnlockSceneStatusOpen also works but there's no fancy animation so might as well save 1 request
	// TODO(tutorial): Just unlock this from the start?
	session.UnlockScene(enum.UnlockSceneTypeStoryMember, enum.UnlockSceneStatusOpened)
	userStoryMember := client.UserStoryMember{
		StoryMemberMasterId: storyMemberMasterId,
		IsNew:               true,
		AcquiredAt:          session.Time.Unix(),
	}
	session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)
}
