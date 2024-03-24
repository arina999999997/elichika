package userdata

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/userdata/database"
	"elichika/utils"

	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func randomKey() []byte {
	if *config.Conf.UseAuthenticationKeys {
		// random 32 bytes
		b := make([]byte, 32)
		_, err := rand.Read(b)
		utils.CheckErr(err)
		return b
	} else {
		return []byte(config.DefaultSessionKey)
	}
}

// load the auth from database, or generate new one
func (session *Session) fetchAuthenticationData() {
	exist, err := session.Db.Table("u_authentication").Where("user_id = ?", session.UserId).Get(&session.AuthenticationData)
	utils.CheckErr(err)
	if !exist {
		session.AuthenticationData = database.UserAuthentication{
			AuthorizationKey: randomKey(),
			SessionKey:       randomKey(),
		}
	}
}

func userAuthenticationDataFinalizer(session *Session) {
	affected, err := session.Db.Table("u_authentication").Where("user_id = ?", session.UserId).AllCols().Update(&session.AuthenticationData)
	utils.CheckErr(err)
	if affected == 0 {
		GenericDatabaseInsert(session, "u_authentication", session.AuthenticationData)
	}
}

func init() {
	AddFinalizer(userAuthenticationDataFinalizer)
}

func (session *Session) GenerateNewSessionKey() {
	session.AuthenticationData.SessionKey = randomKey()
}

func (session *Session) GenerateNewAuthorizationKey() {
	session.AuthenticationData.AuthorizationKey = randomKey()
}

func (session *Session) AuthorizationKey() []byte {
	if *config.Conf.UseAuthenticationKeys {
		return session.AuthenticationData.AuthorizationKey
	} else {
		return []byte(config.DefaultSessionKey)
	}
}
func (session *Session) SessionKey() []byte {
	if *config.Conf.UseAuthenticationKeys {
		return session.AuthenticationData.SessionKey
	} else {
		return []byte(config.DefaultSessionKey)
	}
}

func (session *Session) EncodedAuthorizationKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	utils.CheckErr(err)
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	newKey := utils.Xor(randomBytes, []byte(session.AuthorizationKey()))
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	return newKey64
}

func (session *Session) EncodedSessionKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	utils.CheckErr(err)
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	serverEventReceiverKey, err := hex.DecodeString(config.ServerEventReceiverKey)
	utils.CheckErr(err)
	newKey := utils.Xor(randomBytes, []byte(session.SessionKey()))
	newKey = utils.Xor(newKey, serverEventReceiverKey)
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	return newKey64
}
