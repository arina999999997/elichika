package marathon

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func RemoveBoosterItem(session *userdata.Session, count int32) {
	event := session.Gamedata.EventActive.GetEventMarathon()
	user_content.RemoveContent(session, client.Content{
		ContentType:   enum.ContentTypeEventMarathonBooster,
		ContentId:     event.BoosterItemId,
		ContentAmount: count,
	})
}
