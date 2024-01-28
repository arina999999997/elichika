package story_event_history

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func UnlockStory(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UnlockStoryEventHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UnlockEventStory(req.EventStoryMasterId)
	user_content.RemoveContent(session, item.MemoryKey)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func FinishStory(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishStoryEventHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// there is no cleared tracking so all this request does is set story mode
	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {

	// /storyEventHistory/ all done
	router.AddHandler("/storyEventHistory/unlockStory", UnlockStory)
	router.AddHandler("/storyEventHistory/finishStory", FinishStory)

}
