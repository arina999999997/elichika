package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGuildRankingRewardStep struct {
	MemberMasterId int32          `xorm:"'member_master_id'"`
	UpperRank      int32          `xorm:"'upper_rank'"`
	LowerRank      *int32         `xorm:"'lower_rank'"`
	RankGrade      int32          `xorm:"'rank_grade'"`
	Content        client.Content `xorm:"extends"`
}

type MemberGuildRankingReward struct {
	Steps           []MemberGuildRankingRewardStep
	StepCount       int32
	RankNumberLimit int32
}

func (m *MemberGuildRankingReward) GetRewardContent(rank int32) *client.Content {
	var i int32
	if rank == 0 {
		rank = m.Steps[m.StepCount-1].UpperRank
	}
	for i = m.StepCount - 1; m.Steps[i].UpperRank > rank; i-- {
		continue
	}
	return &m.Steps[i].Content
}

func loadMemberGuildRankingReward(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberGuildRankingReward")

	gamedata.MemberGuildRankingRewardInside = make(map[int32]*MemberGuildRankingReward)
	for memberId := range gamedata.Member {
		gamedata.MemberGuildRankingRewardInside[memberId] = &MemberGuildRankingReward{}
		err := masterdata_db.Table("m_member_guild_ranking_reward_inside").Where("member_master_id = ?", memberId).
			OrderBy("upper_rank").Find(&gamedata.MemberGuildRankingRewardInside[memberId].Steps)
		utils.CheckErr(err)
		m := gamedata.MemberGuildRankingRewardInside[memberId]
		m.StepCount = int32(len(m.Steps))
		m.RankNumberLimit = m.Steps[m.StepCount-1].UpperRank - 1
	}

	gamedata.MemberGuildRankingRewardOutside = make(map[int32]*MemberGuildRankingReward)
	for memberId := range gamedata.Member {
		gamedata.MemberGuildRankingRewardOutside[memberId] = &MemberGuildRankingReward{}
		err := masterdata_db.Table("m_member_guild_ranking_reward_outside").Where("member_master_id = ?", memberId).
			OrderBy("upper_rank").Find(&gamedata.MemberGuildRankingRewardOutside[memberId].Steps)
		utils.CheckErr(err)
		m := gamedata.MemberGuildRankingRewardOutside[memberId]
		m.StepCount = int32(len(m.Steps))
		m.RankNumberLimit = m.Steps[m.StepCount-1].UpperRank - 1
	}
}

func init() {
	addLoadFunc(loadMemberGuildRankingReward)
	addPrequisite(loadMemberGuildRankingReward, loadMember)
}
