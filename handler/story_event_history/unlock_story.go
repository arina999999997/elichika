package story_event_history

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_story_event_history"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func unlockStory(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UnlockStoryEventHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_story_event_history.UnlockEventStory(session, req.EventStoryMasterId)
	user_content.RemoveContent(session, item.MemoryKey)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/storyEventHistory/unlockStory", unlockStory)
}
