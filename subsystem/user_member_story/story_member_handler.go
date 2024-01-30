package user_member_story

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func storyMemberHandler(session *userdata.Session, content *client.Content) any {
	session.InsertMemberStory(content.ContentId)
	content.ContentAmount = 0
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeStoryMember, storyMemberHandler)
}
