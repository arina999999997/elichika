package serverstate

// serverstate used to be part of serverdata, as we moved to support embedded elichika, it need to be separated as embedded version cannot properly generate serverdata.db
import (
	"elichika/db"
	"elichika/log"
	"elichika/utils"

	"xorm.io/xorm"
)

type Initializer = func(*xorm.Session)

var (
	engine                            *xorm.Engine
	Database                          *db.DatabaseSync
	serverstateTableNameToInterface   = map[string]interface{}{}
	serverstateTableNameToInitializer = map[string]Initializer{}

	// whether to reset the server state
	// this do not delete any user data, it only reset the server to the initial state
	// can be used if the server sided tasks are somehow in a bad state
	// this will almost certainly disrupt on-going events
	resetServer bool
)

func addTable(tableName string, structure interface{}, initializer Initializer) {
	_, exist := serverstateTableNameToInterface[tableName]
	if exist {
		log.Panic("table already exist: " + tableName)
	}
	serverstateTableNameToInterface[tableName] = structure
	serverstateTableNameToInitializer[tableName] = initializer
}

func createTable(tableName string, structure interface{}, overwrite bool) bool {
	exist, err := engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)

	if !exist {
		log.Println("Creating new table:", tableName)
		err = engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
		return true
	} else {
		if !overwrite {
			return false
		}
		log.Println("Overwrite existing table:", tableName)
		err := engine.DropTables(tableName)
		utils.CheckErr(err)
		err = engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
		return true
	}
}

func isTableEmpty(tableName string) bool {
	total, err := engine.Table(tableName).Count()
	utils.CheckErr(err)
	return total == 0
}

func InitTables() {
	initializers := []Initializer{}
	for tableName := range serverstateTableNameToInterface {
		newOrEmpty := createTable(tableName, serverstateTableNameToInterface[tableName], resetServer)
		newOrEmpty = newOrEmpty || isTableEmpty(tableName)
		if newOrEmpty {
			initializers = append(initializers, serverstateTableNameToInitializer[tableName])
		}
	}
	session := engine.NewSession()
	defer session.Close()
	session.Begin()
	for _, initializer := range initializers {
		if initializer == nil {
			continue
		}
		initializer(session)
	}
	session.Commit()
}
