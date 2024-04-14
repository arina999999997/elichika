package user

import (
	"elichika/router"
	"elichika/subsystem/user_authentication"
	"elichika/userdata"
	"elichika/utils"

	"encoding/base64"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var respString string
	resp := Response{}

	form := ctx.MustGet("form").(*multipart.Form)

	userPassword := form.Value["user_password"][0]

	session := ctx.MustGet("session").(*userdata.Session)

	if session == nil {
		resp.Error = &respString
		*resp.Error = "User doesn't exist"
	} else if !user_authentication.CheckPassWord(session, userPassword) {
		resp.Error = &respString
		*resp.Error = "Wrong password"
	} else {
		// invalidate existing sessions
		session.GenerateNewSessionKey()
		session.Finalize()

		resp.Response = &respString
		*resp.Response = base64.StdEncoding.EncodeToString(session.SessionKey())
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	router.AddSpecialSetup("/webui/user", func(group *gin.RouterGroup) {
		group.StaticFile("/login", "./webui/user/login.html")
	})
	router.AddHandler("/webui/user", "POST", "/login", Login)
}
