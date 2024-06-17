package user_status

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

// ap can be negative
// return an error when the ap is exceeded (i.e client send a recover request right before their ap recover to full)
func AddUserAp(session *userdata.Session, ap int32) *response.RecoverableExceptionResponse {
	if session.UserStatus.ActivityPointResetAt <= session.Time.Unix() {
		session.UserStatus.ActivityPointCount = session.Gamedata.ConstantInt[enum.ConstantIntActivityPointMaxCount]
		session.UserStatus.ActivityPointResetAt = utils.BeginOfNextHalfDay(session.Time).Unix()
	}

	ap += session.UserStatus.ActivityPointCount
	if ap > session.Gamedata.ConstantInt[enum.ConstantIntActivityPointMaxCount] {
		return &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeApLimitExceeded,
		}
	}
	session.UserStatus.ActivityPointCount = ap
	return nil
}
