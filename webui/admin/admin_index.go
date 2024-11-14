package admin

import (
	"elichika/router"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	starts := `<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"/></head>
	<div><button onclick="location.href='config_editor'" type="button">
			Config Editor</button></div>
	<div><button onclick="location.href='event_selector'" type="button">
			Event Selector</button></div>
	<div><button onclick="location.href='event_scheduler'" type="button">
			Event Scheduler</button></div>
	<div><button onclick="location.href='maintenance_mode'" type="button">
			Maintenance Mode (for updating)</button></div>
	`

	ctx.HTML(http.StatusOK, "logged_in_admin.html", gin.H{
		"body": starts,
	})
}

func init() {
	router.AddHandler("/webui/admin", "GET", "/", Index)
}
