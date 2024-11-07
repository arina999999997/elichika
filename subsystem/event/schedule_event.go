package event

import (
	"elichika/scheduled_task"
	"elichika/serverdata"
	"elichika/utils"

	"strings"

	"xorm.io/xorm"
)

func ScheduleEvent(eventId int32) {
	var err error
	var affected int64
	defer func() {
		{
			err := recover()
			if err != nil {
				panic(err)
			}
		}
		serverdata.Database.Do(func(session *xorm.Session) {
			err = session.Commit()
		})
		utils.CheckErr(err)
		serverdata.Database.Do(func(session *xorm.Session) {
			err = session.Begin()
		})
		utils.CheckErr(err)
	}()
	event := serverdata.EventScheduled{
		EventId: eventId,
	}
	serverdata.Database.Do(func(session *xorm.Session) {
		affected, err = session.Table("s_event_scheduled").Update(&event)
	})
	utils.CheckErr(err)
	if affected == 0 {
		serverdata.Database.Do(func(session *xorm.Session) {
			_, err = session.Table("s_event_scheduled").Insert(&event)
		})
		utils.CheckErr(err)
	}

	tasks := []serverdata.ScheduledTask{}
	serverdata.Database.Do(func(session *xorm.Session) {
		err = session.Table("s_scheduled_task").Find(&tasks)
	})
	utils.CheckErr(err)
	for _, task := range tasks {
		if strings.HasPrefix(task.TaskName, "event_") {
			return
		}
	}
	// no event, so we add a new task
	scheduled_task.AddScheduledTask(serverdata.ScheduledTask{
		Time:     0,
		TaskName: "event_auto_scheduler",
		Priority: 0,
		Params:   "",
	})
}
