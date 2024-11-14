package main

import (
	"elichika/webui/webui_utils"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	body := `<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"/></head>
	`
	version, err := getVersion()
	if err != nil {
		body += fmt.Sprint(`<div>Error getting version: `, err, `</div>`)
	} else {
		body += fmt.Sprintf(`<div>Current elichika version: %s </div>`, version)
	}
	body += `
	<div>Select Run elichika to run elichika, the admin webui will temporarily be disrupted.</div>
	<div>Select Update elichika to update elichika. The update will rebuild the executable too.</div>
	<div><input type="button" value="Run elichika" onclick="submit_form(null, './run_elichika')"/></div>
	<div><input type="button" value="Update elichika" onclick="submit_form(null, './update_elichika')"/></div>
	`
	// body += `<div><input type="button" value="Rebuild elichika" onclick="submit_form(null, './rebuild_elichika')"/></div>
	// `
	ctx.HTML(http.StatusOK, "logged_in_admin.html", gin.H{
		"body": body,
	})
}

func RunElichika(ctx *gin.Context) {
	webui_utils.CommonResponse(ctx, "Trying to run elichika! Note that the maintainance server will come offline, you can access the admin webui once elichika is up!", "")
	runElichika()
}

func UpdateElichika(ctx *gin.Context) {
	message, err := updateElichika()
	if len(message) > 1024 {
		message = message[:1024]
	}
	if err == nil {
		webui_utils.CommonResponse(ctx, message, "")
	} else {
		webui_utils.CommonResponse(ctx, fmt.Sprint(err), "")
	}
}

func RebuildElichika(ctx *gin.Context) {
	message, err := rebuildElichika()
	if len(message) > 1024 {
		message = message[:1024]
	}
	if err == nil {
		webui_utils.CommonResponse(ctx, message, "")
	} else {
		webui_utils.CommonResponse(ctx, fmt.Sprint(err), "")
	}
}
