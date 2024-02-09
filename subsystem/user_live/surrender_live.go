package user_live

import (
	"elichika/generic"
	"elichika/subsystem/user_status"
	"elichika/userdata"
	"elichika/utils"
)

func SurrenderLive(session *userdata.Session) generic.Nullable[int32] {
	exist, _, startReq := LoadUserLive(session)
	utils.MustExist(exist)
	ClearUserLive(session)
	// remove only half the LP
	lpCost := session.Gamedata.LiveDifficulty[startReq.LiveDifficultyId].ConsumedLP / 2
	user_status.AddUserLp(session, -lpCost)
	return generic.NewNullable(lpCost)
}
