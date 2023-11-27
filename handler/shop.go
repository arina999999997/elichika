package handler

import (
	"elichika/config"
	"elichika/userdata"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Not implementing this system seriously

func FetchShopTop(ctx *gin.Context) {
	resp := SignResp(ctx, GetData("fetchShopTop.json"), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchShopPack(ctx *gin.Context) {
	resp := SignResp(ctx, GetData("fetchShopPack.json"), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchShopSnsCoin(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	session.UserModel.UserStatus.FreeSnsCoin += 10000 // add 10000 gems everytime someone try to buy gem
	session.Finalize("{}", "dummy")

	resp := SignResp(ctx, GetData("fetchShopSnsCoin.json"), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
