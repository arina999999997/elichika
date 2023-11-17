package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/model"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FinishStoryMain(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryMainRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	session.UserStatus.IsAutoMode = req.IsAutoMode
	firstClearReward := []model.Content{}
	if session.InsertUserStoryMain(req.CellID) { // newly inserted story, award some gem
		firstClearReward = append(firstClearReward, model.Content{
			ContentType:   enum.ContentTypeSnsCoin,
			ContentID:     0,
			ContentAmount: 10,
		})
		session.AddResource(firstClearReward[0])
	}
	if req.MemberID != nil { // has a member -> select member thingy
		session.UpdateUserStoryMainSelected(req.CellID, *req.MemberID)
	}

	signBody := session.Finalize("{}", "user_model_diff")
	signBody, _ = sjson.Set(signBody, "first_clear_reward", firstClearReward)
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
func SaveBrowseStoryMainDigestMovie(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveBrowseStoryMainDigestMovieRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	session.InsertUserStoryMainPartDigestMovie(req.PartID)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FinishStoryLinkage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishUserStoryLinkageRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	session.UserStatus.IsAutoMode = req.IsAutoMode
	session.InsertUserStoryLinkage(req.CellID)
	signBody := session.Finalize("{}", "user_model_diff")
	// technically correct because it's way past the end of the reward periods
	// but it would be cool to implement some reward handling
	signBody, _ = sjson.Set(signBody, "has_additional_rewards", false)
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
