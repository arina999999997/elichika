package mining

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/log"
	"elichika/scheduled_task"
	"elichika/utils"

	"fmt"
	"strconv"

	"xorm.io/xorm"
)

func StartEventMining(userdata_db *xorm.Session, eventId int32) {
	// Start the event.
	// This is only done once per event.
	// Because the event can be reused, this involve clearing out all the old record and trigger and stuff
	// The story progress will be kept
	_, err := userdata_db.Exec(fmt.Sprintf("UPDATE u_event_mining SET event_point = 0 WHERE event_master_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_event_status WHERE event_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMiningFirstRuleDescription, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMiningShowResult, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec("DELETE FROM u_event_mining_trade_product")
	utils.CheckErr(err)
	_, err = userdata_db.Exec("DELETE FROM u_event_mining_music_best_score")
	utils.CheckErr(err)
}

func startEventScheduledHandler(userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	activeEvent := gamedata.Instance.EventActive.GetActiveEventUnix(task.Time)
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (activeEvent == nil) || (activeEvent.EventId != eventId) ||
		(activeEvent.EventType != enum.EventType1Mining) || (activeEvent.StartAt != task.Time) {

		log.Println("Warning: Failed to start event: ", task)
		log.Println((activeEvent == nil), (activeEvent.EventId != eventId), (activeEvent.EventType != enum.EventType1Mining), (activeEvent.StartAt != task.Time))
		return
	}
	// this will be scheduled by an event scheduler, and called when the event is ready to start
	StartEventMining(userdata_db, eventId)
	scheduled_task.AddScheduledTask(scheduled_task.ScheduledTask{
		Time:     activeEvent.ResultAt,
		TaskName: "event_mining_result",
		Params:   task.Params,
	})
}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_mining_start", startEventScheduledHandler)
}
