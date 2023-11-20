package handler

import (
	"elichika/config"
	"elichika/userdata"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

func FetchMission(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := session.Finalize("{}", "user_model")
	signBody, _ = sjson.Set(signBody, "mission_master_id_list", []any{})
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ClearMissionBadge(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
