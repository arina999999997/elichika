package user_content

import (
	"elichika/client"
	"elichika/userdata"
)

func genericContentHandler(session *userdata.Session, addedContent client.Content) (bool, any) {
	currentContent := GetUserContent(session, addedContent.ContentType, addedContent.ContentId)
	if OverflowCheckedAdd(&currentContent.ContentAmount, addedContent.ContentAmount) {
		UpdateUserContent(session, currentContent)
		return true, nil
	} else {
		return false, nil
	}
}
