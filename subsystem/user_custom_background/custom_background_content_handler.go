package user_custom_background

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"
)

func customBackgroundContentHandler(session *userdata.Session, content *client.Content) any {
	content.ContentAmount = 0
	exists, err := session.Db.Table("u_custom_background").
		Where("user_id = ? AND custom_background_master_id = ?", session.UserId, content.ContentId).Exist(&client.UserCustomBackground{})
	utils.CheckErr(err)
	if !exists {
		userdata.GenericDatabaseInsert(session, "u_custom_background", client.UserCustomBackground{
			CustomBackgroundMasterId: content.ContentId,
			IsNew: true,
		})
	}
	return nil 
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeCustomBackground, customBackgroundContentHandler)
}
