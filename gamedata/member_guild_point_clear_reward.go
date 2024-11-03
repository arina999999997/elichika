package gamedata

import (
	"elichika/client"

	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGuildPointClearReward struct {
	MemberMasterId int32          `xorm:"pk 'member_master_id'"`
	TargetPoint    int32          `xorm:"'target_point'"`
	Content        client.Content `xorm:"extends"`
}

func loadMemberGuildPointClearReward(gamedata *Gamedata) {
	fmt.Println("Loading MemberGuildPointClearReward")
	gamedata.MemberGuildPointClearReward = make(map[int32]*MemberGuildPointClearReward)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_guild_point_clear_reward").Find(&gamedata.MemberGuildPointClearReward)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadMemberGuildPointClearReward)
}
