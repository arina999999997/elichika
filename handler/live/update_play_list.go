package live

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_play_list"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func updatePlayList(ctx *gin.Context) {
	req := request.UpdatePlayListRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_play_list.UpdateUserPlayList(session, req.GroupNum, req.LiveMasterId, req.IsSet)

	common.JsonResponse(ctx, &response.UpdatePlayListResponse{
		IsSuccess:     true,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/live/updatePlayList", updatePlayList)
}
