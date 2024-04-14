package story

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_story_linkage"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func finishStoryLinkage(ctx *gin.Context) {
	req := request.AddStoryLinkageRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	user_story_linkage.InsertUserStoryLinkage(session, req.CellId)

	common.JsonResponse(ctx, &response.AddStoryLinkageResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/story/finishStoryLinkage", finishStoryLinkage)
}
