package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// TODO(refactor): Change to use request and response types
func FinishStoryMain(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryMainRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.IsAutoMode = req.IsAutoMode
	// if session.UserStatus.TutorialPhase == enum.TutorialPhaseStory1 {
	// 	session.UserStatus.TutorialPhase = enum.TutorialPhaseStory2
	// } else if session.UserStatus.TutorialPhase == enum.TutorialPhaseStory3 {
	// 	session.UserStatus.TutorialPhase = enum.TutorialPhaseStory4
	// }
	firstClearReward := []client.Content{}
	if session.InsertUserStoryMain(req.CellId) { // newly inserted story, award some gem
		firstClearReward = append(firstClearReward, client.Content{
			ContentType:   enum.ContentTypeSnsCoin,
			ContentId:     0,
			ContentAmount: 10,
		})
		session.AddResource(firstClearReward[0])
	}
	if req.MemberId != nil { // has a member -> select member thingy
		session.UpdateUserStoryMainSelected(req.CellId, *req.MemberId)
	}

	signBody := session.Finalize("{}", "user_model_diff")
	signBody, _ = sjson.Set(signBody, "first_clear_reward", firstClearReward)
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func SaveBrowseStoryMainDigestMovie(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveBrowseStoryMainDigestMovieRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.InsertUserStoryMainPartDigestMovie(req.PartId)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func FinishStoryLinkage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryLinkageRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UserStatus.IsAutoMode = req.IsAutoMode
	session.InsertUserStoryLinkage(req.CellId)
	signBody := session.Finalize("{}", "user_model_diff")
	// technically correct because it's way past the end of the reward periods
	// but it would be cool to implement some reward handling
	signBody, _ = sjson.Set(signBody, "has_additional_rewards", false)
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
