package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/locale"
	"elichika/userdata"
	"elichika/utils"

	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func StartUpAuthorizationKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	if err != nil {
		panic(err)
	}
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	newKey := utils.Xor(randomBytes, []byte(config.SessionKey))
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	return newKey64
}

func LoginSessionKey(mask64 string) string {
	mask, err := base64.StdEncoding.DecodeString(mask64)
	utils.CheckErr(err)
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	serverEventReceiverKey, err := hex.DecodeString(config.ServerEventReceiverKey)
	utils.CheckErr(err)
	jaKey, err := hex.DecodeString(config.JaKey)
	utils.CheckErr(err)
	newKey := utils.Xor(randomBytes, []byte(config.SessionKey))
	newKey = utils.Xor(newKey, serverEventReceiverKey)
	newKey = utils.Xor(newKey, jaKey)
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	return newKey64
}

func StartUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)
	type StartUpReq struct {
		Mask                        string `json:"mask"`
		ResemaraDetectionIdentifier string `json:"resemara_detection_identifier"` // reset marathon (reroll)
		TimeDifference              int    `json:"time_difference"`               // second different from utc + 0
		RecaptchaToken              string `json:"recaptcha_token"`               // not necessary
	}
	req := StartUpReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	type StartUpResp struct {
		UserID           int    `json:"user_id"`
		AuthorizationKey string `json:"authorization_key"`
	}
	respObj := StartUpResp{}
	respObj.UserID = userdata.CreateNewAccount(ctx, -1, "")
	respObj.AuthorizationKey = StartUpAuthorizationKey(req.Mask)
	startupBody, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(startupBody), ctx.MustGet("locale").(*locale.Locale).StartUpKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func Login(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var mask64 string
	var err error
	req := gjson.Parse(reqBody)
	req.ForEach(func(key, value gjson.Result) bool {
		if value.Get("mask").String() != "" {
			mask64 = value.Get("mask").String()
			return false
		}
		return true
	})
	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()
	if session == nil {
		userdata.CreateNewAccount(ctx, UserID, "")
		session = userdata.GetSession(ctx, UserID)
		defer session.Close()
	}
	session.UserStatus.LastLoginAt = time.Now().Unix()

	loginBody := session.Finalize(GetData("login.json"), "user_model")
	loginBody, _ = sjson.Set(loginBody, "session_key", LoginSessionKey(mask64))
	loginBody, _ = sjson.Set(loginBody, "last_timestamp", time.Now().UnixMilli())

	/* ======== UserData ======== */
	fmt.Println("User logins: ", UserID)

	// live decks
	dbLiveDecks := session.GetAllLiveDecks()
	if len(dbLiveDecks) == 0 {
		panic("no live deck found")
	}
	userLiveDecks := []any{}
	for _, liveDeckInfo := range dbLiveDecks {
		userLiveDecks = append(userLiveDecks, liveDeckInfo.UserLiveDeckID)
		userLiveDecks = append(userLiveDecks, liveDeckInfo)
	}
	loginBody, _ = sjson.Set(loginBody, "user_model.user_live_deck_by_id", userLiveDecks)

	dbLiveParties := session.GetAllLiveParties()
	if len(dbLiveParties) == 0 {
		panic("no live party")
	}
	userLiveParties := []any{}
	for _, livePartyInfo := range dbLiveParties {
		userLiveParties = append(userLiveParties, livePartyInfo.PartyID)
		userLiveParties = append(userLiveParties, livePartyInfo)
	}
	loginBody, _ = sjson.Set(loginBody, "user_model.user_live_party_by_id", userLiveParties)

	// member settings
	dbMembers := session.GetAllMembers()
	if len(dbMembers) == 0 {
		panic("no member found")
	}
	var userMembers []any
	for _, memberInfo := range dbMembers {
		userMembers = append(userMembers, memberInfo.MemberMasterID)
		userMembers = append(userMembers, memberInfo)
	}
	loginBody, _ = sjson.Set(loginBody, "user_model.user_member_by_member_id", userMembers)

	// member love panel settings
	dbLovePanels := session.GetAllMemberLovePanels()
	if len(dbLovePanels) == 0 {
		panic("no member love panel found")
	}
	loginBody, _ = sjson.Set(loginBody, "member_love_panels", dbLovePanels)

	// lesson decks
	dbLessonDecks := session.GetAllLessonDecks()
	if len(dbLessonDecks) == 0 {
		panic("no lesson deck")
	}

	userLessonDecks := []any{}
	for _, userLessonDeck := range dbLessonDecks {
		userLessonDecks = append(userLessonDecks, userLessonDeck.UserLessonDeckID)
		userLessonDecks = append(userLessonDecks, userLessonDeck)
	}
	loginBody, _ = sjson.Set(loginBody, "user_model.user_lesson_deck_by_id", userLessonDecks)

	// user cards
	dbCards := session.GetAllCards()
	if len(dbCards) == 0 {
		panic("no card")
	}

	userCards := []any{}
	for _, userCard := range dbCards {
		userCards = append(userCards, userCard.CardMasterID)
		userCards = append(userCards, userCard)
	}
	loginBody, _ = sjson.Set(loginBody, "user_model.user_card_by_card_id", userCards)

	// user suits
	dbSuits := session.GetAllSuits()
	if len(dbSuits) == 0 {
		panic("no suit")
	}

	userSuits := []any{}
	for _, userSuit := range dbSuits {
		userSuits = append(userSuits, userSuit.SuitMasterID)
		userSuits = append(userSuits, userSuit)
	}
	loginBody, err = sjson.Set(loginBody, "user_model.user_suit_by_suit_id", userSuits)
	utils.CheckErr(err)

	// user accessory
	dbUserAccessories := session.GetAllUserAccessories()
	userAccessories := []any{}
	for _, userAccessory := range dbUserAccessories {
		userAccessories = append(userAccessories, userAccessory.UserAccessoryID)
		userAccessories = append(userAccessories, userAccessory)
	}
	loginBody, _ = sjson.Set(loginBody, "user_model.user_accessory_by_user_accessory_id", userAccessories)

	// song records
	// if return empty, all the song are unlocked, except for bond episide unlocked song
	dbLiveRecords := session.GetAllLiveDifficultyRecords()
	userLiveRecords := []any{}
	for _, userLiveRecord := range dbLiveRecords {
		userLiveRecords = append(userLiveRecords, userLiveRecord.LiveDifficultyID)
		userLiveRecords = append(userLiveRecords, userLiveRecord)
	}
	loginBody, err = sjson.Set(loginBody, "user_model.user_live_difficulty_by_difficulty_id", userLiveRecords)
	utils.CheckErr(err)

	// playlist
	dbPlaylist := session.GetUserPlayList()
	loginBody, err = sjson.Set(loginBody, "user_model.user_play_list_by_id", dbPlaylist)
	utils.CheckErr(err)

	// triggers
	triggersBasics := session.GetAllTriggerBasics()
	loginBody, err = sjson.Set(loginBody, "user_model.user_info_trigger_basic_by_trigger_id", triggersBasics)
	utils.CheckErr(err)
	triggersCardGradeUps := session.GetAllTriggerCardGradeUps()
	loginBody, err = sjson.Set(loginBody, "user_model.user_info_trigger_card_grade_up_by_trigger_id", triggersCardGradeUps)
	utils.CheckErr(err)

	/* ======== UserData ======== */
	resp := SignResp(ctx, loginBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
