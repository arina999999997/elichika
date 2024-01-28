package user_content

import (
	"elichika/client"
	"elichika/config"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserContent(session *userdata.Session, contentType, contentId int32) client.Content {
	_, exist := session.UserContentDiffs[contentType]
	if !exist {
		session.UserContentDiffs[contentType] = make(map[int32]client.Content)
	}
	content, exist := session.UserContentDiffs[contentType][contentId]
	if exist {
		return content
	}
	// load from db
	exist, err := session.Db.Table("u_content").Where("user_id = ? AND content_type = ? AND content_id = ?",
		session.UserId, contentType, contentId).Get(&content)
	utils.CheckErr(err)
	if !exist {
		content = client.Content{
			ContentType:   contentType,
			ContentId:     contentId,
			ContentAmount: *config.Conf.DefaultContentAmount,
		}
	}
	return content
}

func GetUserContentByContent(session *userdata.Session, content client.Content) client.Content {
	return GetUserContent(session, content.ContentType, content.ContentId)
}
