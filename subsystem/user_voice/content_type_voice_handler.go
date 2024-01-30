package user_voice

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func contentTypeVoiceHandler(session *userdata.Session, content *client.Content) any {
	session.UpdateVoice(content.ContentId, false)
	content.ContentAmount = 0
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeVoice, contentTypeVoiceHandler)
}
