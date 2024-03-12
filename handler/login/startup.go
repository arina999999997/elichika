package login

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/locale"
	"elichika/router"
	"elichika/subsystem/user_account"
	// "elichika/subsystem/user_authentication"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	// "net/http"

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

	session := userdata.GetSession(ctx, resp.UserId)
	defer session.Close()
	resp.AuthorizationKey = session.EncodedAuthorizationKey(req.Mask)
	// note that this use a different key than the common one
	ctx.Set("sign_key", ctx.MustGet("locale").(*locale.Locale).StartupKey)
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/login/startup", startup)
}
