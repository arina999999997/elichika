package gamedata

import (
	"elichika/dictionary"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"
	"time"
	"xorm.io/xorm"
)

// event system:
// - events can be one of the following:
//   - marathon event: gain event points for rewards.
//   - mining event: gain event points for exchanging for reward (not supported yet)
//   - super big live: coop SBL (not supported yet)
// - all events available are presents in the server and client:
//   - all the necessary data in client or server is expected to be there already.
// - at all time, EXACTLY ONE  event can be active, stored in "s_event_active":
//   - active event need to set the following values:
//     - event start_at: when the event start
//     - event expire_at: when the event stop (no more ep / exchange items can be earned)
//     - event result_at: when the result is published (usually 3 hours after expire in official server)
//     - event end_at: when the event end (stopped showing up)
//     - we expect the following relationship: start_at < expire_at <= result_at < end_at
//   - active event can be changed with the server being up:
//     - no server restart should be required to start / extend / end events.
//     - howerver, some data might become outdated, so simply changing the database would require a restart.
// - event payout happen at event result time for every user who:
//   - earned at least 1 event point
//   - logged in the range [start_at, end_at]
//   - due to some limitation, the event result animations might be skipped if a new event is already in effect.

// handling object, this can refetch from database if necessary
type EventActive struct {
	Event    *serverdata.EventActive
	Gamedata *Gamedata `xorm:"-"`
}

// reload the event if necessary
// return nil if the event doesn't exist
func (ae *EventActive) GetEventValue() *serverdata.EventActive {
	if ae.Event == nil {
		event := serverdata.EventActive{}
		exist, err := ae.Gamedata.ServerdataDb.Table("s_event_active").Get(&event)
		fmt.Println("trying to load event active: ", event)
		utils.CheckErr(err)
		if exist {
			// do some check
			if (event.StartAt >= event.ExpiredAt) || (event.ExpiredAt > event.ResultAt) || (event.ResultAt >= event.EndAt) {
				panic("active event have bad time constraint")
			}
			ae.Event = &event
		}
	}
	return ae.Event
}

func (ae *EventActive) GetEventMarathon() *EventMarathon {
	event := ae.GetEventValue()
	if event == nil {
		return nil
	}
	return ae.Gamedata.EventMarathon[event.EventId]
}

// get active event given a time point
// if the active event doesn't contain the timePoint, then null is returned
func (ae *EventActive) GetActiveEvent(timePoint time.Time) *serverdata.EventActive {
	return ae.GetActiveEventUnix(timePoint.Unix())
}

func (ae *EventActive) GetActiveEventUnix(timeStamp int64) *serverdata.EventActive {
	event := ae.GetEventValue()
	if (event == nil) || (event.StartAt > timeStamp) || (event.EndAt < timeStamp) {
		return nil
	}
	return event
}

// TODO(extra): This is not the most elegant way to do siths
var eventActiveInstances = []*EventActive{}

func InvalidateActiveEvent() {
	for _, instance := range eventActiveInstances {
		instance.Event = nil
	}
}

func loadEventActive(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	eventActiveInstances = append(eventActiveInstances, &gamedata.EventActive)
	gamedata.EventActive.Gamedata = gamedata
	gamedata.EventActive.GetEventValue()
}

func init() {
	addLoadFunc(loadEventActive)
}
