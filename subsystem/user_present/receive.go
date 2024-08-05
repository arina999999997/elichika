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
	exist, err := session.Db.Table("u_present_item").Where("user_id = ? AND id = ?", session.UserId, presentId).
		Get(&present)
	// LimitExceededItems is for
	utils.CheckErrMustExist(err, exist)

	// has expired item panic so the user has to reload and trigger present delete
	// for now we will give the present to the user anyway if they managed to load it before it expire
	// if (present.ExpiredAt.HasValue && (present.ExpiredAt.Value <= session.Time.Unix())) {
	// 	panic("has expired item")
	// }
	result := user_content.AddContent(session, present.Content)
	if len(session.UnreceivedContent) == 0 {
		// completely received, we just need to delete this thing
		_, err := session.Db.Table("u_present_item").Where("user_id = ? AND id = ?", session.UserId, presentId).
			Delete(&client.PresentItem{})
		utils.CheckErr(err)

		userdata.GenericDatabaseInsert(session, "u_present_history_item", client.PresentHistoryItem{
			Content:          present.Content,
			PresentRouteType: present.PresentRouteType,
			PresentRouteId:   present.PresentRouteId,
			ParamServer:      present.ParamServer,
			ParamClient:      present.ParamClient,
			HistoryCreatedAt: session.Time.Unix(),
		})
		resp.ReceivedPresentItems.Append(present.Content)
		// if it's a card then we need to set this
		if present.Content.ContentType == enum.ContentTypeCard {
			for _, resultCard := range result.([]client.AddedCardResult) {
				resp.CardGradeUpResult.Append(resultCard)
			}
		}
	} else {
		// partially received
		remaining := session.UnreceivedContent[0]
		session.UnreceivedContent = []client.Content{}
		present.Content.ContentAmount -= remaining.ContentAmount
		if present.Content.ContentAmount > 0 { // did receive something
			// we need to update the history database
			userdata.GenericDatabaseInsert(session, "u_present_history_item", client.PresentHistoryItem{
				Content:          present.Content,
				PresentRouteType: present.PresentRouteType,
				PresentRouteId:   present.PresentRouteId,
				ParamServer:      present.ParamServer,
				ParamClient:      present.ParamClient,
				HistoryCreatedAt: session.Time.Unix(),
			})
			resp.ReceivedPresentItems.Append(present.Content)

			// and the existing database
			present.Content.ContentAmount = remaining.ContentAmount
			_, err := session.Db.Table("u_present_item").Where("user_id = ? AND id = ?", session.UserId, presentId).
				Update(&present)
			utils.CheckErr(err)
		}
		present.Content.ContentAmount = remaining.ContentAmount
		resp.LimitExceededItems.Append(present)
	}
}
