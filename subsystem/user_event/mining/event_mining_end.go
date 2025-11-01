package mining

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/log"
	"elichika/scheduled_task"

	"strconv"

	"xorm.io/xorm"
)

// finish the event and pay out the reward for everyone who participated
func endEventScheduledHandler(userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	activeEvent := gamedata.Instance.EventActive.GetActiveEventUnix(task.Time)
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (activeEvent == nil) || (activeEvent.EventId != eventId) ||
		(activeEvent.EventType != enum.EventType1Mining) || (activeEvent.EndAt != task.Time) {
		log.Println("Warning: Failed to end event: ", task)
		return
	}
	// no actual clean up is necessary, we just need to remove the ranking object
	ResetVoltageRanking()
	ResetPointRanking()

	// TODO(event): Add config for other options once we have more than 1 event
	scheduled_task.AddScheduledTask(scheduled_task.ScheduledTask{
		Time:     activeEvent.EndAt + 1,
		TaskName: "event_auto_scheduler",
	})
}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_mining_end", endEventScheduledHandler)
}
