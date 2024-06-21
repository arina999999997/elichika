package admin

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/object_form"
	"elichika/webui/webui_utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigEditor(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	starts := `<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"/></head>
	<div>Update server runtime config.</div>
	<div>Note that some configurations will be applied right away, some will requires restarting the server.</div>
	<div>Finally, you can always delete the config.json to reset everything to default.</div>
	`

	ctx.HTML(http.StatusOK, "logged_in_admin.html", gin.H{
		"body": starts + object_form.GenerateWebForm(config.Conf, "config_form", `onclick="submit_form('config_form', './config_editor')"`, "Reset to current config", "Update config"),
	})
}

func UpdateConfig(ctx *gin.Context) {
	newConfig := config.RuntimeConfig{}
	err := object_form.ParseForm(ctx, &newConfig)
	utils.CheckErr(err)
	config.UpdateConfig(&newConfig)
	webui_utils.CommonResponse(ctx, "Config updated, some changes will require a server restart to work.", "")
}

func init() {
	// TODO(admin): this is the only admin feature for now, so we let it be the main page
	router.AddHandler("/webui/admin", "GET", "/", ConfigEditor)
	router.AddHandler("/webui/admin", "POST", "/config_editor", UpdateConfig)
}
