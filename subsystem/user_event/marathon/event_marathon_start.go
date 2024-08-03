package marathon

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/scheduled_task"
	"elichika/utils"

	"fmt"
	"strconv"

	"xorm.io/xorm"
)

func StartEventMarathon(userdata_db *xorm.Session, eventId int32) {
	// Start the event.
	// This is only done once per event.
	// Because the event can be reused, this involve clearing out all the old record and trigger and stuff
	// The story progress will be kept
	_, err := userdata_db.Exec(fmt.Sprintf("UPDATE u_event_marathon SET event_point = 0 WHERE event_master_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_event_status WHERE event_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMarathonFirstRuleDescription, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMarathonShowResult, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_content WHERE content_type = %d AND content_id = %d",
		enum.ContentTypeEventMarathonBooster, gamedata.Instance.EventMarathon[eventId].BoosterItemId))
	utils.CheckErr(err)
}

func startEventScheduledHandler(serverdata_db *xorm.Session, userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	activeEvent := gamedata.Instance.EventActive.GetActiveEventUnix(task.Time)
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (activeEvent == nil) || (activeEvent.EventId != eventId) ||
		(activeEvent.EventType != enum.EventType1Marathon) || (activeEvent.StartAt != task.Time) {

		fmt.Println("Warning: Failed to start event: ", task)
		fmt.Println((activeEvent == nil), (activeEvent.EventId != eventId), (activeEvent.EventType != enum.EventType1Marathon), (activeEvent.StartAt != task.Time))
		return
	}
	// this will be scheduled by an event scheduler, and called when the event is ready to start
	StartEventMarathon(userdata_db, eventId)

	// schedule the event payout and stuff
	scheduled_task.AddScheduledTask(serverdata_db, scheduled_task.ScheduledTask{
		Time:     activeEvent.ResultAt,
		TaskName: "event_marathon_result",
		Params:   task.Params,
	})
}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_marathon_start", startEventScheduledHandler)
}
