package communication_member

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_live_difficulty"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_story_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func finishUserStoryMember(ctx *gin.Context) {
	req := request.FinishUserStoryMemberRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	storyMemberMaster := session.Gamedata.StoryMember[req.StoryMemberMasterId]
	if user_story_member.FinishStoryMember(session, req.StoryMemberMasterId) {
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
	}
	// always try to unlock the live
	if storyMemberMaster.UnlockLiveId != nil {
		user_live_difficulty.UnlockLive(session, *storyMemberMaster.UnlockLiveId)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/communicationMember/finishUserStoryMember", finishUserStoryMember)
}
