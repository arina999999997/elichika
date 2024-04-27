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
)

func fetchEmblemById(ctx *gin.Context) {
	req := request.EmblemSearchUserIdRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	otherUserSession := userdata.GetSessionWithSharedDb(ctx, req.UserId, session)
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
	router.AddHandler("/", "POST", "/emblem/fetchEmblemById", fetchEmblemById)
}
