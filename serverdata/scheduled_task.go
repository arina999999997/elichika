package serverdata

import (
	"elichika/utils"

	"xorm.io/xorm"
)

type ScheduledTask struct {
	Time     int64  `xorm:"'time'"`
	TaskName string `xorm:"'task_name'"`
	Priority int32  `xorm:"'priority'"`
	Params   string `xorm:"'params'"`
}

func init() {
	addTable("s_scheduled_task", ScheduledTask{}, initScheduledTask)
}

func initScheduledTask(session *xorm.Session) {
	_, err := session.Table("s_scheduled_task").Insert(ScheduledTask{
		Time:     0,
		TaskName: "event_auto_scheduler",
		Priority: 0,
		Params:   "",
	})
	utils.CheckErr(err)
}
