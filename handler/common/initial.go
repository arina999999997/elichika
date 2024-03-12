package common

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/locale"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func initial(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	utils.CheckErr(err)

	defer ctx.Request.Body.Close()

	ctx.Set("reqBody", string(body))

	lang, _ := ctx.GetQuery("l")
	if lang == "" {
		lang = "ja"
	}
	ctx.Set("locale", locale.Locales[lang])
	ctx.Set("gamedata", locale.Locales[lang].Gamedata)
	ctx.Set("dictionary", locale.Locales[lang].Dictionary)
	ctx.Set("sign_key", config.DefaultSessionKey)

	userId, userIdErr := strconv.Atoi(ctx.Query("u"))
	ctx.Set("user_id", userId)

	ctx.Set("ep", ctx.Request.URL.String())

	messages := []json.RawMessage{}
	err = json.Unmarshal(body, &messages)
	utils.CheckErr(err)
	n := len(messages)
	sign := ""
	err = json.Unmarshal(messages[n-1], &sign)
	utils.CheckErr(err)
	if userIdErr == nil {
		session := userdata.GetSession(ctx, int32(userId))
		defer session.Close()
		if session == nil {
			panic("session is nil, use a transfer to get a proper user id")
		}
		// TODO(refactor): Reuse this session for the request instead of ending it
		// ctx.Set("session", session)

		ctx.Set("sign_key", session.SessionKey())
		// signAuth := encrypt.HMAC_SHA1_Encrypt([]byte(ctx.Request.URL.String()+" "+string(messages[n-2])), session.AuthorizationKey())
		// signSession := encrypt.HMAC_SHA1_Encrypt([]byte(ctx.Request.URL.String()+" "+string(messages[n-2])), session.SessionKey())
		// fmt.Println("auth: ", signAuth, "\nactual: ", string(messages[n-1]))
		// fmt.Println("session: ", signSession, "\nactual: ", string(messages[n-1]))
		commandId, _ := strconv.Atoi(ctx.Query("id"))
		if strings.HasPrefix(ctx.Request.URL.String(), "/login/login?") {
			signAuth := encrypt.HMAC_SHA1_Encrypt([]byte(ctx.Request.URL.String()+" "+string(messages[n-2])),
				session.AuthorizationKey())
			if sign != signAuth { // incorrect auth key, reject
				panic("wrong authentication key, sign-in again using password")
			}
			session.AuthenticationData.CommandId = int32(commandId)
		} else {
			signSession := encrypt.HMAC_SHA1_Encrypt([]byte(ctx.Request.URL.String()+" "+string(messages[n-2])),
				session.SessionKey())
			if sign != signSession { // incorrect auth key, reject
				panic("wrong session key")
			}
			if session.AuthenticationData.CommandId+1 != int32(commandId) {
				panic("wrong command id")
			} else {
				session.AuthenticationData.CommandId++
			}
		}
		session.Finalize()
	} else { // no user id, use the start up key
		signStartUp := encrypt.HMAC_SHA1_Encrypt([]byte(ctx.Request.URL.String()+" "+string(messages[n-2])),
			locale.Locales[lang].StartupKey)
		if sign != signStartUp { // incorrect start up key, reject
			fmt.Println("startup: ", signStartUp, "\nactual: ", sign)
			panic("wrong startup key, get the correct app version")
		}
	}

	ctx.Next()
}

func init() {
	router.AddInitialHandler(initial)
}
