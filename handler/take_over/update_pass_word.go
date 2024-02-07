package take_over

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_pass_word"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func updatePassWord(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdatePassWordRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_pass_word.SetPassWord(session, req.PassWord)
	session.Finalize()

	common.JsonResponse(ctx, &response.UpdatePassWordResponse{
		TakeOverId: fmt.Sprint(userId),
	})
}

func init() {
	router.AddHandler("/takeOver/updatePassWord", updatePassWord)
}
