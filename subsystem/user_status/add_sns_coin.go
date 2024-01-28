package user_status

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

// TODO(sns_coin): Handle paid(?)
func addSnsCoin(session *userdata.Session, content client.Content) (bool, any) {
	return user_content.OverflowCheckedAdd(&session.UserStatus.FreeSnsCoin, content.ContentAmount), nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSnsCoin, addSnsCoin)
}
