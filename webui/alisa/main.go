package main

import (
	"elichika/config"
	"elichika/utils"

	"html/template"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

var exitChannel = make(chan struct{})
var runElichikaNext = false

func runElichika() {
	runElichikaNext = true
	exitChannel <- struct{}{}
}

func main() {
	newSessionKey()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// every request by client and elsewhere get hit with the maintenance
	// maybe it's better to filter out but oh well
	router.NoRoute(maintenance)

	router.StaticFile("/favicon.ico", "./webui/favicon.ico")

	funcs := template.FuncMap{}
	funcs["noescape"] = func(s string) template.HTML {
		return template.HTML(s)
	}

	router.SetFuncMap(funcs)
	router.LoadHTMLFiles("./webui/admin/logged_in_admin.html")

	group := router.Group("/webui/admin", adminInitial)

	group.GET("/", Index)
	group.StaticFile("/login", "./webui/admin/login.html")
	group.POST("/login", login)
	group.POST("/update_elichika", UpdateElichika)
	group.POST("/run_elichika", RunElichika)
	group.POST("/rebuild_elichika", RebuildElichika)

	server := &http.Server{
		Addr:    *config.Conf.ServerAddress,
		Handler: router.Handler(),
	}

	go func() {
		for i := 0; i < 10; i++ {
			err := server.ListenAndServe()
			if err == http.ErrServerClosed {
				return
			}
			if err != nil {
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}()

	<-exitChannel
	if runElichikaNext {
		cmd := exec.Command("./elichika")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		utils.CheckErr(err)
		err = cmd.Process.Release()
		utils.CheckErr(err)
	}
}
