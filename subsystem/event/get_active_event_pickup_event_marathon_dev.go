//go:build dev

package event

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/webui/event_marathon_dev"
)

func GetActiveEventPickup(session *userdata.Session) generic.Nullable[client.BootstrapPickupEventInfo] {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		result := generic.NewNullable(client.BootstrapPickupEventInfo{
			EventId:   event_marathon_dev.TopStatus.EventId,
			StartAt:   event_marathon_dev.TopStatus.StartAt,
			ClosedAt:  event_marathon_dev.TopStatus.ExpiredAt,
			EndAt:     event_marathon_dev.TopStatus.EndAt,
			EventType: enum.EventType1Marathon,
		})
		if event_marathon_dev.BoosterItemId != 0 {
			result.Value.BoosterItemId = generic.NewNullable(event_marathon_dev.BoosterItemId)
		}
		return result
	}
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
