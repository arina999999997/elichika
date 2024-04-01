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

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_authentication.SetPassWord(session, req.PassWord)
	session.Finalize()

	common.JsonResponse(ctx, &response.UpdatePassWordResponse{
		TakeOverId: fmt.Sprint(userId),
	})
}

func init() {
	router.AddHandler("/takeOver/updatePassWord", updatePassWord)
}
