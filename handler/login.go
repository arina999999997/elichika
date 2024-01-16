package handler

import (
	"elichika/client/request"
	"elichika/client/response"
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

func StartupAuthorizationKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	utils.CheckErr(err)
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

func Startup(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.StartupRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	resp := response.StartupResponse{}
	resp.UserId = int32(userdata.CreateNewAccount(ctx, -1, ""))
	resp.AuthorizationKey = StartupAuthorizationKey(req.Mask)
	// note that this use a different key than the common one
	startupBody, _ := json.Marshal(resp)
	respBody := SignResp(ctx, string(startupBody), ctx.MustGet("locale").(*locale.Locale).StartupKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, respBody)
}

// TODO(refactor): Change to use request and response types
func Login(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LoginRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session == nil {
		userdata.CreateNewAccount(ctx, userId, "")
		session = userdata.GetSession(ctx, userId)
		defer session.Close()
	}

	fmt.Println("User logins: ", userId)

	resp := session.Login()
	resp.SessionKey = LoginSessionKey(req.Mask)
	session.Finalize("{}", "dummy")
	JsonResponse(ctx, resp)

	{
		backupText, err := json.Marshal(resp)
		utils.CheckErr(err)
		utils.WriteAllText(fmt.Sprint(config.UserDataBackupPath, "login_", userId, ".json"), string(backupText))
	}
}
