package gamedata

import (
	"elichika/enum"
	"elichika/log"

	"fmt"
)

func (gamedata *Gamedata) GetEventType(eventId int32) int32 {
	_, isEventMarathon := gamedata.EventMarathon[eventId]

	if isEventMarathon {
		return enum.EventType1Marathon
	} else {
		log.Panic(fmt.Sprint("Unsupported event: ", eventId))
		return 0
	}
}
