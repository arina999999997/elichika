package live_deck

import (
	// "bytes"
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_deck"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func SaveDeckAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveLiveDeckAllRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_live_deck.UpdateUserLiveDeck(session, req.DeckId, req.CardWithSuit, req.SquadDict)

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseDeckEdit {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseSuitChange
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func FetchLiveDeckSelect(ctx *gin.Context) {
	// return last deck for this song
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchLiveDeckSelectRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, response.FetchLiveDeckSelectResponse{
		LastPlayLiveDifficultyDeck: session.GetLastPlayLiveDifficultyDeck(req.LiveDifficultyId),
	})
}

func SaveSuit(ctx *gin.Context) {
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

func SaveDeck(ctx *gin.Context) {

	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveLiveDeckCardsRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	for position, cardMasterId := range req.CardMasterIds.Map {
		// there should only be 1 pair here
		deck := session.GetUserLiveDeck(req.DeckId)
		replacedCardMasterId := reflect.ValueOf(deck).Field(1 + int(position)).Interface().(generic.Nullable[int32]).Value
		replacedSuitMasterId := reflect.ValueOf(deck).Field(1 + int(position) + 9).Interface().(generic.Nullable[int32]).Value
		suitMasterId := int32(0)
		oldPosition := int32(0)
		for i := 1; i <= 9; i++ {
			currentCardMasterId := reflect.ValueOf(deck).Field(1 + i).Interface().(generic.Nullable[int32]).Value
			if currentCardMasterId == *cardMasterId {
				oldPosition = int32(i)
				suitMasterId = reflect.ValueOf(deck).Field(1 + i + 9).Interface().(generic.Nullable[int32]).Value
				break
			}
		}

		reflect.ValueOf(&deck).Elem().Field(1 + int(position)).Set(reflect.ValueOf(generic.NewNullable(*cardMasterId)))
		if suitMasterId == 0 {
			suitMasterId = gamedata.Card[*cardMasterId].Member.MemberInit.SuitMasterId
		}
		reflect.ValueOf(&deck).Elem().Field(1 + int(position) + 9).Set(reflect.ValueOf(generic.NewNullable(suitMasterId)))

		if oldPosition != 0 {
			// swap the cards
			reflect.ValueOf(&deck).Elem().Field(1 + int(oldPosition)).Set(reflect.ValueOf(generic.NewNullable(replacedCardMasterId)))
			reflect.ValueOf(&deck).Elem().Field(1 + int(oldPosition) + 9).Set(reflect.ValueOf(generic.NewNullable(replacedSuitMasterId)))
		}
		session.UpdateUserLiveDeck(deck)
		// also need to sync up the live party
		parties := []client.UserLiveParty{}
		parties = append(parties, session.GetUserLivePartyWithDeckAndCardId(req.DeckId, replacedCardMasterId))
		if oldPosition != 0 {
			oldParty := session.GetUserLivePartyWithDeckAndCardId(req.DeckId, *cardMasterId)
			if oldParty.PartyId != parties[0].PartyId {
				parties = append(parties, oldParty)
			}
		}

		for _, party := range parties {
			for i := 1; i <= 3; i++ {
				partyCurrentCardMasterId := reflect.ValueOf(party).Field(3 + i).Interface().(generic.Nullable[int32]).Value
				if partyCurrentCardMasterId == *cardMasterId {
					reflect.ValueOf(&party).Elem().Field(3 + i).Set(reflect.ValueOf(generic.NewNullable(replacedCardMasterId)))
				} else if partyCurrentCardMasterId == replacedCardMasterId {
					reflect.ValueOf(&party).Elem().Field(3 + i).Set(reflect.ValueOf(generic.NewNullable(*cardMasterId)))
				}
			}

			party.IconMasterId, party.Name.DotUnderText = gamedata.GetLivePartyInfoByCardMasterIds(
				party.CardMasterId1.Value, party.CardMasterId2.Value, party.CardMasterId3.Value)
			session.UpdateUserLiveParty(party)
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func ChangeDeckNameLiveDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ChangeNameLiveDeckRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	liveDeck := session.GetUserLiveDeck(req.DeckId)
	liveDeck.Name.DotUnderText = req.DeckName
	session.UpdateUserLiveDeck(liveDeck)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/liveDeck/changeDeckNameLiveDeck", ChangeDeckNameLiveDeck)
	router.AddHandler("/liveDeck/fetchLiveDeckSelect", FetchLiveDeckSelect)
	router.AddHandler("/liveDeck/saveDeck", SaveDeck)
	router.AddHandler("/liveDeck/saveDeckAll", SaveDeckAll)
	router.AddHandler("/liveDeck/saveSuit", SaveSuit)
}
