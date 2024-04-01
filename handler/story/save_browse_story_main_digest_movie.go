package story

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_story_main"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func saveBrowseStoryMainDigestMovie(ctx *gin.Context) {
	req := request.SaveBrowseStoryMainDigestMovieRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	user_story_main.InsertUserStoryMainPartDigestMovie(session, req.PartId)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/story/saveBrowseStoryMainDigestMovie", saveBrowseStoryMainDigestMovie)
}
