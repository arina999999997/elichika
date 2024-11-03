package event

import (
	"elichika/scheduled_task"
	"elichika/serverdata"

	"strings"
)

var isDirectEventChanging bool
var targetedEventId int32

// select the event directly instead of waiting for the scheduler
func ChangeEvent(eventId int32) {
	isDirectEventChanging = true
	targetedEventId = eventId

	scheduled_task.ForceRun(nil, func(task serverdata.ScheduledTask) (bool, bool) {
		if task.TaskName == "event_marathon_start" { // TODO(event): Add more event type
			return true, true // run the task and then stop
		}
		return strings.HasPrefix(task.TaskName, "event"), false // run the task if it's event related, to clean up existing events, then stop
	})
	isDirectEventChanging = false
}
