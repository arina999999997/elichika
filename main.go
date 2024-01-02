package main

import (
	"elichika/config"
	"elichika/router"

	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func cli() {
	fmt.Println("Note: cli is not stable and should only be used for testing, use at your own risk!")
	fmt.Println("Note: If you want to do modification that can't be done in game, use the webUI: <your_server>/webui")
}

func main() {

	if len(os.Args) > 1 {
		cli()
		return
	}
	runtime.GC()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)
	fmt.Println("server address: ", *config.Conf.ServerAddress)
	fmt.Println("WebUI address: ", *config.Conf.ServerAddress+"/webui")
	r.Run(*config.Conf.ServerAddress) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
