package handler

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_lesson"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func executeLesson(ctx *gin.Context) {
	req := request.ExecuteLessonRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_lesson.ExecuteLesson(session, req)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/lesson/executeLesson", executeLesson)
}
