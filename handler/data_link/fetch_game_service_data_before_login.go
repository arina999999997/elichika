package data_link

import (
	// "elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	// "elichika/handler/login"
	"elichika/locale"
	"elichika/router"
	// "elichika/userdata"
	"elichika/utils"

	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this is ios exclusive, this doesn't seems to be necessary at all
func fetchGameServiceDataBeforeLogin(ctx *gin.Context) {
	req := request.FetchGameServiceDataBeforeLoginRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)
	// fmt.Println(req)
	resp := response.FetchGameServiceDataBeforeLoginResponse{}

	// TODO(authentication): probably want to check against service id
	// session := userdata.GetSession(ctx, req.UserId)
	// defer session.Close()

	// if session != nil {
	// 	// fill in the data if user exists
	// 	resp.Data.LinkedData = client.UserLinkData{
	// 		UserId:            session.UserId,
	// 		AuthorizationKey:  login.LoginSessionKey(req.Mask),
	// 		Name:              session.UserStatus.Name,
	// 		LastLoginAt:       session.UserStatus.LastLoginAt,
	// 		SnsCoin:           session.UserStatus.FreeSnsCoin + session.UserStatus.AppleSnsCoin + session.UserStatus.GoogleSnsCoin,
	// 		TermsOfUseVersion: session.UserStatus.TermsOfUseVersion,
	// 	}
	// 	resp.Data.CurrentData = client.CurrentUserData{
	// 		UserId:      session.UserId,
	// 		Name:        session.UserStatus.Name,
	// 		LastLoginAt: session.UserStatus.LastLoginAt,
	// 		SnsCoin:     session.UserStatus.FreeSnsCoin + session.UserStatus.AppleSnsCoin + session.UserStatus.GoogleSnsCoin,
	// 	}
	// }

	respBody, _ := json.Marshal(resp)
	signedResp := common.SignResp(ctx, string(respBody), ctx.MustGet("locale").(*locale.Locale).StartupKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, signedResp)
}

func init() {
	router.AddHandler("/", "POST", "/dataLink/fetchGameServiceDataBeforeLogin", fetchGameServiceDataBeforeLogin)
}
