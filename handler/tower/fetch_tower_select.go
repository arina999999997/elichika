package tower

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchTowerSelect(ctx *gin.Context) {
	// there's no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// no need to return anything, the client uses database for this
	// probably used to add DLP without having to add anything to database
	common.JsonResponse(ctx, &response.FetchTowerSelectResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tower/fetchTowerSelect", fetchTowerSelect)
}
