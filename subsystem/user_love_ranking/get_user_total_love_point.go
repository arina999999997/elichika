package user_love_ranking

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

var (
	memberConditions = map[int32]string{}
)

func GetUserTotalLovePoint(session *userdata.Session, otherUserId, condition int32) int32 {
	totalLove, err := session.Db.Table("u_member").Where("user_id = ?"+memberConditions[condition], otherUserId).
		SumInt(&client.UserMember{}, "love_point")
	utils.CheckErr(err)
	return int32(totalLove)
}

func init() {
	// TODO(extra): this assume there's no other member outside of these
	memberConditions[enum.LoveRankingConditionTypeAll] = ""
	memberConditions[enum.LoveRankingConditionTypeMuse] = " AND member_master_id <= 9"
	memberConditions[enum.LoveRankingConditionTypeAqours] = " AND member_master_id >= 101 AND member_master_id <= 109"
	memberConditions[enum.LoveRankingConditionTypeNiji] = " AND member_master_id >= 201"
	memberConditions[enum.LoveRankingConditionTypeMember1] = " AND member_master_id = 1"
	memberConditions[enum.LoveRankingConditionTypeMember2] = " AND member_master_id = 2"
	memberConditions[enum.LoveRankingConditionTypeMember3] = " AND member_master_id = 3"
	memberConditions[enum.LoveRankingConditionTypeMember4] = " AND member_master_id = 4"
	memberConditions[enum.LoveRankingConditionTypeMember5] = " AND member_master_id = 5"
	memberConditions[enum.LoveRankingConditionTypeMember6] = " AND member_master_id = 6"
	memberConditions[enum.LoveRankingConditionTypeMember7] = " AND member_master_id = 7"
	memberConditions[enum.LoveRankingConditionTypeMember8] = " AND member_master_id = 8"
	memberConditions[enum.LoveRankingConditionTypeMember9] = " AND member_master_id = 9"
	memberConditions[enum.LoveRankingConditionTypeMember101] = " AND member_master_id = 101"
	memberConditions[enum.LoveRankingConditionTypeMember102] = " AND member_master_id = 102"
	memberConditions[enum.LoveRankingConditionTypeMember103] = " AND member_master_id = 103"
	memberConditions[enum.LoveRankingConditionTypeMember104] = " AND member_master_id = 104"
	memberConditions[enum.LoveRankingConditionTypeMember105] = " AND member_master_id = 105"
	memberConditions[enum.LoveRankingConditionTypeMember106] = " AND member_master_id = 106"
	memberConditions[enum.LoveRankingConditionTypeMember107] = " AND member_master_id = 107"
	memberConditions[enum.LoveRankingConditionTypeMember108] = " AND member_master_id = 108"
	memberConditions[enum.LoveRankingConditionTypeMember109] = " AND member_master_id = 109"
	memberConditions[enum.LoveRankingConditionTypeMember201] = " AND member_master_id = 201"
	memberConditions[enum.LoveRankingConditionTypeMember202] = " AND member_master_id = 202"
	memberConditions[enum.LoveRankingConditionTypeMember203] = " AND member_master_id = 203"
	memberConditions[enum.LoveRankingConditionTypeMember204] = " AND member_master_id = 204"
	memberConditions[enum.LoveRankingConditionTypeMember205] = " AND member_master_id = 205"
	memberConditions[enum.LoveRankingConditionTypeMember206] = " AND member_master_id = 206"
	memberConditions[enum.LoveRankingConditionTypeMember207] = " AND member_master_id = 207"
	memberConditions[enum.LoveRankingConditionTypeMember208] = " AND member_master_id = 208"
	memberConditions[enum.LoveRankingConditionTypeMember209] = " AND member_master_id = 209"
	memberConditions[enum.LoveRankingConditionTypeMember210] = " AND member_master_id = 210"
	memberConditions[enum.LoveRankingConditionTypeMember211] = " AND member_master_id = 211"
	memberConditions[enum.LoveRankingConditionTypeMember212] = " AND member_master_id = 212"
	memberConditions[enum.LoveRankingConditionTypePrintemps] = " AND member_master_id IN (1, 3, 8)"
	memberConditions[enum.LoveRankingConditionTypeBibi] = " AND member_master_id IN (2, 6, 9)"
	memberConditions[enum.LoveRankingConditionTypeLilyWhite] = " AND member_master_id IN (4, 5, 7)"
	memberConditions[enum.LoveRankingConditionTypeCyaron] = " AND member_master_id IN (101, 105, 109)"
	memberConditions[enum.LoveRankingConditionTypeAzalea] = " AND member_master_id IN (103, 104, 107)"
	memberConditions[enum.LoveRankingConditionTypeGuiltyKiss] = " AND member_master_id IN (102, 106, 108)"
	memberConditions[enum.LoveRankingConditionTypeAzuna] = " AND member_master_id IN (201, 203, 207)"
	memberConditions[enum.LoveRankingConditionTypeDiverDiva] = " AND member_master_id IN (204, 205)"
	memberConditions[enum.LoveRankingConditionTypeQu4rtz] = " AND member_master_id IN (202, 206, 208, 209)"
	memberConditions[enum.LoveRankingConditionTypeNijiUnit4] = " AND member_master_id >= 210"
}
