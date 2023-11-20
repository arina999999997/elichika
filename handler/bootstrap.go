package handler

import (
	"elichika/config"
	"elichika/userdata"

	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchBootstrap(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type BootstrapReq struct {
		BootstrapFetchTypes []int  `json:"bootstrap_fetch_types"`
		DeviceToken         string `json:"device_token"`
		DeviceName          string `json:"device_name"`
	}
	req := BootstrapReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	session.UserStatus.BootstrapSifidCheckAt = time.Now().UnixMilli()
	session.UserStatus.DeviceToken = req.DeviceToken
	signBody := session.Finalize(GetData("fetchBootstrap.json"), "user_model_diff")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetClearedPlatformAchievement(ctx *gin.Context) {
	signBody := GetData("getClearedPlatformAchievement.json")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
