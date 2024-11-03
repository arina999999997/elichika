package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGuildConstant struct {
	StartAt                         int64 `xorm:"'start_at'"`
	LovePointCalculationNum         int32 `xorm:"'love_point_calculation_num'"`
	VoltageCalculationNum           int32 `xorm:"'voltage_calculation_num'"`
	JoinConditionPoint              int32 `xorm:"'join_condition_point'"`
	JoinConditionRank               int32 `xorm:"'join_condition_rank'"`
	InsideRankingTopRangeLowerLimit int32 `xorm:"'inside_ranking_top_range_lower_limit'"`
	DailyLimitPoint                 int32 `xorm:"'daily_limit_point'"`
	SupportPoint                    int32 `xorm:"'support_point'"`
	BackgroundMasterId              int32 `xorm:"'background_master_id'"`
	RankingViewBoder                int32 `xorm:"'ranking_view_border'"`
}

func loadMemberGuildConstant(gamedata *Gamedata) {
	fmt.Println("Loading MemberGuildConstant")
	var exist bool
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		exist, err = session.Table("m_member_guild_constant").OrderBy("start_at DESC").Limit(1).Get(&gamedata.MemberGuildConstant)
	})
	utils.CheckErrMustExist(err, exist)
}

func init() {
	addLoadFunc(loadMemberGuildConstant)
}
