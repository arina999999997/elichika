package handler

import (
	"elichika/config"
	"elichika/gacha"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchGachaMenu(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gachaList := session.GetGachaList()
	signBody := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "gacha_list", gachaList)
	signBody, _ = sjson.Set(signBody, "gacha_unconfirmed", nil)
	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GachaDraw(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := model.GachaDrawReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	ctx.Set("session", session)
	gacha, resultCards := gacha.HandleGacha(ctx, req)
	// TODO: add items
	signBody := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "gacha", gacha)
	signBody, _ = sjson.Set(signBody, "result_cards", resultCards)
	signBody, _ = sjson.Set(signBody, "result_bonuses", nil)
	signBody, _ = sjson.Set(signBody, "retry_gacha", nil)
	signBody, _ = sjson.Set(signBody, "stepup_next_step", nil)

	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
