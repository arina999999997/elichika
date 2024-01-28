package user_content

import (
	"elichika/client"
	"elichika/userdata"
)

func RemoveContent(session *userdata.Session, content client.Content) (bool, any) {
	content.ContentAmount = -content.ContentAmount
	added, result := AddContent(session, content)
	return added, result
}
