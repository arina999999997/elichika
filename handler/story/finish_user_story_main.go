package story

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_story_main"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func finishUserStoryMain(ctx *gin.Context) {
	req := request.StoryMainRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	resp := response.StoryMainResponse{
		UserModelDiff: &session.UserModel,
	}

	if user_story_main.InsertUserStoryMain(session, req.CellId) { // newly inserted story, award some gem
		resp.FirstClearReward.Append(item.StarGem.Amount(10))
		user_present.AddPresent(session, client.PresentItem{
			Content:          item.StarGem.Amount(10),
			PresentRouteType: enum.PresentRouteTypeStoryMain,
			PresentRouteId:   generic.NewNullable(req.CellId),
		})
	}
	if req.MemberId.HasValue { // has a member -> select member thingy
		user_story_main.UpdateUserStoryMainSelected(session, req.CellId, req.MemberId.Value)
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/story/finishUserStoryMain", finishUserStoryMain)
}
