package main

import (
	"elichika/config"
	_ "elichika/handler"
	_ "elichika/subsystem"
	"elichika/userdata"
	_ "elichika/webui"

	"elichika/router"

	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func cli() {
	fmt.Println("CLI is reserved for special behaviour, the server will now exit, start it again without any argument!")
}

func main() {
	if len(os.Args) > 1 {
		cli()
		return
	}
	userdata.Init()
	runtime.GC()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)
	fmt.Println("server address: ", *config.Conf.ServerAddress)
	fmt.Println("WebUI address: ", *config.Conf.ServerAddress+"/webui")
	r.Run(*config.Conf.ServerAddress)
}
