package user_card

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func cardContentHandler(session *userdata.Session, content *client.Content) any {
	content.ContentAmount = 0
	return AddUserCardByCardMasterId(session, content.ContentId)
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeCard, cardContentHandler)
}
