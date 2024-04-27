package userdata

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/userdata/database"
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
	SessionTypeDirectDbWrite = 3
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

	CommandId int32

	UniqueCount int64

	SendMissionDetail bool

	AuthenticationData database.UserAuthentication

	Finalized  bool
	IsSharedDb bool
}

func (session *Session) NextUniqueId() int64 {
	result := session.Time.UnixNano() + session.UniqueCount
	session.UniqueCount++
	return result
}

// Commit changes into the db
// Calling multiple time is fine as it allow for some specific use case
// Note that calling Finalize again does nothing, after calling it once, the session can no longer be used to write or read
func (session *Session) Finalize() {
	if (session == nil) || session.Finalized {
		return
	}
	session.Finalized = true
	var err error
	if session.SessionType != SessionTypeDirectDbWrite {
		if session.SessionType == SessionTypeLogin {
			// if login then we only need to update a thing
			userStatusFinalizer(session)
			userAuthenticationDataFinalizer(session)
		} else {
			for _, finalizer := range finalizers {
				finalizer(session)
			}
		}
	}
	if !session.IsSharedDb {
		err = session.Db.Commit()
		utils.CheckErr(err)
	}
}

func (session *Session) Close() {
	// fmt.Printf("close: %p\n", session)
	if (session == nil) || session.IsSharedDb {
		return
	}
	err := recover()
	if err != nil {
		session.Db.Rollback()
		session.Db.Close()
		panic(err)
	} else {
		session.Db.Close()
	}
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
	AddFinalizer(userStatusFinalizer)
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
	return GetSessionWithSharedDb(ctx, userId, nil)
}

func GetSessionWithSharedDb(ctx *gin.Context, userId int32, otherSession *Session) *Session {
	s := Session{}
	s.Time = time.Now()
	s.Ctx = ctx
	s.UserId = userId
	{
		g, exist := ctx.Get("gamedata")
		if exist {
			s.Gamedata = g.(*gamedata.Gamedata)
		}
	}
	defer func() {
		err := recover()
		if err != nil {
			s.Db.Close()
			panic(err)
		}
	}()
	if otherSession != nil {
		s.Db = otherSession.Db
		s.IsSharedDb = true
	} else {
		s.Db = Engine.NewSession()
		err := s.Db.Begin()
		utils.CheckErr(err)
	}

	exist, err := s.Db.Table("u_status").Where("user_id = ?", userId).Get(&s.UserModel.UserStatus)
	utils.CheckErr(err)
	if !exist {
		s.Close()
		return nil
	}
	s.fetchAuthenticationData()
	s.UserStatus = &s.UserModel.UserStatus
	s.UserContentDiffs = make(map[int32](map[int32]client.Content))

	s.MemberLovePanelDiffs = make(map[int32]client.MemberLovePanel)
	return &s
}
