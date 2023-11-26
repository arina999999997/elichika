package handler

import (
	"elichika/account"
	"elichika/config"
	"elichika/encrypt"
	"elichika/locale"
	"elichika/userdata"
	"elichika/utils"

	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func StartUpAuthorizationKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	if err != nil {
		panic(err)
	}
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	newKey := utils.Xor(randomBytes, []byte(config.SessionKey))
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	return newKey64
}

func LoginSessionKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	utils.CheckErr(err)
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	serverEventReceiverKey, err := hex.DecodeString(config.ServerEventReceiverKey)
	utils.CheckErr(err)
	jaKey, err := hex.DecodeString(config.JaKey)
	utils.CheckErr(err)
	newKey := utils.Xor(randomBytes, []byte(config.SessionKey))
	newKey = utils.Xor(newKey, serverEventReceiverKey)
	newKey = utils.Xor(newKey, jaKey)
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	return newKey64
}

func StartUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type StartUpReq struct {
		Mask                        string `json:"mask"`
		ResemaraDetectionIdentifier string `json:"resemara_detection_identifier"` // reset marathon (reroll)
		TimeDifference              int    `json:"time_difference"`               // second different from utc + 0
		RecaptchaToken              string `json:"recaptcha_token"`               // not necessary
	}
	req := StartUpReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	type StartUpResp struct {
		UserID           int    `json:"user_id"`
		AuthorizationKey string `json:"authorization_key"`
	}
	respObj := StartUpResp{}
	respObj.UserID = userdata.CreateNewAccount(ctx, -1, "")
	respObj.AuthorizationKey = StartUpAuthorizationKey(req.Mask)
	startupBody, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(startupBody), ctx.MustGet("locale").(*locale.Locale).StartUpKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func Login(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var mask64 string
	req := gjson.Parse(reqBody)
	req.ForEach(func(key, value gjson.Result) bool {
		if value.Get("mask").String() != "" {
			mask64 = value.Get("mask").String()
			return false
		}
		return true
	})
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	if session == nil {
		userdata.CreateNewAccount(ctx, userID, "")
		session = userdata.GetSession(ctx, userID)
		defer session.Close()
	}
	fmt.Println("User logins: ", userID)
	loginResponse := session.Login()
	loginResponse.SessionKey = LoginSessionKey(mask64)
	session.Finalize("{}", "user_model")
	loginBody, err := json.Marshal(loginResponse)
	utils.CheckErr(err)
	resp := SignResp(ctx, string(loginBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
	utils.WriteAllText(fmt.Sprint("login_", userID, ".json"), account.ExportUser(ctx))
}
