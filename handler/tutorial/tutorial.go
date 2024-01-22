package tutorial

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"reflect"
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

	// TODO(refactor): Used some common code to do this instead

	userLiveDeck := session.GetUserLiveDeck(1)
	for position, cardMasterId := range req.CardWithSuitDict.Order {
		suitMasterId := *req.CardWithSuitDict.GetOnly(cardMasterId)
		if !suitMasterId.HasValue {
			// TODO: maybe we can assign the suit of the card instead
			suitMasterId = generic.NewNullable(session.Gamedata.Card[cardMasterId].Member.MemberInit.SuitMasterId)
		}
		reflect.ValueOf(&userLiveDeck).Elem().Field(position + 2).Set(reflect.ValueOf(generic.NewNullable(cardMasterId)))
		reflect.ValueOf(&userLiveDeck).Elem().Field(position + 2 + 9).Set(reflect.ValueOf(suitMasterId))
	}
	session.UpdateUserLiveDeck(userLiveDeck)
	for partyId, liveSquad := range req.SquadDict.Map {
		userLiveParty := client.UserLiveParty{
			PartyId:        partyId,
			UserLiveDeckId: 1,
		}
		userLiveParty.IconMasterId, userLiveParty.Name.DotUnderText = session.Gamedata.GetLivePartyInfoByCardMasterIds(
			liveSquad.CardMasterIds.Slice[0], liveSquad.CardMasterIds.Slice[1], liveSquad.CardMasterIds.Slice[2])
		for position := 0; position < 3; position++ {
			reflect.ValueOf(&userLiveParty).Elem().Field(position + 4).Set(
				reflect.ValueOf(generic.NewNullable(liveSquad.CardMasterIds.Slice[position])))
			reflect.ValueOf(&userLiveParty).Elem().Field(position + 4 + 3).Set(
				reflect.ValueOf(liveSquad.UserAccessoryIds.Slice[position]))
		}
		session.UpdateUserLiveParty(userLiveParty)
	}

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
