package user_love_ranking

import (
	"elichika/enum"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

// for simplicity we will generate a total love point table from the member table
// for now this table is not delta patched on each request, but delta patched once in a while
var lastGeneratedAt int64 = 0

func generateTotalLovePointTable(session *userdata.Session) {
	if session.Time.Unix()-lastGeneratedAt <= int64(enum.MinuteSecondCount) {
		return
	}
	userIds := []int32{}
	err := session.Db.Table("u_status").Cols("user_id").Find(&userIds)
	utils.CheckErr(err)
	for _, userId := range userIds {
		userTotalLovePointSummary := database.UserTotalLovePointSummary{
			UserId:     userId,
			All:        GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeAll),
			Muse:       GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeMuse),
			Aqours:     GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeAqours),
			Niji:       GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeNiji),
			Printemps:  GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypePrintemps),
			Bibi:       GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeBibi),
			LilyWhite:  GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeLilyWhite),
			Cyaron:     GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeCyaron),
			Azalea:     GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeAzalea),
			GuiltyKiss: GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeGuiltyKiss),
			Azuna:      GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeAzuna),
			DiverDiva:  GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeDiverDiva),
			Qu4rtz:     GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeQu4rtz),
			NijiUnit4:  GetUserTotalLovePoint(session, userId, enum.LoveRankingConditionTypeNijiUnit4),
		}
		affected, err := session.Db.Table("u_total_love_point_summary").
			Where("user_id = ?", userId).Update(&userTotalLovePointSummary)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_total_love_point_summary").Insert(&userTotalLovePointSummary)
			utils.CheckErr(err)
		}
	}
	lastGeneratedAt = session.Time.Unix()
}
