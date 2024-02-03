package handler

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// For now return fixed good skills:
// Appeal+ Ms
// Skill+ Ms
// Type Effect +
func resultLesson(ctx *gin.Context) {
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
		// 30000041, // Appeal+ (L)
		30000482, // Appeal+ (M):Group
		30000517, // Appeal+ (M):Same Attribute
		30000502, // Appeal+ (M):Same School
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

func init() {
	router.AddHandler("/lesson/resultLesson", resultLesson)
}
