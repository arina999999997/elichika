package admin

import (
	"elichika/router"
	"elichika/utils"

	"crypto/rand"
)

var adminSessionKey []byte

func randomKey() []byte {
	// random 32 bytes
	b := make([]byte, 32)
	_, err := rand.Read(b)
	utils.CheckErr(err)
	return b
}

func newSessionKey() {
	adminSessionKey = randomKey()
}

func init() {
	newSessionKey()
	router.AddTemplates("./webui/admin/logged_in_admin.html")
}
