package communication_member

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_live_difficulty"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func finishUserStoryMember(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryMemberRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	if session.FinishStoryMember(req.StoryMemberMasterId) {
		storyMemberMaster := gamedata.StoryMember[req.StoryMemberMasterId]
		if storyMemberMaster.Reward != nil {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *storyMemberMaster.Reward,
				PresentRouteType: enum.PresentRouteTypeStoryMember,
				PresentRouteId:   generic.NewNullable(req.StoryMemberMasterId),
			})
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeStoryMemberReward,
				ParamInt:        generic.NewNullable(req.StoryMemberMasterId),
			})
		}
		if storyMemberMaster.UnlockLiveId != nil {
			user_live_difficulty.UnlockLive(session, *storyMemberMaster.UnlockLiveId)
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/communicationMember/finishUserStoryMember", finishUserStoryMember)
}
