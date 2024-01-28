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
	// TODO(trigger): Maybe this need a trigger too
	_, err := session.Db.Table("u_present").
		Where("user_id = ? AND expired_at != \"null\" AND CAST(expired_at AS BIGINT) <= ?", session.UserId, session.Time.Unix()).
		Delete(&client.PresentItem{})
	utils.CheckErr(err)

	err = session.Db.Table("u_present").Where("user_id = ?", session.UserId).
		OrderBy("posted_at DESC").Limit(PresentsPerPageCount).
		Find(&presents.Slice)
	utils.CheckErr(err)
	return presents
}
