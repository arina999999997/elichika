package handler

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_lesson"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func resultLesson(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_lesson.ResultLesson(session)

	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/lesson/resultLesson", resultLesson)
}
