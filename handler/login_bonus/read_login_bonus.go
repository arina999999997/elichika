package login_bonus

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func readLoginBonus(ctx *gin.Context) {
	// this doesn't need to do anything, at least with this way of handling things
	// reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// req := request.ReadLoginBonusRequest{}
	// err := json.Unmarshal([]byte(reqBody), &req)
	// utils.CheckErr(err)
	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func init() {
	router.AddHandler("/loginBonus/readLoginBonus", readLoginBonus)
}
