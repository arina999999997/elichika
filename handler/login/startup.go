package login

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/locale"
	"elichika/router"
	"elichika/subsystem/user_account"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func startup(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.StartupRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	resp := response.StartupResponse{}
	resp.UserId = int32(user_account.CreateNewAccount(ctx, -1, ""))
	resp.AuthorizationKey = StartupAuthorizationKey(req.Mask)
	// note that this use a different key than the common one
	startupBody, _ := json.Marshal(resp)
	respBody := common.SignResp(ctx, string(startupBody), ctx.MustGet("locale").(*locale.Locale).StartupKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, respBody)
}

func init() {
	router.AddHandler("/login/startup", startup)
}
