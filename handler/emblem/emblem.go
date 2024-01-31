package emblem

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchEmblem(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// this is official server behavior:
	// populate UserEmblemByEmblemId like login, then mirror that data for EmblemIsNewDataList
	// this is probably some spandrel in the way that the server was developed
	session.PopulateUserModelField("UserEmblemByEmblemId")
	resp := response.FetchEmblemResponse{
		UserModel: &session.UserModel,
	}
	for _, emblem := range session.UserModel.UserEmblemByEmblemId.Map {
		resp.EmblemIsNewDataList.Append(client.EmblemIsNewData{
			EmblemMasterId: emblem.EmblemMId,
			IsNew:          emblem.IsNew,
		})
	}

	common.JsonResponse(ctx, &resp)
}

func ActivateEmblem(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ActivateEmblemRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.EmblemId = req.EmblemMasterId

	session.Finalize()
	common.JsonResponse(ctx, response.ActivateEmblemResponse{
		UserModel: &session.UserModel,
	})
}

func FetchEmblemById(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.EmblemSearchUserIdRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	otherUserSession := userdata.GetSession(ctx, req.UserId)
	otherUserSession.PopulateUserModelField("UserEmblemByEmblemId")
	resp := response.FetchEmblemResponse{
		UserModel: &session.UserModel,
	}
	for _, emblem := range otherUserSession.UserModel.UserEmblemByEmblemId.Map {
		resp.EmblemIsNewDataList.Append(client.EmblemIsNewData{
			EmblemMasterId: emblem.EmblemMId,
			IsNew:          emblem.IsNew,
		})
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	// TODO(refactor): move to individual files. 
	router.AddHandler("/emblem/activateEmblem", ActivateEmblem)
	router.AddHandler("/emblem/fetchEmblem", FetchEmblem)
	router.AddHandler("/emblem/fetchEmblemById", FetchEmblemById)
}
