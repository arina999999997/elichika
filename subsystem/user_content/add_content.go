// directly add content to the account, return bool, set the relevant user model fields,
// and a return interface if it's actually added
// if it's not added, add the items into a unreceived list
// finally there will be a finalizer that add the unreceived items to present box
// when receving from present box, we can clear the list so there would be no doubly added presents

package user_content

import (
	"elichika/client"
	"elichika/userdata"

	"fmt"
)

func AddContent(session *userdata.Session, content client.Content) any {
	if content.ContentAmount == 0 { // caller should gracefully accept this
		return nil
	}
	handler, exist := contentHandlerByContentType[content.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", content.ContentType)
		return nil
	}
	result := handler(session, &content)
	if content.ContentAmount > 0 { // if not fully received then we need to track it
		session.UnreceivedContent = append(session.UnreceivedContent, content)
	}
	return result
}
