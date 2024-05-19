package user_love_ranking

import (
	"elichika/enum"
)

func getMemberIdFromCondition(condtiionType int32) int32 {
	if (condtiionType < enum.LoveRankingConditionTypeMember1) || (condtiionType > enum.LoveRankingConditionTypeMember212) {
		panic("must be a single member condition")
	}
	if condtiionType <= enum.LoveRankingConditionTypeMember9 {
		return condtiionType - enum.LoveRankingConditionTypeMember1 + 1
	} else if condtiionType <= enum.LoveRankingConditionTypeMember109 {
		return condtiionType - enum.LoveRankingConditionTypeMember101 + 101
	} else {
		return condtiionType - enum.LoveRankingConditionTypeMember201 + 201
	}
}
