package admin

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/webui_utils"

	"encoding/base64"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	var respString string
	resp := webui_utils.Response{}
	form := ctx.MustGet("form").(*multipart.Form)

	// TODO: this is vulnerable to timing attack but it's whatever
	adminPassword := form.Value["admin_password"][0]
	if *config.Conf.AdminPassword != adminPassword {
		resp.Error = &respString
		*resp.Error = "Wrong password!"
	} else {
		newSessionKey()
		resp.Response = &respString
		*resp.Response = base64.StdEncoding.EncodeToString(adminSessionKey)
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	router.AddSpecialSetup("/webui/admin", func(group *gin.RouterGroup) {
		group.StaticFile("/login", "./webui/admin/login.html")
	})
	router.AddHandler("/webui/admin", "POST", "/login", login)
}
