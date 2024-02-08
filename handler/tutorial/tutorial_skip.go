package tutorial

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_deck"
	"elichika/subsystem/user_reference_book"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func tutorialSkip(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SkipTutorialRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	// TODO(tutorial): Replicate the steps we would go through instead of just skipping.
	// if tutorial is skipped, then it can send stuff to update the live deck which would have been updated during the tutorial itself
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTutorialEnd {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTutorialEnd
		session.UserStatus.TutorialEndAt = session.Time.Unix()
	}

	user_live_deck.SaveUserLiveDeck(session, 1, req.CardWithSuitDict, req.SquadDict)

	{ // unlock systems
		session.UnlockScene(enum.UnlockSceneTypeLesson, enum.UnlockSceneStatusOpen)
		session.UnlockScene(enum.UnlockSceneTypeFreeLive, enum.UnlockSceneStatusOpen)
		session.UnlockScene(enum.UnlockSceneTypeAccessory, enum.UnlockSceneStatusOpen)
		session.UnlockScene(enum.UnlockSceneTypeStoryMember, enum.UnlockSceneStatusOpen)
		session.UnlockScene(enum.UnlockSceneTypeEvent, enum.UnlockSceneStatusOpen)
		session.UnlockScene(enum.UnlockSceneTypeReferenceBookSelect, enum.UnlockSceneStatusOpen)
	}
	{
		// mark lesson as finished
		user_reference_book.InsertUserReferenceBook(session, 1001)
		user_reference_book.InsertUserReferenceBook(session, 1002)
		user_reference_book.InsertUserReferenceBook(session, 1003)
		user_reference_book.InsertUserReferenceBook(session, 1004)
	}

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tutorial/tutorialSkip", tutorialSkip)
}
