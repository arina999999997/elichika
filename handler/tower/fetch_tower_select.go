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
	session := ctx.MustGet("session").(*userdata.Session)

	// no need to return anything, the client uses database for this
	// probably used to add DLP without having to add anything to database
	common.JsonResponse(ctx, &response.FetchTowerSelectResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/tower/fetchTowerSelect", fetchTowerSelect)
}
