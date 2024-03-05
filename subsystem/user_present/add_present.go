package user_present

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"
)

func AddPresent(session *userdata.Session, present client.PresentItem) {
	switch present.Content.ContentType {
	case enum.ContentTypeStorySide:
		fallthrough
	case enum.ContentTypeStoryMember:
		// for story just add it
		user_content.AddContent(session, present.Content)
	default: // for items add to present box
		stat := UserPresentStat{}
		exist, err := session.Db.Table("u_present_stat").Where("user_id = ?", session.UserId).Get(&stat)
		utils.CheckErr(err)
		if !exist {
			stat.UserId = session.UserId
			stat.PresentCount = 1
			userdata.GenericDatabaseInsert(session, "u_present_stat", stat)
		} else {
			stat.PresentCount++
			_, err = session.Db.Table("u_present_stat").Where("user_id = ?", session.UserId).Update(&stat)
			utils.CheckErr(err)
		}
		present.Id = stat.PresentCount
		present.PostedAt = session.Time.Unix()
		present.IsNew = true
		// fill in the id
		userdata.GenericDatabaseInsert(session, "u_present_item", present)
	}
}
