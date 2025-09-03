package serverdata

import (
	"elichika/config"
	"elichika/db"
	"elichika/log"
	"elichika/serverstate"
	"elichika/utils"

	"os"

	"xorm.io/xorm"
)

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite", config.ServerdataPath)
	migrateToServerstate(engine)
	utils.CheckErr(err)
	engine.SetMaxOpenConns(50)
	engine.SetMaxIdleConns(10)
	rebuildAsset = (len(os.Args) >= 2) && (os.Args[1] == "rebuild_assets")
	InitTables()
	engine.Close()
	Database, err = db.NewDatabase(config.ServerdataPath)
	utils.CheckErr(err)
}

// TODO(finalise): Remove this in final versions
func migrateToServerstate(engine *xorm.Engine) {
	exist, err := engine.Table("s_event_active").IsTableExist("s_event_active")
	utils.CheckErr(err)
	if !exist {
		return
	}
	exist, err = engine.Table("s_event_scheduled").IsTableExist("s_event_scheduled")
	utils.CheckErr(err)
	if !exist {
		return
	}
	exist, err = engine.Table("s_scheduled_task").IsTableExist("s_scheduled_task")
	utils.CheckErr(err)
	if !exist {
		return
	}
	log.Println("Migrating data from serverdata into serverstate")
	// only when all 3 tables exist will we perform migration
	ea := []serverstate.EventActive{}
	es := []serverstate.EventScheduled{}
	st := []serverstate.ScheduledTask{}
	err = engine.Table("s_event_active").Find(&ea)
	utils.CheckErr(err)
	err = engine.Table("s_event_scheduled").Find(&es)
	utils.CheckErr(err)
	err = engine.Table("s_scheduled_task").Find(&st)
	utils.CheckErr(err)
	if len(ea)+len(es)+len(st) > 0 {
		// only when there's an active event will we migrate
		serverstate.Database.Do(func(session *xorm.Session) {
			// drop all scheduled tasks related to events
			good := true
			_, err := session.Table("s_scheduled_task").Where("task_name LIKE \"event_%%\"").Delete()
			utils.CheckErr(err)
			good = good && (err == nil)
			// then insert all the items
			if len(ea) > 0 {
				count, err := session.Table("s_event_active").Insert(ea)
				good = good && (err == nil) && (count == int64(len(ea)))
			}
			if len(es) > 0 {
				count, err := session.Table("s_event_scheduled").Insert(es)
				good = good && (err == nil) && (count == int64(len(es)))
			}
			if len(st) > 0 {
				count, err := session.Table("s_scheduled_task").Insert(st)
				good = good && (err == nil) && (count == int64(len(st)))
			}
			if good {
				session.Commit()
				log.Println("Migrated")
			} else {
				log.Println("Error migrating event data, dropping it")
				session.Rollback()
			}
		})
	}
	// drop all these tables no matter of it's good or not
	err = engine.DropTables("s_event_active")
	utils.CheckErr(err)
	err = engine.DropTables("s_event_scheduled")
	utils.CheckErr(err)
	err = engine.DropTables("s_scheduled_task")
	utils.CheckErr(err)
}
