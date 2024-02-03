package live_deck

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func saveSuit(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveLiveDeckMemberSuitRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseSuitChange {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseGacha
	}

	userLiveDeck := session.GetUserLiveDeck(req.DeckId)
	reflect.ValueOf(&userLiveDeck).Elem().Field(int(1 + req.CardIndex + 9)).Set(reflect.ValueOf(generic.NewNullable(req.SuitMasterId)))
	session.UpdateUserLiveDeck(userLiveDeck)

	// Rina-chan board toggle
	if session.Gamedata.Suit[req.SuitMasterId].Member.Id == enum.MemberMasterIdRina {
		RinaChan := user_member.GetMember(session, enum.MemberMasterIdRina)
		RinaChan.ViewStatus = req.ViewStatus
		user_member.UpdateMember(session, RinaChan)
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/liveDeck/saveSuit", saveSuit)
}
