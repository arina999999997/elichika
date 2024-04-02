package webui

import (
	"elichika/config"
	"elichika/utils"
	"elichika/webui/object_form"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigEditor(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	starts := "<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"/></head><div>Update server runtime config. Note that some configurations will be applied right away, some will requires restarting the server.</div>"

	ctx.String(http.StatusOK, "%s\n%s", starts,
		object_form.GenerateWebForm(config.Conf, "/update_config", "Reset to current config", "Update config"))
}

func UpdateConfig(ctx *gin.Context) {
	newConfig := config.RuntimeConfig{}
	err := object_form.ParseForm(ctx, &newConfig)
	utils.CheckErr(err)
	config.UpdateConfig(&newConfig)
	ctx.Redirect(http.StatusFound, commonPrefix+"Updated runtime config")

}
