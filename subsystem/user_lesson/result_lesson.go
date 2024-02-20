package user_lesson

import (
	"elichika/client/response"
	"elichika/userdata"
	"elichika/utils"
)

func ResultLesson(session *userdata.Session) response.LessonResultResponse {
	resp := response.LessonResultResponse{}
	exists, err := session.Db.Table("u_lesson").Where("user_id = ?", session.UserId).Get(&resp)
	utils.CheckErrMustExist(err, exists)

	resp.UserModelDiff = &session.UserModel
	return resp
}
