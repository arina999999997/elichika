package event

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

func GetActiveEventPickup(session *userdata.Session) generic.Nullable[client.BootstrapPickupEventInfo] {
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if event == nil {
		return generic.Nullable[client.BootstrapPickupEventInfo]{}
	}
	result := generic.NewNullable(client.BootstrapPickupEventInfo{
		EventId:   event.EventId,
		StartAt:   event.StartAt,
		ClosedAt:  event.ExpiredAt,
		EndAt:     event.EndAt,
		EventType: event.EventType,
	})
	if session.Time.Unix() < event.ExpiredAt {
		if event.EventType == enum.EventType1Marathon {
			result.Value.BoosterItemId = generic.NewNullable(session.Gamedata.EventActive.GetEventMarathon().BoosterItemId)
		}
	}
	return result
}
