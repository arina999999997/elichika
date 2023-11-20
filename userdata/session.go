package userdata

import (
	"elichika/gamedata"
	"elichika/generic"
	"elichika/model"
	"elichika/protocol/response"
	"elichika/utils"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"xorm.io/xorm"
)

// A session is a complete transation between server and client
// so 1 session per request
// A session fetch the data needs to be modified, and sometime modify the data if the code is shared between handlers.
// session can use ctx to get things like user id / master db, but it should not make any network operation

type Session struct {
	Db                                  *xorm.Session
	Ctx                                 *gin.Context
	Gamedata                            *gamedata.Gamedata
	UserStatus                          *model.UserStatus // link to UserModel.UserStatus
	UserCardMapping                     generic.ObjectByObjectIDMapping[model.UserCard]
	UserMemberMapping                   generic.ObjectByObjectIDMapping[model.UserMember]
	UserLessonDeckMapping               generic.ObjectByObjectIDMapping[model.UserLessonDeck]
	UserLiveDeckMapping                 generic.ObjectByObjectIDMapping[model.UserLiveDeck]
	UserLivePartyMapping                generic.ObjectByObjectIDMapping[model.UserLiveParty]
	UserAccessoryMapping                generic.ObjectByObjectIDMapping[model.UserAccessory]
	UserLiveDifficultyMapping           generic.ObjectByObjectIDMapping[model.UserLiveDifficulty]
	UserTriggerCardGradeUpMapping       generic.ObjectByObjectIDMapping[model.TriggerCardGradeUp]
	UserTriggerBasicMapping             generic.ObjectByObjectIDMapping[model.TriggerBasic]
	UserTriggerMemberLoveLevelUpMapping generic.ObjectByObjectIDMapping[model.TriggerMemberLoveLevelUp]
	UserMemberLovePanelDiffs            map[int]model.UserMemberLovePanel
	UserResourceDiffs                   map[int](map[int]UserResource) // content_type then content_id

	// for now only store delta patch, i.e. user_model_diff
	// should be fine until we want to keep user state entirely in ram
	UserModel response.UserModel

	// unix nano timestamps, mixed with counting up trigger id
	// maybe expands to other things too
	TimeStamp   int64
	UniqueCount int64
}

// Push update into the db and create the diff
// The actual response depend on the API, but they often contain the diff somewhere
// The mainKey is the key to the diff
func (session *Session) Finalize(jsonBody string, mainKey string) string {
	memberLovePanels := session.FinalizeMemberLovePanelDiffs()
	var err error
	if len(memberLovePanels) != 0 {
		jsonBody, err = sjson.Set(jsonBody, "member_love_panels", memberLovePanels)
		utils.CheckErr(err)
	}
	for _, finalizer := range finalizers {
		finalizer(session)
	}
	session.Db.Commit()
	jsonBody, err = sjson.Set(jsonBody, mainKey, session.UserModel)
	utils.CheckErr(err)
	return jsonBody
}

func (session *Session) Close() {
	// fmt.Printf("close: %p\n", session)
	if session == nil {
		return
	}
	session.Db.Close()
}

func userStatusFinalizer(session *Session) {
	_, err := session.Db.Table("u_info").Where("user_id = ?", session.UserStatus.UserID).AllCols().Update(session.UserStatus)
	utils.CheckErr(err)
}
func init() {
	addFinalizer(userStatusFinalizer)
}

func GetSession(ctx *gin.Context, userID int) *Session {
	s := Session{}
	s.Ctx = ctx
	s.Gamedata = ctx.MustGet("gamedata").(*gamedata.Gamedata)
	s.Db = Engine.NewSession()
	err := s.Db.Begin()
	utils.CheckErr(err)

	exists, err := s.Db.Table("u_info").Where("user_id = ?", userID).Get(&s.UserModel.UserStatus)
	utils.CheckErr(err)
	if !exists {
		s.Close()
		return nil
	}
	s.UserStatus = &s.UserModel.UserStatus
	s.UserResourceDiffs = make(map[int](map[int]UserResource))

	s.UserMemberLovePanelDiffs = make(map[int]model.UserMemberLovePanel)
	s.TimeStamp = time.Now().UnixNano()
	return &s
}
