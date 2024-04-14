package present

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// there is no request body

	session := ctx.MustGet("session").(*userdata.Session)
	// TODO(database): Have a common function to sync present state maybe
	resp := response.FetchPresentResponse{
		PresentItems:        user_present.FetchPresentItems(session),
		PresentHistoryItems: user_present.FetchPresentHistoryItems(session),
	}

	session.Finalize() // this is because the fetch request can cause server to delete expired item
	resp.PresentCount = user_present.FetchPresentCount(session)
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/present/fetch", fetch)
}
