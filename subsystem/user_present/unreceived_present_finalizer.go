package user_present

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
)

func unreceivedPresentFinalizer(session *userdata.Session) {
	for _, content := range session.UnreceivedContent {
		// there doesn't seems to be a meaningful difference between these
		if content.ContentType == enum.ContentTypeAccessory {
			AddPresent(session, client.PresentItem{
				Content:          content,
				PresentRouteType: enum.PresentRouteTypeLiveAccessoryItemFull,
			})
		} else {
			AddPresent(session, client.PresentItem{
				Content:          content,
				PresentRouteType: enum.PresentRouteTypeItemFull,
			})
		}
	}
}

func init() {
	userdata.AddFinalizer(unreceivedPresentFinalizer)
}
