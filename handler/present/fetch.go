package present

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(now): Implement presents
func fetch(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchPresentResponse{
		PresentItems: user_present.FetchPresentItems(session),
	}
	resp.PresentCount = int32(resp.PresentItems.Size())

	session.Finalize() // this is because the fetch request can cause server to delete expired item
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/present/fetch", fetch)
}
