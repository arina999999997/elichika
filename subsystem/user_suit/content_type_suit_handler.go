package user_suit

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func contentTypeSuitHandler(session *userdata.Session, content client.Content) (bool, any) {
	InsertUserSuit(session, content.ContentId)
	return true, nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSuit, contentTypeSuitHandler)
}
