package take_over

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/handler/login"
	"elichika/locale"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

/*
The take over system is used as a pseudo account system.
Use to switch account:
- Transfer Id should be the same as user Id (9 digits).
- The password is the login password.
Use to create new account:
If the user Id is new, then a new account will be created.
- The password entered will be the login password.
- User can user the datalink function to change the password as long as they have access to the account.
TODO(password): Password is stored in plaintext, maybe something like bcrypt would be better but the password is always sent using plain text anyway
*/
func checkTakeOver(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.CheckTakeOverRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	resp := response.CheckTakeOverResponse{}

	var currentSession, linkedSession (*userdata.Session)
	var linkedUserId int32
	_linkedUserId, err := strconv.Atoi(req.TakeOverId)
	if (err != nil) || (len(req.TakeOverId) > 9) {
		resp.IsNotTakeOver = true
		goto FINISH_RESPONSE
	}
	linkedUserId = int32(_linkedUserId)

	currentSession = userdata.GetSession(ctx, req.UserId)
	defer currentSession.Close()
	linkedSession = userdata.GetSession(ctx, linkedUserId)
	defer linkedSession.Close()

	if currentSession != nil { // has current session, fill in current user
		resp.CurrentData.UserId = int32(currentSession.UserId)
		resp.CurrentData.LastLoginAt = currentSession.UserStatus.LastLoginAt
		resp.CurrentData.SnsCoin = currentSession.UserStatus.FreeSnsCoin +
			currentSession.UserStatus.AppleSnsCoin + currentSession.UserStatus.GoogleSnsCoin
	}
	if linkedSession != nil { // user exist
		if !linkedSession.CheckPassWord(req.PassWord) { // incorrect password
			resp.IsNotTakeOver = true
			goto FINISH_RESPONSE
		}
		resp.LinkedData.UserId = int32(linkedSession.UserId)
		resp.LinkedData.AuthorizationKey = login.LoginSessionKey(req.Mask)
		resp.LinkedData.Name = linkedSession.UserStatus.Name
		resp.LinkedData.LastLoginAt = linkedSession.UserStatus.LastLoginAt
		resp.LinkedData.SnsCoin = linkedSession.UserStatus.FreeSnsCoin +
			linkedSession.UserStatus.AppleSnsCoin + linkedSession.UserStatus.GoogleSnsCoin
		resp.LinkedData.TermsOfUseVersion = linkedSession.UserStatus.TermsOfUseVersion

	} else { // user doesn't exist, but we won't create an account until setTakeOver is called
		resp.LinkedData.UserId = int32(linkedUserId)
		resp.LinkedData.AuthorizationKey = login.LoginSessionKey(req.Mask)
		resp.LinkedData.Name.DotUnderText = "Newcomer"
		resp.LinkedData.LastLoginAt = time.Now().Unix()
		resp.LinkedData.SnsCoin = 100000
		resp.LinkedData.TermsOfUseVersion = 4
	}

FINISH_RESPONSE:

	respBody, _ := json.Marshal(resp)
	signedResp := common.SignResp(ctx, string(respBody), ctx.MustGet("locale").(*locale.Locale).StartupKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, signedResp)
}

func init() {
	router.AddHandler("/takeOver/checkTakeOver", checkTakeOver)
}
