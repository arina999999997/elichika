package user_accessory
import (
	"elichika/userdata"
)
func UnequipAccessoryFromAllDeck(session *userdata.Session, userAccessoryId int64) {
	
	liveParties := session.GetAllLivePartiesWithAccessory(userAccessoryId)
	for _, liveParty := range liveParties {
		if (liveParty.UserAccessoryId1.HasValue) && (liveParty.UserAccessoryId1.Value == userAccessoryId) {
			liveParty.UserAccessoryId1.HasValue = false
		}
		if (liveParty.UserAccessoryId2.HasValue) && (liveParty.UserAccessoryId2.Value == userAccessoryId) {
			liveParty.UserAccessoryId2.HasValue = false
		}
		if (liveParty.UserAccessoryId3.HasValue) && (liveParty.UserAccessoryId3.Value == userAccessoryId) {
			liveParty.UserAccessoryId3.HasValue = false
		}
		session.UpdateUserLiveParty(liveParty)
	}
}