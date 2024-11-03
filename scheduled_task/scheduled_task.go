package scheduled_task

import (
	"elichika/serverdata"
	"elichika/userdata/database"
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

// the general task take a session to the user database
type TaskHandler = func(*xorm.Session, ScheduledTask)

var taskHandlers = map[string]TaskHandler{}

func AddScheduledTaskHandler(taskName string, handler TaskHandler) {
	_, exist := taskHandlers[taskName]
	if exist {
		panic(fmt.Sprint("task already has handler: ", taskName))
	}
	taskHandlers[taskName] = handler
}

func AddScheduledTask(scheduledTask ScheduledTask) {
	var err error
	serverdata.Database.Do(func(session *xorm.Session) {
		_, err = session.Table("s_scheduled_task").Insert(scheduledTask)
		if err != nil {
			return
		}
		err = session.Commit()
		if err != nil {
			return
		}
		err = session.Begin()
	})
	utils.CheckErr(err)
}

func HandleScheduledTasks(userdata_db *xorm.Session, currentTime time.Time) {
	for {
		task := []ScheduledTask{}
		var err error
		serverdata.Database.Do(func(session *xorm.Session) {
			err = session.Table("s_scheduled_task").OrderBy("time, priority").Limit(1).Find(&task)
		})
		utils.CheckErr(err)
		if len(task) == 0 || (task[0].Time > currentTime.Unix()) {
			break
		}
		handler, exist := taskHandlers[task[0].TaskName]
		if !exist {
			fmt.Println("Warning: Ignored task with no handler: ", task[0].TaskName)
		} else {
			handler(userdata_db, task[0])
		}
		serverdata.Database.Do(func(session *xorm.Session) {
			_, err = session.Table("s_scheduled_task").Where("time = ? AND task_name = ? AND priority = ?",
				task[0].Time, task[0].TaskName, task[0].Priority).Delete(&task[0])
		})
		utils.CheckErr(err)
		serverdata.Database.Do(func(session *xorm.Session) {
			session.Commit()
			session.Begin()
		})
	}
}

// repeatedly select the tasks, check them, and run them if necessary
// this is useful to clear out some tasks early
// for each iteration, select at most 1 task to do and then continue to the next iteration
// - the check function should return whether a task should be run or not
// - it should also return if the force run should be terminated or not
// - if no task is run, then the termination happens anyway
// - note that if a task is run, then it is run first, then termination happen, if signified

func ForceRun(userdata_db *xorm.Session, check func(ScheduledTask) (bool, bool)) {
	ownedServerdataDb := (*xorm.Session)(nil)
	ownedUserdataDb := (*xorm.Session)(nil)
	defer func() {
		if ownedServerdataDb != nil {
			ownedServerdataDb.Close()
		}
		if ownedUserdataDb != nil {
			ownedUserdataDb.Close()
		}
	}()
	if userdata_db == nil {
		ownedUserdataDb = database.Engine.NewSession()
		userdata_db = ownedUserdataDb
	}
	for {
		tasks := []ScheduledTask{}
		var err error
		serverdata.Database.Do(func(session *xorm.Session) {
			err = session.Table("s_scheduled_task").OrderBy("time, priority").Find(&tasks)
		})
		utils.CheckErr(err)
		if len(tasks) == 0 {
			break
		}
		for _, task := range tasks {
			run, terminate := check(task)
			if run {
				handler, exist := taskHandlers[task.TaskName]
				if !exist {
					fmt.Println("Warning: Ignored task with no handler: ", task.TaskName)
				} else {
					handler(userdata_db, task)
				}
				serverdata.Database.Do(func(session *xorm.Session) {
					_, err = session.Table("s_scheduled_task").Where("time = ? AND task_name = ? AND priority = ?",
						task.Time, task.TaskName, task.Priority).Delete(&task)
				})
				utils.CheckErr(err)
				serverdata.Database.Do(func(session *xorm.Session) {
					session.Commit()
					session.Begin()
				})
			}
			if terminate {
				return
			}
			if run { // move on to the next iteration, we need to fetch the tasks again as some tasks might have been added
				break
			}
		}
	}
}
