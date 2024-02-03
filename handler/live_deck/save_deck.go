package live_deck

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func saveDeck(ctx *gin.Context) {
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

func init() {
	router.AddHandler("/liveDeck/saveDeck", saveDeck)
}
