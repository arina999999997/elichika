package gamedata

import (
	"elichika/enum"

	"fmt"
)

func (gamedata *Gamedata) GetEventType(eventId int32) int32 {
	_, isEventMarathon := gamedata.EventMarathon[eventId]

	if isEventMarathon {
		return enum.EventType1Marathon
	} else {
		panic(fmt.Sprint("Unsupported event: ", eventId))
	}
}
