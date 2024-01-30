package user_status

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func addSubscriptionCoin(session *userdata.Session, content *client.Content) any {
	user_content.OverflowCheckedAdd(&session.UserStatus.SubscriptionCoin, &content.ContentAmount)
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSubscriptionCoin, addSubscriptionCoin)
}
