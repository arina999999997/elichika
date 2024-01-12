package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(refactor): Change to use request and response types
func SaveUnlockedScene(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveUnlockedSceneRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	for _, sceneType := range req.UnlockSceneTypes {
		session.UnlockScene(sceneType, enum.UnlockSceneStatusOpened)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func SaveSceneTipsType(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveSceneTipsTypeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	// response with user model
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.SaveSceneTips(req.SceneTipsType)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
