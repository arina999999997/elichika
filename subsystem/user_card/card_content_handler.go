package user_card

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func cardContentHandler(session *userdata.Session, content *client.Content) any {
	if content.ContentAmount > 6 {
		// there's no instance of adding more than 6 of a card in offcial server
		// this is to prevent webui usage from lagging/crashing the server if the amount is too much
		content.ContentAmount = 6
	}
	results := []client.AddedCardResult{}
	for content.ContentAmount > 0 {
		content.ContentAmount--
		results = append(results, AddUserCardByCardMasterId(session, content.ContentId))
	}
	return results
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeCard, cardContentHandler)
}
