package emblem

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchEmblem(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

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
		if emblem.IsNew {
			emblem.IsNew = false
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/emblem/fetchEmblem", fetchEmblem)
}
