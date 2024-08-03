package event

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/scheduled_task"
	"elichika/serverdata"
	"elichika/utils"
	"xorm.io/xorm"

	"fmt"
	"time"
)

// The auto scheduler works as follow:
// - The timeline is divided into rest period and event period
// - rest period means there's no event, not even event that has ended and is only showing results
// - the event period is further divided into:
//   - event period: from StartAt to ExpireAt, event points can be gained.
//   - tally period: from ExpiredAt to ResultAt, event points can no longer be gained and ranking can't be viewed.
//   - result period: from ResultAt to EndAt, ranking can be viewed and reward can be received
// - The scheduler take the duration of the above 4 periods along with 2 more variable:
//   - Last event ended at: the EndAt of the last event, default to 0
//   - Event period start time of day: The time of day when an event should start, default to 15:00
// - If there is no active event, the scheduler will schedule an event such that:
//   - The start time of day is respected
//   - There is at least the duration of the rest period between last event ended and new event start.
//   - If the start time must be in the future, then pick the first possible time.
//   - Otherwise pick the latest past time that is before current time.
// - The picked event is decided using the s_events table:
//   - Treat the events there as a circular array
//   - Based on the last event, pick the one after it
// TODO(config): For now these are limited to one of the premade config

// These constants exist for testing only, but they are also used for some constraint
const HourDuration = 60 * 60
const DayDuration = 24 * HourDuration
const ActualDayDuration = 24 * 60 * 60

type AutoSchedulerConfig struct {
	RestDuration   int32
	EventDuration  int32
	TallyDuration  int32
	ResultDuration int32

	TimeOfDayStartAt int32
}

var autoSchedulerConfigs = map[string]AutoSchedulerConfig{}

func init() {
	if ActualDayDuration%DayDuration != 0 {
		panic("DayDuration must be a divisor of ActualDayDuration")
	}
	// event cycle every day
	// used if user want to quickly cycle through the events to get all the rewards
	autoSchedulerConfigs["1_day"] = AutoSchedulerConfig{
		RestDuration:     HourDuration,
		EventDuration:    DayDuration - HourDuration*4,
		TallyDuration:    HourDuration,
		ResultDuration:   HourDuration * 2,
		TimeOfDayStartAt: 15 * HourDuration,
	}
	// event cycle every week, and there's no permanent rest
	// this is the default
	autoSchedulerConfigs["7_days"] = AutoSchedulerConfig{
		RestDuration:     HourDuration,      // 1 hour
		EventDuration:    DayDuration * 6,   // 6 days
		TallyDuration:    HourDuration * 3,  // 3 hours
		ResultDuration:   HourDuration * 20, // 20 hours
		TimeOfDayStartAt: 15 * HourDuration,
	}
}

func eventAutoScheduler(serverdata_db *xorm.Session, userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	// the scheduler use real time instead of scheduled time when schduling
	// but it will use scheduled time for checking
	// the auto scheduler should be invoked by the event end task or directly
	// trying to schedule an event while an other event hasn't ended will not work
	configObj := autoSchedulerConfigs[*config.Conf.EventAutoSchedulerPeriod]
	now := time.Now()

	year, month, day := now.Date()
	startTime := time.Date(year, month, day, 0, 0, int(configObj.TimeOfDayStartAt), 0, now.Location()).Unix()

	for startTime < now.Unix() {
		startTime += DayDuration
	}
	for startTime > now.Unix() {
		startTime -= DayDuration
	}

	lastEvent := gamedata.Instance.EventActive.GetEventValue()
	lastEndedAt := int64(0)
	if lastEvent != nil {
		if lastEvent.EndAt >= task.Time {
			fmt.Println("Warning: active event hasn't ended, event auto scheduler ignored")
			return
		}
		lastEndedAt = lastEvent.EndAt
	}

	for startTime < lastEndedAt+int64(configObj.RestDuration) {
		startTime += DayDuration
	}

	eventId := gamedata.Instance.EventAvailable.GetNextEvent(lastEvent)
	eventType := gamedata.Instance.GetEventType(eventId)

	// need to fill the delete condition with some stuff because of xorm
	_, err := serverdata_db.Table("s_event_active").Where("event_id >= 0").Delete(&serverdata.EventActive{})
	utils.CheckErr(err)

	_, err = serverdata_db.Table("s_event_active").Insert(serverdata.EventActive{
		EventId:   eventId,
		EventType: eventType,
		StartAt:   startTime,
		ExpiredAt: startTime + int64(configObj.EventDuration),
		ResultAt:  startTime + int64(configObj.EventDuration+configObj.TallyDuration),
		EndAt:     startTime + int64(configObj.EventDuration+configObj.TallyDuration+configObj.ResultDuration),
	})
	utils.CheckErr(err)
	gamedata.InvalidateActiveEvent()

	// schedule the event start at start time to truly begin the event
	if eventType == enum.EventType1Marathon {
		scheduled_task.AddScheduledTask(serverdata_db, serverdata.ScheduledTask{
			Time:     startTime,
			TaskName: "event_marathon_start",
			Priority: 0,
			Params:   fmt.Sprint(eventId),
		})
	} else {
		panic("unsupported event type")
	}
}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_auto_scheduler", eventAutoScheduler)
}
