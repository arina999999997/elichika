package live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func LiveUpdatePlayList(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdatePlayListRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if req.IsSet {
		session.AddUserPlayList(client.UserPlayList{
			UserPlayListId: req.GroupNum + req.LiveMasterId*10,
			GroupNum:       req.GroupNum,
			LiveId:         req.LiveMasterId,
		})
	} else {
		session.DeleteUserPlayList(req.GroupNum + req.LiveMasterId*10)
	}

	session.Finalize()
	common.JsonResponse(ctx, &response.UpdatePlayListResponse{
		IsSuccess:     true,
		UserModelDiff: &session.UserModel,
	})
}
