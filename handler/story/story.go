package story

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FinishStoryMain(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.StoryMainRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	resp := response.StoryMainResponse{
		UserModelDiff: &session.UserModel,
	}

	if session.InsertUserStoryMain(req.CellId) { // newly inserted story, award some gem
		resp.FirstClearReward.Append(item.StarGem.Amount(10))
		session.AddContent(item.StarGem.Amount(10))
	}
	if req.MemberId.HasValue { // has a member -> select member thingy
		session.UpdateUserStoryMainSelected(req.CellId, req.MemberId.Value)
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func SaveBrowseStoryMainDigestMovie(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveBrowseStoryMainDigestMovieRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.InsertUserStoryMainPartDigestMovie(req.PartId)

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func FinishStoryLinkage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AddStoryLinkageRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	session.InsertUserStoryLinkage(req.CellId)

	session.Finalize()
	common.JsonResponse(ctx, &response.AddStoryLinkageResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	// /story/ all done
	router.AddHandler("/story/finishStoryLinkage", FinishStoryLinkage)
	router.AddHandler("/story/finishUserStoryMain", FinishStoryMain)
	router.AddHandler("/story/saveBrowseStoryMainDigestMovie", SaveBrowseStoryMainDigestMovie)
}
