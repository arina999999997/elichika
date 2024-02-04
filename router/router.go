package router

import (
	// "elichika/handler"
	// "elichika/handler/live"

	"elichika/middleware"
	"elichika/webui"

	"github.com/gin-gonic/gin"
)

// Other packages should import the router package and declare which API it want to handle
// In main.go, import the relevant packages to support them
type Handler = func(*gin.Context)

var (
	handlers = map[string]Handler{}
)

func AddHandler(path string, handler Handler) {
	_, exist := handlers[path]
	if exist {
		panic("Multiple handler for path: " + path)
	}
	handlers[path] = handler
}

func Router(r *gin.Engine) {
	r.Static("/static", "static")
	{
		api := r.Group("/", middleware.Common)
		for path, handler := range handlers {
			api.POST(path, handler)
		}
	}

	{
		webapi := r.Group("/webui", webui.Common)
		r.Static("/webui", "webui")
		// the web ui cover for functionality that can't be done by the client or is currently missing
		webapi.POST("/birthday", webui.Birthday)
		webapi.POST("/accessory", webui.Accessory)
		webapi.POST("/import_account", webui.ImportAccount)
		webapi.POST("/export_account", webui.ExportAccount)
		webapi.POST("/reset_story_main", webui.ResetProgress)
		webapi.POST("/reset_story_side", webui.ResetProgress)
		webapi.POST("/reset_story_member", webui.ResetProgress)
		webapi.POST("/reset_story_linkage", webui.ResetProgress)
		webapi.POST("/reset_story_event", webui.ResetProgress)
		webapi.POST("/reset_dlp", webui.ResetProgress)
		webapi.POST("/add_present", webui.AddPresent)
	}
}
