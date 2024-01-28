package user_member

import (
	"elichika/userdata"
)

func IncreaseMemberLoveLevelLimit(session *userdata.Session, memberMasterId, increasedLoveLevel int32) (int32, int32) {
	member := GetMember(session, memberMasterId)
	beforeLoveLevelLimit := session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit)
	afterLoveLevelLimit := beforeLoveLevelLimit + increasedLoveLevel
	if afterLoveLevelLimit > session.Gamedata.MemberLoveLevelCount {
		afterLoveLevelLimit = session.Gamedata.MemberLoveLevelCount
	}
	member.LovePointLimit = session.Gamedata.MemberLoveLevelLovePoint[afterLoveLevelLimit]
	UpdateMember(session, member)
	return beforeLoveLevelLimit, afterLoveLevelLimit
}
