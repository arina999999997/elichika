package live_mv

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func start(ctx *gin.Context) {
	// we don't really need the request
	// maybe it's once needed or it's only used for gathering data
	// reqBody := ctx.Get("reqBody").(json.RawMessage)
	// req := request.StartLiveMvRequest{}
	// err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	// utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, &response.StartLiveMvResponse{
		UniqId:        session.Time.UnixNano(),
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/liveMv/start", start)
}
