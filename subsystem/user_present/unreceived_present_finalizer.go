package user_present

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
)

func unreceivedPresentFinalizer(session *userdata.Session) {
	// there has to be a duration, especially for accessory, otherwise it would be very questionable
	for _, content := range session.UnreceivedContent {
		// there doesn't seems to be a meaningful difference between these
		if content.ContentType == enum.ContentTypeAccessory {
			AddPresentWithDuration(session, client.PresentItem{
				Content:          content,
				PresentRouteType: enum.PresentRouteTypeLiveAccessoryItemFull,
			}, Duration30Days)
		} else {
			AddPresentWithDuration(session, client.PresentItem{
				Content:          content,
				PresentRouteType: enum.PresentRouteTypeItemFull,
			}, Duration30Days)
		}
	}
}

func init() {
	userdata.AddFinalizer(unreceivedPresentFinalizer)
}
