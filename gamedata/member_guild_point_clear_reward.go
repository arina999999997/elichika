package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGuildPointClearReward struct {
	MemberMasterId int32          `xorm:"pk 'member_master_id'"`
	TargetPoint    int32          `xorm:"'target_point'"`
	Content        client.Content `xorm:"extends"`
}

func loadMemberGuildPointClearReward(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberGuildPointClearReward")
	gamedata.MemberGuildPointClearReward = make(map[int32]*MemberGuildPointClearReward)
	err := masterdata_db.Table("m_member_guild_point_clear_reward").Find(&gamedata.MemberGuildPointClearReward)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadMemberGuildPointClearReward)
}
