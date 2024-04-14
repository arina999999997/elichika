package take_over

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/locale"
	"elichika/router"
	"elichika/subsystem/user_authentication"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

Special behavior when user_id is -1 for when someone want to recover account without trashing the db with a random user id
*/
func checkTakeOver(ctx *gin.Context) {
	req := request.CheckTakeOverRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
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
	if linkedUserId < -1 {
		panic("user_id can't be negative except for the recovery ones")
	}

	currentSession = userdata.GetSession(ctx, req.UserId)
	defer currentSession.Close()
	linkedSession = userdata.GetSession(ctx, linkedUserId)
	defer linkedSession.Close()

	if currentSession != nil { // has current session, fill in current user
		resp.CurrentData.UserId = currentSession.UserId
		resp.CurrentData.LastLoginAt = currentSession.UserStatus.LastLoginAt
		resp.CurrentData.SnsCoin = currentSession.UserStatus.FreeSnsCoin +
			currentSession.UserStatus.AppleSnsCoin + currentSession.UserStatus.GoogleSnsCoin
	}
	if linkedSession != nil { // user exist
		if !user_authentication.CheckPassWord(linkedSession, req.PassWord) { // incorrect password
			resp.IsNotTakeOver = true
			goto FINISH_RESPONSE
		}
		resp.LinkedData.UserId = linkedSession.UserId
		// resp.LinkedData.AuthorizationKey = user_authentication.LoginSessionKey(nil, req.Mask)
		resp.LinkedData.AuthorizationKey = linkedSession.EncodedAuthorizationKey(req.Mask)
		resp.LinkedData.Name = linkedSession.UserStatus.Name
		resp.LinkedData.LastLoginAt = linkedSession.UserStatus.LastLoginAt
		resp.LinkedData.SnsCoin = linkedSession.UserStatus.FreeSnsCoin +
			linkedSession.UserStatus.AppleSnsCoin + linkedSession.UserStatus.GoogleSnsCoin
		resp.LinkedData.TermsOfUseVersion = linkedSession.UserStatus.TermsOfUseVersion

	} else { // user doesn't exist, but we won't create an account until setTakeOver is called
		resp.LinkedData.UserId = linkedUserId
		// resp.LinkedData.AuthorizationKey = user_authentication.LoginSessionKey(nil, req.Mask)
		// resp.LinkedData.AuthorizationKey = linkedSession.EncodedAuthorizationKey(req.Mask)
		// resp.LinkedData.AuthorizationKey = ""
		if linkedUserId == -1 {
			resp.LinkedData.Name.DotUnderText = "Recovery"
			resp.LinkedData.LastLoginAt = time.Now().Unix()
			resp.LinkedData.SnsCoin = 0
			resp.LinkedData.TermsOfUseVersion = 4
		} else {
			resp.LinkedData.Name.DotUnderText = "Newcomer"
			resp.LinkedData.LastLoginAt = time.Now().Unix()
			resp.LinkedData.SnsCoin = 100000
			resp.LinkedData.TermsOfUseVersion = 4
		}
	}

FINISH_RESPONSE:

	respBody, _ := json.Marshal(resp)
	signedResp := common.SignResp(ctx, string(respBody), ctx.MustGet("locale").(*locale.Locale).StartupKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, signedResp)
}

func init() {
	router.AddHandler("/", "POST", "/takeOver/checkTakeOver", checkTakeOver)
}
