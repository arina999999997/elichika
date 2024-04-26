package admin

import (
	"elichika/router"
	"elichika/utils"

	"bytes"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

func adminInitial(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		form, err := ctx.MultipartForm()
		utils.CheckErr(err)
		ctx.Set("form", form)
		if !strings.HasPrefix(ctx.Request.URL.String(), "/webui/admin/login") {
			sessionKey, err := base64.StdEncoding.DecodeString(form.Value["admin_session_key"][0])
			utils.CheckErr(err)
			if !bytes.Equal(sessionKey, adminSessionKey) {
				panic("wrong session key")
			}
		}
	}
	ctx.Next()
}

func init() {
	router.AddInitialHandler("/webui/admin", adminInitial)
}
