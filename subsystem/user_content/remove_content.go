package user_content

import (
	"elichika/client"
	"elichika/userdata"
)

func RemoveContent(session *userdata.Session, content client.Content) any {
	content.ContentAmount = -content.ContentAmount
	return AddContent(session, content)
}
