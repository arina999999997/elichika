package user_story_member

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertMemberStory(session *userdata.Session, storyMemberMasterId int32) {
	userStoryMember := client.UserStoryMember{
		StoryMemberMasterId: storyMemberMasterId,
		IsNew:               true,
		AcquiredAt:          session.Time.Unix(),
	}
	session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)
}
