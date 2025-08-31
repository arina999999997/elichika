package clientdb

import (
	"elichika/config"
	"elichika/log"
	"elichika/utils"

	"bufio"
	"fmt"
	"os"
	"os/exec"

	"xorm.io/xorm"
)

func isNotChanged(file string) bool {
	cmd := exec.Command("git", "diff", "--exit-code", "--quiet", file)
	cmd.Dir = config.AssetPath
	err := cmd.Run()
	if err == nil {
		return true // exit code is 0
	}
	exitError, ok := err.(*exec.ExitError)
	if !ok {
		panic(err)
	}
	if exitError.ExitCode() != 1 {
		panic(err)
	}
	return false
}

// note that this is subject to change, do not depend on it too much
func initLocale(locale string) {

	sqlDir := fmt.Sprint(config.AssetPath, "sql/", locale, "/")
	dbDir := fmt.Sprint("db/", locale, "/")
	files, err := os.ReadDir(sqlDir)
	if err != nil {
		return
	}
	// file name has the format <order>.filename.sql
	// order must be exactly 3 digits (technically it can be any 3 characters)
	// for each file, if it has not changed then apply the update
	// if an error is encountered, no change would be made to any of the file
	needUpdate := map[string]bool{}
	engines := map[string]*xorm.Engine{}
	sessions := map[string]*xorm.Session{}

	for _, file := range files {
		dbName := file.Name()[4 : len(file.Name())-4]
		need, exists := needUpdate[dbName]
		if !exists {
			needUpdate[dbName] = isNotChanged(dbDir + dbName)
			need = needUpdate[dbName]
		}
		if !need {
			continue
		}
		session, exists := sessions[dbName]
		if !exists {
			engines[dbName], err = xorm.NewEngine("sqlite", config.AssetPath+dbDir+dbName)
			utils.CheckErr(err)
			engines[dbName].SetMaxOpenConns(50)
			engines[dbName].SetMaxIdleConns(10)
			sessions[dbName] = engines[dbName].NewSession()
			session = sessions[dbName]
			session.Begin()
		}
		log.Println("Running SQL file: ", file.Name())

		f, err := os.Open(sqlDir + file.Name())
		utils.CheckErr(err)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			_, err = session.Exec(scanner.Text())
			utils.CheckErr(err)

		}
		utils.CheckErr(scanner.Err())
	}
	for _, session := range sessions {
		err := session.Commit()
		utils.CheckErr(err)
		session.Close()
	}
}

// initialise the database inside of the db repository, if necessary
func databaseInit() {
	initLocale("gl")
	initLocale("jp")
}
