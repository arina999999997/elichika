package take_over

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/handler/login"
	"elichika/locale"
	"elichika/router"
	"elichika/subsystem/user_account"
	"elichika/subsystem/user_pass_word"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func setTakeOver(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetTakeOverRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	_linkedUserId, err := strconv.Atoi(req.TakeOverId)
	utils.CheckErr(err)
	linkedUserId := int32(_linkedUserId)
	linkedSession := userdata.GetSession(ctx, linkedUserId)
	defer linkedSession.Close()

	if linkedSession == nil { // new account
		user_account.CreateNewAccount(ctx, linkedUserId, req.PassWord)
		linkedSession = userdata.GetSession(ctx, linkedUserId)
		defer linkedSession.Close()
	} else if !user_pass_word.CheckPassWord(linkedSession, req.PassWord) {
		panic("wrong pass word")
	}

	resp := response.SetTakeOverResponse{
		Data: client.UserLinkData{
			UserId:            int32(linkedSession.UserId),
			AuthorizationKey:  login.StartupAuthorizationKey(req.Mask),
			Name:              linkedSession.UserStatus.Name,
			LastLoginAt:       linkedSession.UserStatus.LastLoginAt,
			SnsCoin:           linkedSession.UserStatus.FreeSnsCoin + linkedSession.UserStatus.AppleSnsCoin + linkedSession.UserStatus.GoogleSnsCoin,
			TermsOfUseVersion: linkedSession.UserStatus.TermsOfUseVersion,
		},
	}

	respBody, _ := json.Marshal(resp)
	signedResp := common.SignResp(ctx, string(respBody), ctx.MustGet("locale").(*locale.Locale).StartupKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, signedResp)
}

func init() {
	router.AddHandler("/takeOver/setTakeOver", setTakeOver)
}
