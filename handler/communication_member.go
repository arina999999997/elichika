package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchCommunicationMemberDetail(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	var memberId int
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("member_id").String() != "" {
			memberId = int(value.Get("member_id").Int())
			return false
		}
		return true
	})

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	lovePanelCellIds := session.GetLovePanelCellIDs(memberId)

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

// where is this called?
// called when checked a thing that is marked as new?
func UpdateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateUserCommunicationMemberDetailBadgeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	fmt.Println(req)
	detailBadge := session.GetUserCommunicationMemberDetailBadge(req.MemberMasterID)
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

func UpdateUserLiveDifficultyNewFlag(ctx *gin.Context) {
	// mark all the song that this member is featured in as not new
	// only choose from the song user has access to, so no bond song and story locked songs
	// TODO: also need to mark some flag to get rid of the ! on the button
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type UpdateUserLiveDifficultyNewFlag struct {
		MemberMasterID int `json:"member_master_id"`
	}
	req := UpdateUserLiveDifficultyNewFlag{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")

	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	liveDifficultyRecords := session.GetAllLiveDifficulties()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	for _, liveDifficultyRecord := range liveDifficultyRecords {
		if !liveDifficultyRecord.IsNew { // no need to update
			continue
		}
		// update if it feature this member
		liveDifficultyMaster := gamedata.LiveDifficulty[liveDifficultyRecord.LiveDifficultyID]
		if liveDifficultyMaster == nil {
			// some song no longer exists but official server still send them
			// it's ok to ignore these for now
			continue
		}
		_, exist := liveDifficultyMaster.Live.LiveMemberMapping[req.MemberMasterID]
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

func FinishUserStorySide(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStorySideRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	session.UserStatus.IsAutoMode = req.IsAutoMode
	session.FinishStorySide(req.StorySideMasterID)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FinishUserStoryMember(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryMemberRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	session.UserStatus.IsAutoMode = req.IsAutoMode
	if session.FinishStoryMember(req.StoryMemberMasterID) {
		storyMemberMaster := gamedata.StoryMember[req.StoryMemberMasterID]
		if storyMemberMaster.Reward != nil {
			session.AddResource(*storyMemberMaster.Reward)
		}
		if storyMemberMaster.UnlockLiveID != nil {
			masterLive := gamedata.Live[*storyMemberMaster.UnlockLiveID]
			// insert empty record for relevant items
			for _, masterLiveDifficulty := range masterLive.LiveDifficulties {
				liveDifficulty := session.GetLiveDifficulty(masterLiveDifficulty.LiveDifficultyID)
				session.UpdateLiveDifficulty(liveDifficulty)
			}
		}
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetTheme(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetThemeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	member := session.GetMember(req.MemberMasterID)
	member.SuitMasterID = req.SuitMasterID
	member.CustomBackgroundMasterID = req.CustomBackgroundMasterID
	session.UpdateMember(member)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetFavoriteMember(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetFavoriteMemberRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	session.UserStatus.FavoriteMemberID = req.MemberMasterID

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
