package handler

import (
	"bytes"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func SaveDeckAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SaveDeckAllReq struct {
		DeckId       int   `json:"deck_id"`
		CardWithSuit []int `json:"card_with_suit"`
		SquadDict    []any `json:"squad_dict"`
	}

	req := SaveDeckAllReq{}
	decoder := json.NewDecoder(strings.NewReader(reqBody))
	decoder.UseNumber()
	err := decoder.Decode(&req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	deckInfo := session.GetUserLiveDeck(req.DeckId)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	for i := 0; i < 9; i++ {
		if req.CardWithSuit[i*2+1] == 0 {
			req.CardWithSuit[i*2+1] = int(gamedata.Card[int32(req.CardWithSuit[i*2])].Member.MemberInit.SuitMasterId)
		}
	}

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseDeckEdit {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseSuitChange
	}

	deckByte, _ := json.Marshal(deckInfo)
	deckJson := string(deckByte)
	for i := 0; i < 9; i++ {
		deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("card_master_id_%d", i+1), req.CardWithSuit[i*2])
		deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", i+1), req.CardWithSuit[i*2+1])
	}

	if err := json.Unmarshal([]byte(deckJson), &deckInfo); err != nil {
		panic(err)
	}

	session.UpdateUserLiveDeck(deckInfo)

	for k, v := range req.SquadDict {
		if k%2 == 0 {
			partyId, err := v.(json.Number).Int64()
			if err != nil {
				panic(err)
			}

			rDictInfo, err := json.Marshal(req.SquadDict[k+1])
			utils.CheckErr(err)

			dictInfo := model.DeckSquadDict{}
			decoder := json.NewDecoder(bytes.NewReader(rDictInfo))
			decoder.UseNumber()
			err = decoder.Decode(&dictInfo)
			utils.CheckErr(err)

			partyInfo := model.UserLiveParty{}
			partyInfo.PartyId = int32(partyId)
			partyInfo.IconMasterId, partyInfo.Name.DotUnderText = gamedata.GetLivePartyInfoByCardMasterIds(
				int32(dictInfo.CardMasterIds[0]), int32(dictInfo.CardMasterIds[1]), int32(dictInfo.CardMasterIds[2]),
			)
			partyInfo.UserLiveDeckId = int32(req.DeckId)
			partyInfo.CardMasterId1 = int32(dictInfo.CardMasterIds[0])
			partyInfo.CardMasterId2 = int32(dictInfo.CardMasterIds[1])
			partyInfo.CardMasterId3 = int32(dictInfo.CardMasterIds[2])
			partyInfo.UserAccessoryId1 = dictInfo.UserAccessoryIds[0]
			partyInfo.UserAccessoryId2 = dictInfo.UserAccessoryIds[1]
			partyInfo.UserAccessoryId3 = dictInfo.UserAccessoryIds[2]
			session.UpdateUserLiveParty(partyInfo)
		}
	}

	respBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, respBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchLiveDeckSelect(ctx *gin.Context) {
	// return last deck for this song
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type FetchLiveDeckSelectReq struct {
		LiveDifficultyId int `json:"live_difficulty_id"`
	}
	req := FetchLiveDeckSelectReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	deck := session.GetLastPlayLiveDifficultyDeck(req.LiveDifficultyId)
	signBody, err := sjson.Set("{}", "last_play_live_difficulty_deck", deck)
	utils.CheckErr(err)

	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SaveSuit(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SaveSuitReq struct {
		DeckId       int `json:"deck_id"`
		CardIndex    int `json:"card_index"`
		SuitMasterId int `json:"suit_master_id"`
		ViewStatus   int `json:"view_status"` // 2 for Rina-chan board off, 1 for everyone else
	}

	req := SaveSuitReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseSuitChange {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseGacha
	}

	deck := session.GetUserLiveDeck(req.DeckId)
	deckJsonByte, err := json.Marshal(deck)
	utils.CheckErr(err)
	deckJson := string(deckJsonByte)
	deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", req.CardIndex), req.SuitMasterId)
	err = json.Unmarshal([]byte(deckJson), &deck)
	utils.CheckErr(err)
	session.UpdateUserLiveDeck(deck)

	// Rina-chan board toggle
	if session.Gamedata.Suit[req.SuitMasterId].Member.Id == enum.MemberMasterIdRina {
		RinaChan := session.GetMember(enum.MemberMasterIdRina)
		RinaChan.ViewStatus = int32(req.ViewStatus)
		session.UpdateMember(RinaChan)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SaveDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SaveDeckReq struct {
		DeckId        int    `json:"deck_id"`
		CardMasterIds [2]int `json:"card_master_ids"`
	}
	req := SaveDeckReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	position := req.CardMasterIds[0]
	newCardMasterId := req.CardMasterIds[1]
	newSuitMasterId := newCardMasterId

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	// fetch the deck and parties affected
	deck := session.GetUserLiveDeck(req.DeckId)
	deckJsonByte, err := json.Marshal(deck)
	utils.CheckErr(err)
	deckJson := string(deckJsonByte)

	oldCardMasterId := int(gjson.Get(deckJson, fmt.Sprintf("card_master_id_%d", position)).Int())
	oldSuitMasterId := int(gjson.Get(deckJson, fmt.Sprintf("suit_master_id_%d", position)).Int())
	if newCardMasterId == oldCardMasterId {
		panic("same card master id")
	}

	oldPosition := 0
	// old position != 0 then new card is in current deck, we have to swap
	gjson.Parse(deckJson).ForEach(func(key, value gjson.Result) bool {
		if strings.Contains(key.String(), "card_master_id") {
			if int(value.Int()) == newCardMasterId {
				oldPosition = int(key.String()[len(key.String())-1] - '0')
				// don't change suit_master_id if we just swap card around
				newSuitMasterId = int(gjson.Get(deckJson, fmt.Sprintf("suit_master_id_%d", oldPosition)).Int())
				return false
			}
		}
		return true
	})

	parties := []model.UserLiveParty{}
	parties = append(parties, session.GetUserLivePartyWithDeckAndCardId(req.DeckId, oldCardMasterId))
	if oldPosition != 0 {
		oldParty := session.GetUserLivePartyWithDeckAndCardId(req.DeckId, newCardMasterId)
		if oldParty.PartyId != parties[0].PartyId {
			parties = append(parties, oldParty)
		}
	}

	// change card master id in deck, then change the suit master id to default
	deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("card_master_id_%d", position), newCardMasterId)
	deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", position), newSuitMasterId)
	if oldPosition != 0 {
		deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("card_master_id_%d", oldPosition), oldCardMasterId)
		deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", oldPosition), oldSuitMasterId)
	}
	err = json.Unmarshal([]byte(deckJson), &deck)
	utils.CheckErr(err)
	session.UpdateUserLiveDeck(deck)

	for _, party := range parties {
		partyJsonByte, err := json.Marshal(party)
		utils.CheckErr(err)
		partyJson := string(partyJsonByte)
		// change the live party and update the names
		gjson.Parse(partyJson).ForEach(func(key, value gjson.Result) bool {
			if strings.Contains(key.String(), "card_master_id") {
				if int(value.Int()) == oldCardMasterId {
					partyJson, _ = sjson.Set(partyJson, key.String(), newCardMasterId)
				} else if int(value.Int()) == newCardMasterId {
					partyJson, _ = sjson.Set(partyJson, key.String(), oldCardMasterId)
				}
			}
			return true
		})

		partyInfo := gjson.Parse(partyJson)
		partyIcon, partyName := gamedata.GetLivePartyInfoByCardMasterIds(
			int32(partyInfo.Get("card_master_id_1").Int()),
			int32(partyInfo.Get("card_master_id_2").Int()),
			int32(partyInfo.Get("card_master_id_3").Int()),
		)

		partyJson, _ = sjson.Set(partyJson, "name.dot_under_text", partyName)
		partyJson, _ = sjson.Set(partyJson, "icon_master_id", partyIcon)
		err = json.Unmarshal([]byte(partyJson), &party)
		utils.CheckErr(err)
		session.UpdateUserLiveParty(party)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeDeckNameLiveDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ChangeDeckNameReq struct {
		DeckId   int    `json:"deck_id"`
		DeckName string `json:"deck_name"`
	}
	req := ChangeDeckNameReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	liveDeck := session.GetUserLiveDeck(req.DeckId)
	liveDeck.Name.DotUnderText = req.DeckName
	session.UpdateUserLiveDeck(liveDeck)
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
