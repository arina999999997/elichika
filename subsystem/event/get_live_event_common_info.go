package event

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

func GetLiveEventCommonInfo(session *userdata.Session) generic.Nullable[client.LiveEventCommonInfo] {
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if (event == nil) || (event.ExpiredAt <= session.Time.Unix()) {
		return generic.Nullable[client.LiveEventCommonInfo]{}
	}
	result := client.LiveEventCommonInfo{
		EventId:   event.EventId,
		EventType: event.EventType,
		ClosedAt:  event.ExpiredAt,
	}
	if event.EventType == enum.EventType1Marathon {
		result.PointBoostContentId = generic.NewNullable(session.Gamedata.EventActive.GetEventMarathon().BoosterItemId)
	} else {
		// TODO(event mining)
		panic("not supported")
	}
	return generic.NewNullable(result)
}
