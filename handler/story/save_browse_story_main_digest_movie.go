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

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	user_story_main.InsertUserStoryMainPartDigestMovie(session, req.PartId)

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/story/saveBrowseStoryMainDigestMovie", saveBrowseStoryMainDigestMovie)
}
