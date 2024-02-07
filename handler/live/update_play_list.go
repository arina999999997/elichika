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
	"github.com/tidwall/gjson"
)

func updatePlayList(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdatePlayListRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_play_list.UpdateUserPlayList(session, req.GroupNum, req.LiveMasterId, req.IsSet)

	session.Finalize()
	common.JsonResponse(ctx, &response.UpdatePlayListResponse{
		IsSuccess:     true,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/live/updatePlayList", updatePlayList)
}
