package user_present

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"
)

func Receive(session *userdata.Session, presentId int32, resp *response.ReceivePresentResponse) {
	present := client.PresentItem{}
	exist, err := session.Db.Table("u_present").Where("user_id = ? AND id = ?", session.UserId, presentId).
		Get(&present)
	// LimitExceededItems is for
	utils.CheckErrMustExist(err, exist)

	// has expired item panic so the user has to reload and trigger present delete
	// for now we will give the present to the user anyway if they managed to load it before it expire
	// if (present.ExpiredAt.HasValue && (present.ExpiredAt.Value <= session.Time.Unix())) {
	// 	panic("has expired item")
	// }
	received, result := user_content.AddContent(session, present.Content)
	if !received {
		resp.LimitExceededItems.Append(present)
	} else {
		// received, we need to add this item to the history and set relevant things
		// TODO(now): insert to deleted table
		_, err = session.Db.Table("u_present").Where("user_id = ? AND id = ?", session.UserId, presentId).
			Delete(&client.PresentItem{})
		utils.CheckErr(err)

		resp.ReceivedPresentItems.Append(present.Content)

		// if it's a card then we need to set this
		if present.Content.ContentType == enum.ContentTypeCard {
			resp.CardGradeUpResult.Append(result.(client.AddedCardResult))
		}
	}
}
