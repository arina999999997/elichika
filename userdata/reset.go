package userdata

import (
	"elichika/item"
	"elichika/utils"
)

func (session *Session) RemoveUserProgress(table string) {
	count, err := session.Db.Table(table).Where("user_id = ?", session.UserId).Delete()
	utils.CheckErr(err)
	if table == "u_story_event_history" {
		session.AddResource(item.MemoryKey.Amount(int32(count)))
	}
}

func (session *Session) MarkIsNew(table string, isNew bool) {
	_, err := session.Db.Table(table).Where("user_id = ?", session.UserId).Update(map[string]interface{}{"is_new": isNew})
	utils.CheckErr(err)

}
