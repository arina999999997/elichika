package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// TODO(refactor): Change to use request and response types
func FetchCommunicationMemberDetail(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	var memberId int32
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("member_id").String() != "" {
			memberId = int32(value.Get("member_id").Int())
			return false
		}
		return true
	})

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	lovePanelCellIds := session.GetLovePanelCellIds(memberId)

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	signBody := GetData("fetchCommunicationMemberDetail.json")
	signBody, _ = sjson.Set(signBody, "member_love_panels.0.member_id", memberId)
	signBody, _ = sjson.Set(signBody, "member_love_panels.0.member_love_panel_cell_ids", lovePanelCellIds)
	signBody, _ = sjson.Set(signBody, "weekday_state.weekday", weekday)
	signBody, _ = sjson.Set(signBody, "weekday_state.next_weekday_at", tomorrow)
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func UpdateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateUserCommunicationMemberDetailBadgeRequest{}
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

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func UpdateUserLiveDifficultyNewFlag(ctx *gin.Context) {
	// mark all the song that this member is featured in as not new
	// only choose from the song user has access to, so no bond song and story locked songs
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type UpdateUserLiveDifficultyNewFlag struct {
		MemberMasterId int `json:"member_master_id"`
	}
	req := UpdateUserLiveDifficultyNewFlag{}
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
		liveDifficultyMaster := gamedata.LiveDifficulty[int(liveDifficultyRecord.LiveDifficultyId)]
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

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func FinishUserStorySide(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStorySideRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.IsAutoMode = req.IsAutoMode
	session.FinishStorySide(req.StorySideMasterId)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func FinishUserStoryMember(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryMemberRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	session.UserStatus.IsAutoMode = req.IsAutoMode
	if session.FinishStoryMember(req.StoryMemberMasterId) {
		storyMemberMaster := gamedata.StoryMember[req.StoryMemberMasterId]
		if storyMemberMaster.Reward != nil {
			session.AddResource(*storyMemberMaster.Reward)
		}
		if storyMemberMaster.UnlockLiveId != nil {
			masterLive := gamedata.Live[*storyMemberMaster.UnlockLiveId]
			// insert empty record for relevant items
			for _, masterLiveDifficulty := range masterLive.LiveDifficulties {
				liveDifficulty := session.GetLiveDifficulty(int32(masterLiveDifficulty.LiveDifficultyId))
				session.UpdateLiveDifficulty(liveDifficulty)
			}
		}
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func SetTheme(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetThemeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	member := session.GetMember(int32(req.MemberMasterId))
	member.SuitMasterId = int32(req.SuitMasterId)
	member.CustomBackgroundMasterId = int32(req.CustomBackgroundMasterId)
	session.UpdateMember(member)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
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
			member.OwnedCardCount++
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

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
