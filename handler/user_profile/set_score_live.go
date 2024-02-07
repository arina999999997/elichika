package user_profile

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_profile"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func setScoreLive(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_profile.SetScoreLive(session, req.LiveDifficultyMasterId)

	session.Finalize()
	common.JsonResponse(ctx, response.SetLiveResponse{
		LiveDifficultyMasterId: req.LiveDifficultyMasterId,
	})
}

func init() {
	router.AddHandler("/userProfile/setScoreLive", setScoreLive)
}
