package marathon

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/scheduled_task"

	"fmt"
	"strconv"

	"xorm.io/xorm"
)

// finish the event and pay out the reward for everyone who participated
func endEventScheduledHandler(serverdata_db *xorm.Session, userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	activeEvent := gamedata.Instance.EventActive.GetActiveEventUnix(task.Time)
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (activeEvent == nil) || (activeEvent.EventId != eventId) ||
		(activeEvent.EventType != enum.EventType1Marathon) || (activeEvent.EndAt != task.Time) {
		fmt.Println("Warning: Failed to end event: ", task)
		return
	}
	// no actual clean up is necessary, we just need to remove the ranking object
	ResetRanking()

	// TODO(event): Add config for other options once we have more than 1 event
	scheduled_task.AddScheduledTask(serverdata_db, scheduled_task.ScheduledTask{
		Time:     activeEvent.EndAt + 1,
		TaskName: "event_auto_scheduler",
	})
}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_marathon_end", endEventScheduledHandler)
}
