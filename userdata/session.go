package userdata

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/utils"

	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

// A session is a complete transation between server and client
// so 1 session per request
// A session fetch the data needs to be modified, and sometime modify the data if the code is shared between handlers.
// session can use ctx to get things like user id / master db, but it should not make any network operation
const (
	SessionTypeGameplay      = 0
	SessionTypeLogin         = 1
	SessionTypeImportAccount = 2
)

type Session struct {
	Time       time.Time
	Db         *xorm.Session
	Ctx        *gin.Context
	UserId     int32
	Gamedata   *gamedata.Gamedata
	UserStatus *client.UserStatus // link to UserModel.UserStatus
	// TODO: change the map to index map?
	MemberLovePanelDiffs map[int32]client.MemberLovePanel
	MemberLovePanels     []client.MemberLovePanel
	UserContentDiffs     map[int32](map[int32]client.Content) // content_type then content_id

	UnreceivedContent []client.Content

	SessionType int
	UserModel   client.UserModel

	UniqueCount int64
}

func (session *Session) NextUniqueId() int64 {
	result := session.Time.UnixNano() + session.UniqueCount
	session.UniqueCount++
	return result
}

// Push update into the db and create the diff
// The actual response depend on the API, but they often contain the diff somewhere
// The mainKey is the key to the diff
func (session *Session) Finalize() {
	var err error
	if session.SessionType == SessionTypeLogin {
		// if login then we only need to update a thing
		userStatusFinalizer(session)
	} else {
		for _, finalizer := range finalizers {
			finalizer(session)
		}
		finalizeMemberLovePanelDiffs(session)
	}
	err = session.Db.Commit()
	utils.CheckErr(err)
}

func (session *Session) Close() {
	// fmt.Printf("close: %p\n", session)
	if session == nil {
		return
	}
	session.Db.Close()
}

func userStatusFinalizer(session *Session) {
	affected, err := session.Db.Table("u_status").Where("user_id = ?", session.UserId).AllCols().Update(session.UserStatus)
	utils.CheckErr(err)
	if affected != 1 {
		if session.SessionType != SessionTypeImportAccount {
			panic("user doesn't exist in u_info")
		} else {
			GenericDatabaseInsert(session, "u_status", *session.UserStatus)
		}
	}
}
func init() {
	AddContentFinalizer(userStatusFinalizer)
}

func UserExist(userId int32) bool {
	exist, err := Engine.Table("u_status").Exist(
		&generic.UserIdWrapper[client.UserStatus]{
			UserId: userId,
		})
	utils.CheckErr(err)
	return exist
}

func GetSession(ctx *gin.Context, userId int32) *Session {
	s := Session{}
	s.Time = time.Now()
	s.Ctx = ctx
	s.UserId = userId
	s.Gamedata = ctx.MustGet("gamedata").(*gamedata.Gamedata)
	s.Db = Engine.NewSession()
	err := s.Db.Begin()
	utils.CheckErr(err)

	exist, err := s.Db.Table("u_status").Where("user_id = ?", userId).Get(&s.UserModel.UserStatus)
	utils.CheckErr(err)
	if !exist {
		s.Close()
		return nil
	}
	s.UserStatus = &s.UserModel.UserStatus
	s.UserContentDiffs = make(map[int32](map[int32]client.Content))

	s.MemberLovePanelDiffs = make(map[int32]client.MemberLovePanel)
	return &s
}

func SessionFromImportedLoginData(ctx *gin.Context, loginData *response.LoginResponse, userId int32) *Session {
	s := Session{}
	s.Time = time.Now()
	s.SessionType = SessionTypeImportAccount
	s.Ctx = ctx
	s.UserId = userId
	s.Gamedata = ctx.MustGet("gamedata").(*gamedata.Gamedata)
	s.Db = Engine.NewSession()
	err := s.Db.Begin()
	utils.CheckErr(err)
	s.UserModel = *loginData.UserModel
	s.UserStatus = &s.UserModel.UserStatus

	s.UserContentDiffs = make(map[int32](map[int32]client.Content))

	s.MemberLovePanels = loginData.MemberLovePanels.Slice
	GenericDatabaseInsert(&s, "u_login", *loginData)
	utils.CheckErr(err)
	return &s
}
