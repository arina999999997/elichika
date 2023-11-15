package userdata

import (
	"elichika/gamedata"
	"elichika/model"
	"elichika/protocol/response"
	"elichika/utils"

	// "encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"xorm.io/xorm"
)

// A session is a complete transation between server and client
// so 1 session per request
// A session fetch the data needs to be modified, and sometime modify the data if the code is shared between handlers.
// session can use ctx to get things like user id / master db, but it should not make any network operation

type Session struct {
	Db                        *xorm.Session
	Ctx                       *gin.Context
	Gamedata                  *gamedata.Gamedata
	UserStatus                model.UserStatus
	CardDiffs                 map[int]model.UserCard
	UserMemberDiffs           map[int]model.UserMember
	UserLessonDeckDiffs       map[int]model.UserLessonDeck
	UserLiveDeckDiffs         map[int]model.UserLiveDeck
	UserLivePartyDiffs        map[int]model.UserLiveParty
	UserMemberLovePanelDiffs  map[int]model.UserMemberLovePanel
	UserLiveDifficultyDiffs   map[int]model.UserLiveDifficulty
	UserAccessoryDiffs        map[int64]model.UserAccessory
	UserResourceDiffs         map[int](map[int]UserResource) // content_type then content_id
	UserSuitDiffs             []model.UserSuit
	TriggerCardGradeUps       []any
	TriggerBasics             []any
	TriggerMemberLoveLevelUps []any
	UserModelCommon           response.UserModelCommon
}

// Push update into the db and create the diff
// The actual response depend on the API, but they often contain the diff somewhere
// The mainKey is the key to the diff
func (session *Session) Finalize(jsonBody string, mainKey string) string {
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_member_by_member_id", session.FinalizeUserMemberDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_card_by_card_id", session.FinalizeCardDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_lesson_deck_by_id", session.FinalizeUserLessonDeckDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_live_deck_by_id", session.FinalizeUserLiveDeckDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_live_party_by_id", session.FinalizeUserLivePartyDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_suit_by_suit_id", session.FinalizeUserSuitDiffs())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_accessory_by_user_accessory_id", session.FinalizeUserAccessories())
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_live_difficulty_by_difficulty_id", session.FinalizeLiveDifficulties())
	resourceKeys, resourceValues := session.FinalizeUserResources()
	for i, _ := range resourceKeys {
		jsonBody, _ = sjson.SetRaw(jsonBody, mainKey+"."+resourceKeys[i], resourceValues[i])
	}
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_status", session.FinalizeUserInfo())
	session.UserModelCommon.UserStatus = session.UserStatus
	memberLovePanels := session.FinalizeMemberLovePanelDiffs()
	if len(memberLovePanels) != 0 {
		jsonBody, _ = sjson.Set(jsonBody, "member_love_panels", memberLovePanels)
	}
	session.Db.Commit()
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_info_trigger_card_grade_up_by_trigger_id", session.TriggerCardGradeUps)
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_info_trigger_basic_by_trigger_id", session.TriggerBasics)
	jsonBody, _ = sjson.Set(jsonBody, mainKey+".user_info_trigger_member_love_level_up_by_trigger_id", session.TriggerMemberLoveLevelUps)
	if len(jsonBody) < 100000 {
	fmt.Println("Before\n", jsonBody)
	jsonBody, _ = sjson.Set(jsonBody, mainKey, session.UserModelCommon)
	fmt.Println("After\n", jsonBody)
	}
	// testJson, err := json.Marshal(session.UserModelCommon)
	// utils.CheckErr(err)
	// fmt.Println(string(testJson))
	return jsonBody
}

func (session *Session) Close() {
	// fmt.Printf("close: %p\n", session)
	if session == nil {
		return
	}
	session.Db.Close()
}

func (session *Session) FinalizeUserInfo() model.UserStatus {
	_, err := session.Db.Table("u_info").Where("user_id = ?", session.UserStatus.UserID).Update(&session.UserStatus)
	if err != nil {
		panic(err)
	}
	return session.UserStatus
}

func GetSession(ctx *gin.Context, userID int) *Session {
	s := Session{}
	s.Ctx = ctx
	s.Gamedata = ctx.MustGet("gamedata").(*gamedata.Gamedata)
	s.Db = Engine.NewSession()
	err := s.Db.Begin()
	utils.CheckErr(err)
	// fmt.Printf("session: %p\n", &s)
	s.UserStatus.UserID = userID
	exists, err := s.Db.Table("u_info").Where("user_id = ?", userID).Get(&s.UserStatus)
	utils.CheckErr(err)
	if !exists {
		s.Close()
		return nil
	}

	s.CardDiffs = make(map[int]model.UserCard)
	s.UserMemberDiffs = make(map[int]model.UserMember)
	s.UserLessonDeckDiffs = make(map[int]model.UserLessonDeck)
	s.UserLiveDeckDiffs = make(map[int]model.UserLiveDeck)
	s.UserLivePartyDiffs = make(map[int]model.UserLiveParty)
	s.UserMemberLovePanelDiffs = make(map[int]model.UserMemberLovePanel)
	s.UserLiveDifficultyDiffs = make(map[int]model.UserLiveDifficulty)
	s.UserAccessoryDiffs = make(map[int64]model.UserAccessory)
	s.TriggerCardGradeUps = make([]any, 0)
	s.TriggerBasics = make([]any, 0)
	s.TriggerMemberLoveLevelUps = make([]any, 0)
	s.UserSuitDiffs = make([]model.UserSuit, 0)
	s.UserResourceDiffs = make(map[int](map[int]UserResource))
	return &s
}
