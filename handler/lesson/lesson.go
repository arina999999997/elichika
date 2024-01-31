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

func ExecuteLesson(ctx *gin.Context) {
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

// For now return fixed good skills:
// Appeal+ Ms
// Skill+ Ms
// Type Effect +
func ResultLesson(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// TODO(lesson): Get the data from the database and fill it in
	// Note that the items are already decided and added to the account by this point, but we need to display it
	// Also the selected deck id is relevant so need to save it
	resp := response.LessonResultResponse{
		UserModelDiff:  &session.UserModel,
		SelectedDeckId: 1,
	}
	// can only return 12 max
	skills := []int32{
		30000041, // Appeal+ (L)
		30000482, // Appeal+ (M):Group
		30000517, // Appeal+ (M):Same Attribute
		// 30000502, // Appeal+ (M):Same School
		30000507, // Appeal+ (M):Same Strategy
		30000492, // Appeal+ (M):Same Year
		30000512, // Appeal+ (M):Type
		30000044, // Skill Activation %+ (L)
		30000485, // Skill Activation %+ (M):Group
		30000520, // Skill Activation %+ (M):Same Attribute
		// 30000505, // Skill Activation %+ (M):Same School
		30000510, // Skill Activation %+ (M):Same Strategy
		30000495, // Skill Activation %+ (M):Same Year
		// 30000515, // Skill Activation %+ (M):Type
		30000045, // Type Effect+ (M)
	}
	for position := int32(1); position <= 9; position++ {
		for _, skillId := range skills {
			resp.DropSkillList.Append(client.LessonResultDropPassiveSkill{
				Position:       position,
				PassiveSkillId: skillId,
			})
		}
	}

	common.JsonResponse(ctx, resp)
}

func SkillEditResult(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SkillEditResultRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for cardMasterId, selectedSkills := range req.SelectedSkillIds.Map {
		card := session.GetUserCard(cardMasterId)
		for i, skillId := range selectedSkills.Slice {
			switch i {
			case 0:
				card.AdditionalPassiveSkill1Id = skillId
			case 1:
				card.AdditionalPassiveSkill2Id = skillId
			case 2:
				card.AdditionalPassiveSkill3Id = skillId
			case 3:
				card.AdditionalPassiveSkill4Id = skillId
			}
		}
		session.UpdateUserCard(card) // this is always updated even if no skill change happen
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SaveDeckLesson(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveLessonDeckRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userLessonDeck := session.GetUserLessonDeck(req.DeckId)
	for position, cardMasterId := range req.CardMasterIds.Map {
		reflect.ValueOf(&userLessonDeck).Elem().Field(int(position) + 1).Set(reflect.ValueOf(*cardMasterId))
	}
	session.UpdateLessonDeck(userLessonDeck)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func ChangeDeckNameLessonDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ChangeNameLessonDeckRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	lessonDeck := session.GetUserLessonDeck(req.DeckId)
	lessonDeck.Name = req.DeckName
	session.UpdateLessonDeck(lessonDeck)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	// TODO(refactor): move to individual files. 
	router.AddHandler("/lesson/executeLesson", ExecuteLesson)
	router.AddHandler("/lesson/resultLesson", ResultLesson)
	router.AddHandler("/lesson/saveDeck", SaveDeckLesson)
	router.AddHandler("/lesson/skillEditResult", SkillEditResult)
	router.AddHandler("/lesson/changeDeckNameLessonDeck", ChangeDeckNameLessonDeck)
}
