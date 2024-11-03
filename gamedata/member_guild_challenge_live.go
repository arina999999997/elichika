package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGuildChallengeLive struct {
	LiveMasterIds []int32
	Count         int32
}

func (m *MemberGuildChallengeLive) GetLiveId(memberGuildId int32) int32 {
	return m.LiveMasterIds[(memberGuildId-1)%m.Count]
}

func loadMemberGuildChallengeLive(gamedata *Gamedata) {
	fmt.Println("Loading MemberGuildChallengeLive")

	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_guild_challenge_live").OrderBy("order_num").Cols("live_master_id").Find(&gamedata.MemberGuildChallengeLive.LiveMasterIds)
	})
	utils.CheckErr(err)
	gamedata.MemberGuildChallengeLive.Count = int32(len(gamedata.MemberGuildChallengeLive.LiveMasterIds))
}

func init() {
	addLoadFunc(loadMemberGuildChallengeLive)
}
