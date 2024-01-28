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
	"elichika/subsystem/time"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchCommunicationMemberDetail(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchCommunicationMemberDetailRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchCommunicationMemberDetailResponse{}
	resp.MemberLovePanels.Append(session.GetMemberLovePanel(req.MemberId))

	resp.WeekdayState = time.GetWeekdayState(session)
	common.JsonResponse(ctx, resp)
}

func UpdateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateMemberDetailBadgeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	detailBadge := user_member.GetUserCommunicationMemberDetailBadge(session, req.MemberMasterId)
	switch req.CommunicationMemberDetailBadgeType {
	case enum.CommunicationMemberDetailBadgeTypeStoryMember:
		detailBadge.IsStoryMemberBadge = false
	case enum.CommunicationMemberDetailBadgeTypeStorySide:
		detailBadge.IsStorySideBadge = false
	case enum.CommunicationMemberDetailBadgeTypeVoice:
		detailBadge.IsVoiceBadge = false
	case enum.CommunicationMemberDetailBadgeTypeTheme:
		detailBadge.IsThemeBadge = false
	case enum.CommunicationMemberDetailBadgeTypeCard:
		detailBadge.IsCardBadge = false
	case enum.CommunicationMemberDetailBadgeTypeMusic:
		detailBadge.IsMusicBadge = false
	default:
		panic("unknown type")
	}
	user_member.UpdateUserCommunicationMemberDetailBadge(session, detailBadge)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func UpdateUserLiveDifficultyNewFlag(ctx *gin.Context) {
	// mark all the song that this member is featured in as not new
	// only choose from the song user has access to, so no bond song and story locked songs
	// this use the same request as the above, must have been coded around the same time
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateMemberDetailBadgeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	liveDifficultyRecords := session.GetAllLiveDifficulties()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	for _, liveDifficultyRecord := range liveDifficultyRecords {
		if !liveDifficultyRecord.IsNew { // no need to update
			continue
		}
		// update if it feature this member
		liveDifficultyMaster := gamedata.LiveDifficulty[liveDifficultyRecord.LiveDifficultyId]
		if liveDifficultyMaster == nil {
			// some song no longer exists but official server still send them
			// it's ok to ignore these for now
			continue
		}
		_, exist := liveDifficultyMaster.Live.LiveMemberMapping[req.MemberMasterId]
		if exist {
			liveDifficultyRecord.IsNew = false
			session.UpdateLiveDifficulty(liveDifficultyRecord)
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func FinishUserStorySide(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStorySideRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	session.FinishStorySide(req.StorySideMasterId)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func FinishUserStoryMember(ctx *gin.Context) {
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
			user_content.AddContent(session, *storyMemberMaster.Reward)
			session.AddTriggerBasic(client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeStoryMemberReward,
				ParamInt:        generic.NewNullable(req.StoryMemberMasterId),
			})
		}
		if storyMemberMaster.UnlockLiveId != nil {
			masterLive := gamedata.Live[int32(*storyMemberMaster.UnlockLiveId)]
			// insert empty record for relevant items
			for _, masterLiveDifficulty := range masterLive.LiveDifficulties {
				userLiveDifficulty := session.GetUserLiveDifficulty(masterLiveDifficulty.LiveDifficultyId)
				session.UpdateLiveDifficulty(userLiveDifficulty)
			}
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SetTheme(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetThemeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	member := user_member.GetMember(session, req.MemberMasterId)
	member.SuitMasterId = req.SuitMasterId
	member.CustomBackgroundMasterId = req.CustomBackgroundMasterId
	user_member.UpdateMember(session, member)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SetFavoriteMember(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetFavoriteMemberRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.FavoriteMemberId = int32(req.MemberMasterId)
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseFavoriateMember {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseLovePointUp
		// award the initial SR
		// TODO(magic_id)
		intiialSr := int32(100002001 + req.MemberMasterId*10000)
		user_card.AddUserCardByCardMasterId(session, intiialSr)
		session.UserStatus.RecommendCardMasterId = intiialSr // this is for the pop up
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/communicationMember/fetchCommunicationMemberDetail", FetchCommunicationMemberDetail)
	router.AddHandler("/communicationMember/finishUserStoryMember", FinishUserStoryMember)
	router.AddHandler("/communicationMember/finishUserStorySide", FinishUserStorySide)
	router.AddHandler("/communicationMember/setFavoriteMember", SetFavoriteMember)
	router.AddHandler("/communicationMember/setTheme", SetTheme)
	router.AddHandler("/communicationMember/updateUserCommunicationMemberDetailBadge", UpdateUserCommunicationMemberDetailBadge)
	router.AddHandler("/communicationMember/updateUserLiveDifficultyNewFlag", UpdateUserLiveDifficultyNewFlag)
}
