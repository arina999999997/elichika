package user_voice

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func contentTypeVoiceHandler(session *userdata.Session, content client.Content) (bool, any) {
	session.UpdateVoice(content.ContentId, false)
	return true, nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeVoice, contentTypeVoiceHandler)
}
