package scheduled_task

import (
	"elichika/serverdata"
	"elichika/utils"

	"fmt"
	"time"

	"xorm.io/xorm"
)

// this module handle scheduled server sided tasks
// for example:
// - cleaning up outdated data
// - start / end event period
// - pay out rewards
// - other things
// because the server might not be up all the time, the tasks are actually handled in the following manner:
// - tasks are processed "occasionally":
//   - for now, everytime a proper request come in, we try to process tasks
//   - but we can also add server startup or regular interval or whatever
//
// - everytime tasks are processed, they are done from earliest to latest, if their timestamp is not exceeding the current timestamp
// - scheduled tasks happening at the same time are ordered by priority:
//   - lower priority number is processed first
//
// - if the schedule tasks have the same time and priority, then the order are not guaranteed
// - scheduled tasks are allowed to spawn new scheduled tasks:
//   - this can allow for repeating tasks
//   - or lead to multistage tasks.
//
// - scheduled tasks also take a string params, allowing for stuffs
type ScheduledTask = serverdata.ScheduledTask

// the general task take a session to the user database and a session to server database
type TaskHandler = func(*xorm.Session, *xorm.Session, ScheduledTask)

var taskHandlers = map[string]TaskHandler{}

func AddScheduledTaskHandler(taskName string, handler TaskHandler) {
	_, exist := taskHandlers[taskName]
	if exist {
		panic(fmt.Sprint("task already has handler: ", taskName))
	}
	taskHandlers[taskName] = handler
}

func AddScheduledTask(serverdata_db *xorm.Session, scheduledTask ScheduledTask) {
	_, err := serverdata_db.Table("s_scheduled_task").Insert(scheduledTask)
	utils.CheckErr(err)
}

func HandleScheduledTasks(serverdata_db *xorm.Session, userdata_db *xorm.Session, currentTime time.Time) {
	for {
		task := []ScheduledTask{}

		err := serverdata_db.Table("s_scheduled_task").OrderBy("time, priority").Limit(1).Find(&task)
		utils.CheckErr(err)
		if len(task) == 0 || (task[0].Time > currentTime.Unix()) {
			break
		}
		handler, exist := taskHandlers[task[0].TaskName]
		if !exist {
			fmt.Println("Warning: Ignored task with no handler: ", task[0].TaskName)
			continue
		}
		handler(serverdata_db, userdata_db, task[0])
		_, err = serverdata_db.Table("s_scheduled_task").Where("time = ? AND task_name = ? AND priority = ?",
			task[0].Time, task[0].TaskName, task[0].Priority).Delete(&task[0])
		utils.CheckErr(err)
		serverdata_db.Commit()
		serverdata_db.Begin()
	}
}
