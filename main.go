package main

import (
	"elichika/config"
	_ "elichika/handler"
	"elichika/log"
	"elichika/router"
	"elichika/shutdown"
	_ "elichika/subsystem"
	"elichika/subsystem/user_training_tree"
	"elichika/userdata"
	_ "elichika/webui"

	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

// with some cli, we keep the server open
// return true to keep open
func checkCli() bool {
	if os.Args[1] == "rebuild_assets" {
		if len(os.Args) > 2 && os.Args[2] == "keep_alive" {
			return true
		}
	}
	if os.Args[1] == "fix_training_trees" {
		user_training_tree.FixUsersTrainingTrees()
	}
	log.Println("CLI is reserved for special behaviour, the server will now exit, start it again without any argument!")
	return false
}

func main() {
	if len(os.Args) > 1 {
		if !checkCli() {
			return
		}
	}
	userdata.Init()
	runtime.GC()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)
	log.Println("server address: ", *config.Conf.ServerAddress)
	log.Println("WebUI address: ", *config.Conf.ServerAddress+"/webui/...")
	go func() {
		r.Run(*config.Conf.ServerAddress)
	}()
	shutdown.ReceiveFinalSignal()
}
