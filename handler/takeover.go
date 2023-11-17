package handler

import (
	"elichika/config"
	"elichika/utils"
	// "elichika/encrypt"
	"elichika/locale"
	"elichika/userdata"

	"fmt"
	"net/http"
	"strconv"
	"time"
	// "encoding/base64"
	// "encoding/hex"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

/*
The take over system is used as a pseudo account system.
Use to switch account:
- Transfer ID should be the same as user ID (9 digits).
- The password is the login password.
Use to create new account:
If the user ID is new, then a new account will be created.
- The password entered will be the login password.
- User can user the datalink function to change the password as long as they have access to the account.
Password is stored in plaintext, if you want things like bcrypt, do it yourself.
*/

type TakeOverData struct {
	UserID           int    `json:"user_id"`
	AuthorizationKey string `json:"authorization_key"`
	Name             struct {
		DotUnderText string `json:"dot_under_text"`
	} `json:"name"`
	// Unix second
	LastLoginAt          int64   `json:"last_login_at"`
	SnsCoin              int     `json:"sns_coin"`
	TermsOfUseVersion    int     `json:"terms_of_use_version"`
	ServiceUserCommonKey *string `json:"service_user_common_key"`
}

func CheckTakeOver(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type CheckTakeOverReq struct {
		CurrentUserID int    `json:"user_id"`
		TakeOverID    string `json:"take_over_id"`
		PassWord      string `json:"pass_word"`
		Mask          string `json:"mask"`
	}
	req := CheckTakeOverReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	type TakeOverResp struct {
		IsNotTakeOver bool          `json:"is_not_take_over"`
		LinkedData    TakeOverData  `json:"linked_data"`
		CurrentData   *TakeOverData `json:"current_data"`
	}
	resp := TakeOverResp{}
	resp.IsNotTakeOver = false
	var currentSession, linkedSession (*userdata.Session)
	linkedUserID, err := strconv.Atoi(req.TakeOverID)
	if (err != nil) || (len(req.TakeOverID) > 9) {
		resp.IsNotTakeOver = true
		goto FINISH_RESPONSE
	}

	currentSession = userdata.GetSession(ctx, req.CurrentUserID)
	defer currentSession.Close()
	linkedSession = userdata.GetSession(ctx, linkedUserID)
	defer linkedSession.Close()

	if currentSession != nil { // has current session, fill in current user
		resp.CurrentData = new(TakeOverData)
		resp.CurrentData.UserID = currentSession.UserStatus.UserID
		resp.CurrentData.LastLoginAt = currentSession.UserStatus.LastLoginAt
		resp.CurrentData.SnsCoin = currentSession.UserStatus.FreeSnsCoin +
			currentSession.UserStatus.AppleSnsCoin + currentSession.UserStatus.GoogleSnsCoin
	}
	if linkedSession != nil { // user exists
		if linkedSession.UserStatus.PassWord != req.PassWord { // incorrect password
			resp.IsNotTakeOver = true
			goto FINISH_RESPONSE
		}
		resp.LinkedData.UserID = linkedSession.UserStatus.UserID
		resp.LinkedData.AuthorizationKey = LoginSessionKey(req.Mask)
		resp.LinkedData.ServiceUserCommonKey = nil
		resp.LinkedData.Name.DotUnderText = linkedSession.UserStatus.Name.DotUnderText
		resp.LinkedData.LastLoginAt = linkedSession.UserStatus.LastLoginAt
		resp.LinkedData.SnsCoin = linkedSession.UserStatus.FreeSnsCoin +
			linkedSession.UserStatus.AppleSnsCoin + linkedSession.UserStatus.GoogleSnsCoin
		resp.LinkedData.TermsOfUseVersion = linkedSession.UserStatus.TermsOfUseVersion

	} else { // user doesn't exists, but we won't create an account until setTakeOver is called
		resp.LinkedData.UserID = linkedUserID
		resp.LinkedData.AuthorizationKey = LoginSessionKey(req.Mask)
		resp.LinkedData.ServiceUserCommonKey = nil
		resp.LinkedData.Name.DotUnderText = "Newcomer"
		resp.LinkedData.LastLoginAt = time.Now().Unix()
		resp.LinkedData.SnsCoin = 100000
		resp.LinkedData.TermsOfUseVersion = 4
	}

FINISH_RESPONSE:
	respBody, _ := json.Marshal(resp)
	signedResp := SignResp(ctx, string(respBody), ctx.MustGet("locale").(*locale.Locale).StartUpKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, signedResp)
}

func SetTakeOver(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SetTakeOverReq struct {
		TakeOverID string `json:"take_over_id"`
		PassWord   string `json:"pass_word"`
		Mask       string `json:"mask"`
	}
	req := SetTakeOverReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	linkedUserID, err := strconv.Atoi(req.TakeOverID)
	utils.CheckErr(err)
	linkedSession := userdata.GetSession(ctx, linkedUserID)
	defer linkedSession.Close()
	if linkedSession == nil { // new account
		userdata.CreateNewAccount(ctx, linkedUserID, req.PassWord)
		linkedSession = userdata.GetSession(ctx, linkedUserID)
		defer linkedSession.Close()
	} else { // existing account, have to check password
		if linkedSession.UserStatus.PassWord != req.PassWord {
			panic("wrong password")
		}
	}
	resp := TakeOverData{}
	resp.UserID = linkedSession.UserStatus.UserID
	resp.AuthorizationKey = StartUpAuthorizationKey(req.Mask)
	resp.ServiceUserCommonKey = nil
	resp.Name.DotUnderText = linkedSession.UserStatus.Name.DotUnderText
	resp.LastLoginAt = linkedSession.UserStatus.LastLoginAt
	resp.SnsCoin = linkedSession.UserStatus.FreeSnsCoin +
		linkedSession.UserStatus.AppleSnsCoin + linkedSession.UserStatus.GoogleSnsCoin
	resp.TermsOfUseVersion = linkedSession.UserStatus.TermsOfUseVersion

	signedResp, _ := sjson.Set("{}", "data", resp)

	signedResp = SignResp(ctx, signedResp, ctx.MustGet("locale").(*locale.Locale).StartUpKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, signedResp)

}

func UpdatePassWord(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type UpdatePassWordReq struct {
		PassWord string `json:"pass_word"`
	}
	req := UpdatePassWordReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	type UpdatePassWordResp struct {
		TakeOverID string `json:"take_over_id"`
	}
	session.UserStatus.PassWord = req.PassWord
	session.Finalize("", "")
	respObj := UpdatePassWordResp{}
	respObj.TakeOverID = fmt.Sprintf("%09d", userID)
	startupBody, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(startupBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
