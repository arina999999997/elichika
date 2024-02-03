package handler

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"reflect"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func executeLesson(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ExecuteLessonRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.ExecuteLessonResponse{
		UserModelDiff: &session.UserModel,
	}

	deck := session.GetUserLessonDeck(req.SelectedDeckId)
	repeatCount := int32(1)
	if req.IsThreeTimes {
		repeatCount = 3
	}

	for lesson := int32(1); lesson <= 4; lesson++ {
		// TODO(lesson): Generate the drops and put them into database here
		// Note that the items are already generated and added at this step based on official server response
		actions := generic.List[client.LessonMenuAction]{}
		drops := generic.List[int32]{}
		for i := 1; i <= 9; i++ {
			cardMasterId := reflect.ValueOf(deck).Field(i + 1).Interface().(generic.Nullable[int32]).Value
			actions.Append(client.LessonMenuAction{
				CardMasterId: cardMasterId,
				Position:     int32(i),
			})
		}
		if lesson <= 3 {
			for i := 1; i <= int(repeatCount*lesson*10); i++ {
				drops.Append(int32(i%2 + 1))
			}
		}
		resp.LessonMenuActions.Set(lesson%4, actions)
		resp.LessonDropRarityList.Set(lesson%4, drops)
	}
	resp.IsSubscription = true
	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/lesson/executeLesson", executeLesson)
}
