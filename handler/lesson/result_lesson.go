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

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_lesson.ResultLesson(session)

	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/lesson/resultLesson", resultLesson)
}
