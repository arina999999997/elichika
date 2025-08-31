//go:build embedded

package main

import (
	"elichika/config"
	_ "elichika/handler"
	"elichika/log"
	"elichika/router"
	_ "elichika/subsystem"
	"elichika/userdata"
	_ "elichika/webui"

	"github.com/gin-gonic/gin"

	"C"
)

func init() {
	log.Println("Entered main init")
	log.Println("Start loading userdata")
	userdata.Init()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	router.Router(r)
	log.Println("server address: ", *config.Conf.ServerAddress)
	log.Println("WebUI address: ", *config.Conf.ServerAddress+"/webui/...")
	go func() {
		r.Run(*config.Conf.ServerAddress)
	}()
	log.Println("Exit main init, server should be live")
}

func main() {}
