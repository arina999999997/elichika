package user_present

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func FetchPresentItems(session *userdata.Session) generic.List[client.PresentItem] {
	var presents generic.List[client.PresentItem]
	// delete the expired item because the client can't claim them
	// TODO(database): this is ugly but to make it better we would need to either get rid of xorm or control it much better
	_, err := session.Db.Table("u_present_item").
		Where("user_id = ? AND expired_at != \"null\" AND CAST(expired_at AS BIGINT) <= ?", session.UserId, session.Time.Unix()).
		Delete(&client.PresentItem{})
	utils.CheckErr(err)

	err = session.Db.Table("u_present_item").Where("user_id = ?", session.UserId).
		OrderBy("id DESC").Limit(PresentsPerPageCount).
		Find(&presents.Slice)
	utils.CheckErr(err)

	// TODO(database): This set everything, which arguablly is the correct thing to do but it might be slow

	_, err = session.Db.Exec("UPDATE u_present_item SET is_new = 0 WHERE user_id = ?", session.UserId)
	utils.CheckErr(err)
	return presents
}
