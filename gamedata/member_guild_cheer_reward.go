package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/generic/drop"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

func loadMemberGuildCheerReward(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberGuildCheerReward")
	type MemberGuildRewardLot struct {
		// Id
		MemberMasterId int32
		Content        client.Content `xorm:"extends"`
		// DisplayOrder
	}
	gamedata.MemberGuildCheerReward = map[int32]*drop.DropList[client.Content]{}
	rewards := []MemberGuildRewardLot{}
	err := masterdata_db.Table("m_member_guild_reward_lot").Find(&rewards)
	utils.CheckErr(err)
	for _, reward := range rewards {
		dropList, exist := gamedata.MemberGuildCheerReward[reward.MemberMasterId]
		if !exist {
			gamedata.MemberGuildCheerReward[reward.MemberMasterId] = new(drop.DropList[client.Content])
			dropList = gamedata.MemberGuildCheerReward[reward.MemberMasterId]
		}
		dropList.AddItem(reward.Content)
	}
}

func init() {
	addLoadFunc(loadMemberGuildCheerReward)
}
