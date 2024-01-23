package tutorial

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_deck"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func CorePlayableEnd(ctx *gin.Context) {
	// there's no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase != enum.TutorialPhaseCorePlayable {
		panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseTimingAdjuster

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TimingAdjusterEnd(ctx *gin.Context) {
	// there's no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTimingAdjuster {
		panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseFavoriateMember

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func PhaseEnd(ctx *gin.Context) {
	// there's no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase != enum.TutorialPhaseFinal {
		panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseTutorialEnd
	session.UserStatus.TutorialEndAt = session.Time.Unix()

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TutorialSkip(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SkipTutorialRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// if tutorial is skipped, then it can send stuff to update the live deck which would have been updated during the tutorial itself
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTutorialEnd {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTutorialEnd
		session.UserStatus.TutorialEndAt = time.Now().Unix()
	}

	user_live_deck.UpdateUserLiveDeck(session, 1, req.CardWithSuitDict, req.SquadDict)

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tutorial/corePlayableEnd", CorePlayableEnd)
	router.AddHandler("/tutorial/phaseEnd", PhaseEnd)
	router.AddHandler("/tutorial/tutorialSkip", TutorialSkip)
	router.AddHandler("/tutorial/timingAdjusterEnd", TimingAdjusterEnd)
}
