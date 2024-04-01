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
)

func setScoreLive(ctx *gin.Context) {
	req := request.SetLiveRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
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
