package user_present

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func FetchPresentHistoryItems(session *userdata.Session) generic.List[client.PresentHistoryItem] {
	var presentHistory generic.List[client.PresentHistoryItem]
	// TODO(database): Need to delete old items somewhere
	err := session.Db.Table("u_present_history_item").Where("user_id = ?", session.UserId).
		OrderBy("history_created_at DESC").Limit(PresentsPerPageCount).
		Find(&presentHistory.Slice)
	utils.CheckErr(err)
	return presentHistory
}
