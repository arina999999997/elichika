package gamedata

import (
	"elichika/dictionary"
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type EventAvailable struct {
	Count          int32
	EventIds       []int32
	EventIdToOrder map[int32]int32
}

func (ea *EventAvailable) Build() {
	ea.Count = int32(len(ea.EventIds))
	ea.EventIdToOrder = map[int32]int32{}
	for i, eventId := range ea.EventIds {
		ea.EventIdToOrder[eventId] = int32(i)
	}
}

// reload the event if necessary
// return nil if the event doesn't exist
func (ea *EventAvailable) GetNextEvent(currentEvent *serverdata.EventActive) int32 {
	if currentEvent == nil {
		return ea.EventIds[0]
	}
	order, exist := ea.EventIdToOrder[currentEvent.EventId]
	if !exist {
		return ea.EventIds[0]
	} else {
		return ea.EventIds[(order+1)%ea.Count]
	}

}

func loadEventAvailable(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := serverdata_db.Table("s_event_available").OrderBy(`"order"`).Cols("event_id").Find(&gamedata.EventAvailable.EventIds)
	utils.CheckErr(err)
	gamedata.EventAvailable.Build()
}

func init() {
	addLoadFunc(loadEventAvailable)
}
