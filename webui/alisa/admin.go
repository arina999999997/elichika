package main

import (
	"elichika/config"
	"elichika/utils"
	"elichika/webui/webui_utils"

	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var adminSessionKey []byte

func randomKey() []byte {
	// random 32 bytes
	b := make([]byte, 32)
	_, err := rand.Read(b)
	utils.CheckErr(err)
	return b
}

func newSessionKey() {
	adminSessionKey = randomKey()
}

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

func login(ctx *gin.Context) {
	var respString string
	resp := webui_utils.Response{}
	form := ctx.MustGet("form").(*multipart.Form)

	adminPassword := form.Value["admin_password"][0]
	if subtle.ConstantTimeCompare([]byte(*config.Conf.AdminPassword), []byte(adminPassword)) == 1 {
		newSessionKey()
		resp.Response = &respString
		*resp.Response = base64.StdEncoding.EncodeToString(adminSessionKey)
	} else {
		resp.Error = &respString
		*resp.Error = "Wrong password!"
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}
