package take_over

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_authentication"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func updatePassWord(ctx *gin.Context) {
	req := request.UpdatePassWordRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_authentication.SetPassWord(session, req.PassWord)

	common.JsonResponse(ctx, &response.UpdatePassWordResponse{
		TakeOverId: fmt.Sprint(session.UserId),
	})
}

func init() {
	router.AddHandler("/takeOver/updatePassWord", updatePassWord)
}
