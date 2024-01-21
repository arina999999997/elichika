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
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchCommunicationMemberDetail(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchCommunicationMemberDetailRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchCommunicationMemberDetailResponse{}
	resp.MemberLovePanels.Append(session.GetMemberLovePanel(req.MemberId))

	year, month, day := session.Time.Date()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, session.Time.Location()).Unix()
	resp.WeekdayState.Weekday = int32(session.Time.Weekday())
	if resp.WeekdayState.Weekday == 0 {
		resp.WeekdayState.Weekday = 7
	}
	resp.WeekdayState.NextWeekdayAt = tomorrow
	common.JsonResponse(ctx, resp)
}

func UpdateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateMemberDetailBadgeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	detailBadge := session.GetUserCommunicationMemberDetailBadge(req.MemberMasterId)
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
	session.UpdateUserCommunicationMemberDetailBadge(detailBadge)

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

	userId := ctx.GetInt("user_id")
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

	userId := ctx.GetInt("user_id")
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

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	if session.FinishStoryMember(req.StoryMemberMasterId) {
		storyMemberMaster := gamedata.StoryMember[req.StoryMemberMasterId]
		if storyMemberMaster.Reward != nil {
			session.AddResource(*storyMemberMaster.Reward)
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

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	member := session.GetMember(req.MemberMasterId)
	member.SuitMasterId = req.SuitMasterId
	member.CustomBackgroundMasterId = req.CustomBackgroundMasterId
	session.UpdateMember(member)

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

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.FavoriteMemberId = int32(req.MemberMasterId)
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseFavoriateMember {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseLovePointUp
		// award the initial SR
		// TODO(refactor): use the common method to add a card instead
		card := session.GetUserCard(int32(100002001 + req.MemberMasterId*10000))
		session.UserStatus.RecommendCardMasterId = int32(card.CardMasterId) // this is for the pop up
		cardRarity := int32(20)
		member := session.GetMember(int32(req.MemberMasterId))
		beforeLoveLevelLimit := session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit)
		afterLoveLevelLimit := beforeLoveLevelLimit + cardRarity/10
		if afterLoveLevelLimit > session.Gamedata.MemberLoveLevelCount {
			afterLoveLevelLimit = session.Gamedata.MemberLoveLevelCount
		}
		member.LovePointLimit = session.Gamedata.MemberLoveLevelLovePoint[afterLoveLevelLimit]
		card.Grade++ // new grade,
		if card.Grade == 0 {
		} else {
			// add trigger card grade up so animation play when opening the card
			session.AddTriggerCardGradeUp(client.UserInfoTriggerCardGradeUp{
				CardMasterId:         card.CardMasterId,
				BeforeLoveLevelLimit: int32(afterLoveLevelLimit), // this is correct
				AfterLoveLevelLimit:  int32(afterLoveLevelLimit),
			})
		}
		// update the card and member
		session.UpdateUserCard(card)
		session.UpdateMember(member)
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
