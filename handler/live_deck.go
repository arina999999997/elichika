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
	// fmt.Println(reqBody)
	type SaveDeckAllReq struct {
		DeckID       int   `json:"deck_id"`
		CardWithSuit []int `json:"card_with_suit"`
		SquadDict    []any `json:"squad_dict"`
	}

	req := SaveDeckAllReq{}
	decoder := json.NewDecoder(strings.NewReader(reqBody))
	decoder.UseNumber()
	err := decoder.Decode(&req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	deckInfo := session.GetUserLiveDeck(req.DeckID)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	for i := 0; i < 9; i++ {
		if req.CardWithSuit[i*2+1] == 0 {
			req.CardWithSuit[i*2+1] = gamedata.Card[req.CardWithSuit[i*2]].Member.MemberInit.SuitMasterID
		}
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
	// fmt.Println(deckInfo)

	session.UpdateUserLiveDeck(deckInfo)

	for k, v := range req.SquadDict {
		if k%2 == 0 {
			partyId, err := v.(json.Number).Int64()
			if err != nil {
				panic(err)
			}
			// fmt.Println("Party ID:", partyId)

			rDictInfo, err := json.Marshal(req.SquadDict[k+1])
			utils.CheckErr(err)

			dictInfo := model.DeckSquadDict{}
			decoder := json.NewDecoder(bytes.NewReader(rDictInfo))
			decoder.UseNumber()
			err = decoder.Decode(&dictInfo)
			utils.CheckErr(err)

			partyInfo := model.UserLiveParty{}
			partyInfo.UserID = userID
			partyInfo.PartyID = int(partyId)
			partyInfo.IconMasterID, partyInfo.Name.DotUnderText = gamedata.GetLivePartyInfoByCardMasterIDs(
				dictInfo.CardMasterIDs[0], dictInfo.CardMasterIDs[1], dictInfo.CardMasterIDs[2],
			)
			partyInfo.UserLiveDeckID = req.DeckID
			partyInfo.CardMasterID1 = dictInfo.CardMasterIDs[0]
			partyInfo.CardMasterID2 = dictInfo.CardMasterIDs[1]
			partyInfo.CardMasterID3 = dictInfo.CardMasterIDs[2]
			partyInfo.UserAccessoryID1 = dictInfo.UserAccessoryIDs[0]
			partyInfo.UserAccessoryID2 = dictInfo.UserAccessoryIDs[1]
			partyInfo.UserAccessoryID3 = dictInfo.UserAccessoryIDs[2]
			session.UpdateUserLiveParty(partyInfo)
		}
	}

	respBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, respBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchLiveDeckSelect(ctx *gin.Context) {
	// return last deck for this song
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type FetchLiveDeckSelectReq struct {
		LiveDifficultyID int `json:"live_difficulty_id"`
	}
	req := FetchLiveDeckSelectReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()
	deck := session.GetLastPlayLiveDifficultyDeck(req.LiveDifficultyID)
	signBody, err := sjson.Set("{}", "last_play_live_difficulty_deck", deck)

	// utils.CheckErr(err)

	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SaveSuit(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SaveSuitReq struct {
		DeckID       int `json:"deck_id"`
		CardIndex    int `json:"card_index"`
		SuitMasterID int `json:"suit_master_id"`
		ViewStatus   int `json:"view_status"` // 2 for Rina-chan board off, 1 for everyone else
	}

	req := SaveSuitReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()
	deck := session.GetUserLiveDeck(req.DeckID)
	deckJsonByte, err := json.Marshal(deck)
	deckJson := string(deckJsonByte)
	deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", req.CardIndex), req.SuitMasterID)
	err = json.Unmarshal([]byte(deckJson), &deck)
	session.UpdateUserLiveDeck(deck)

	// Rina-chan board toggle
	if session.Gamedata.Suit[req.SuitMasterID].Member.ID == enum.MemberMasterIDRina {
		RinaChan := session.GetMember(enum.MemberMasterIDRina)
		RinaChan.ViewStatus = req.ViewStatus
		session.UpdateMember(RinaChan)
	}

	signBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SaveDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)
	type SaveDeckReq struct {
		DeckID        int    `json:"deck_id"`
		CardMasterIDs [2]int `json:"card_master_ids"`
	}
	req := SaveDeckReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	position := req.CardMasterIDs[0]
	newCardMasterID := req.CardMasterIDs[1]
	newSuitMasterID := newCardMasterID

	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	// fetch the deck and parties affected
	deck := session.GetUserLiveDeck(req.DeckID)
	deckJsonByte, err := json.Marshal(deck)
	utils.CheckErr(err)
	deckJson := string(deckJsonByte)

	oldCardMasterID := int(gjson.Get(deckJson, fmt.Sprintf("card_master_id_%d", position)).Int())
	oldSuitMasterID := int(gjson.Get(deckJson, fmt.Sprintf("suit_master_id_%d", position)).Int())
	if newCardMasterID == oldCardMasterID {
		panic("same card master id")
	}

	oldPosition := 0
	// old position != 0 then new card is in current deck, we have to swap
	gjson.Parse(deckJson).ForEach(func(key, value gjson.Result) bool {
		if strings.Contains(key.String(), "card_master_id") {
			if int(value.Int()) == newCardMasterID {
				oldPosition = int(key.String()[len(key.String())-1] - '0')
				// don't change suit_master_id if we just swap card around
				newSuitMasterID = int(gjson.Get(deckJson, fmt.Sprintf("suit_master_id_%d", oldPosition)).Int())
				return false
			}
		}
		return true
	})

	parties := []model.UserLiveParty{}
	parties = append(parties, session.GetUserLivePartyWithDeckAndCardID(req.DeckID, oldCardMasterID))
	if oldPosition != 0 {
		oldParty := session.GetUserLivePartyWithDeckAndCardID(req.DeckID, newCardMasterID)
		if oldParty.PartyID != parties[0].PartyID {
			parties = append(parties, oldParty)
		}
	}

	// change card master id in deck, then change the suit master id to default
	deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("card_master_id_%d", position), newCardMasterID)
	deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", position), newSuitMasterID)
	if oldPosition != 0 {
		deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("card_master_id_%d", oldPosition), oldCardMasterID)
		deckJson, _ = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", oldPosition), oldSuitMasterID)
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
				if int(value.Int()) == oldCardMasterID {
					partyJson, _ = sjson.Set(partyJson, key.String(), newCardMasterID)
				} else if int(value.Int()) == newCardMasterID {
					partyJson, _ = sjson.Set(partyJson, key.String(), oldCardMasterID)
				}
			}
			return true
		})

		partyInfo := gjson.Parse(partyJson)
		partyIcon, partyName := gamedata.GetLivePartyInfoByCardMasterIDs(
			int(partyInfo.Get("card_master_id_1").Int()),
			int(partyInfo.Get("card_master_id_2").Int()),
			int(partyInfo.Get("card_master_id_3").Int()),
		)

		partyJson, _ = sjson.Set(partyJson, "name.dot_under_text", partyName)
		partyJson, _ = sjson.Set(partyJson, "icon_master_id", partyIcon)
		err = json.Unmarshal([]byte(partyJson), &party)
		utils.CheckErr(err)
		session.UpdateUserLiveParty(party)
	}

	signBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeDeckNameLiveDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ChangeDeckNameReq struct {
		DeckID   int    `json:"deck_id"`
		DeckName string `json:"deck_name"`
	}
	req := ChangeDeckNameReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	liveDeck := session.GetUserLiveDeck(req.DeckID)
	liveDeck.Name.DotUnderText = req.DeckName
	session.UpdateUserLiveDeck(liveDeck)
	signBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
