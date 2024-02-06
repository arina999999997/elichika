package user_profile

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_live_difficulty"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserProfilePlayHistory(session *userdata.Session, otherUserId int32) client.ProfilePlayHistory {
	// build from liveDifficulty
	res := client.ProfilePlayHistory{}

	// clear stats
	userLiveDifficulties := []client.UserLiveDifficulty{}
	err := session.Db.Table("u_live_difficulty").Where("user_id = ?", otherUserId).Find(&userLiveDifficulties)
	utils.CheckErr(err)
	playCount := map[int32]int32{}
	clearCount := map[int32]int32{}
	for _, liveDifficulty := range userLiveDifficulties {
		masterLiveDifficulty := session.Gamedata.LiveDifficulty[liveDifficulty.LiveDifficultyId]
		if masterLiveDifficulty == nil {
			continue // song doesn't exist in the server, but exist in records
		}
		if !masterLiveDifficulty.IsCountTarget {
			continue
		}
		// the naming is a bit inconsistent, but playCount in profileHistory context mean time actually cleared or skipped
		// is cleared
		if liveDifficulty.ClearCount > 0 {
			playCount[masterLiveDifficulty.LiveDifficultyType] += liveDifficulty.ClearCount
			clearCount[masterLiveDifficulty.LiveDifficultyType]++
		}
	}
	types := []int32{enum.LiveDifficultyTypeNormal, enum.LiveDifficultyTypeHard, enum.LiveDifficultyTypeExpert,
		enum.LiveDifficultyTypeExpertPlus, enum.LiveDifficultyTypeExpertPlusPlus}
	for _, t := range types {
		res.LivePlayCount.Set(t, playCount[t])
		res.LiveClearCount.Set(t, clearCount[t])
	}

	// card stats
	cardMasterIds := []int32{}
	err = session.Db.Table("u_card_play_count_stat").Where("user_id = ?", otherUserId).
		OrderBy("live_join_count DESC").Limit(3).Cols("card_master_id").Find(&cardMasterIds)
	utils.CheckErr(err)
	for _, cardMasterId := range cardMasterIds {
		res.JoinedLiveCardRanking.Append(user_card.GetOtherUserProfileUserCard(session, otherUserId, cardMasterId))
	}
	cardMasterIds = []int32{}

	err = session.Db.Table("u_card_play_count_stat").Where("user_id = ?", otherUserId).
		OrderBy("active_skill_play_count DESC").Limit(3).Cols("card_master_id").Find(&cardMasterIds)
	utils.CheckErr(err)
	for _, cardMasterId := range cardMasterIds {
		res.PlaySkillCardRanking.Append(user_card.GetOtherUserProfileUserCard(session, otherUserId, cardMasterId))
	}

	// custom profile
	customProfile := session.GetOtherUserSetProfile(otherUserId)
	if customProfile.VoltageLiveDifficultyId != 0 {
		res.MaxScoreLiveDifficulty.LiveDifficultyMasterId = generic.NewNullable(customProfile.VoltageLiveDifficultyId)
		res.MaxScoreLiveDifficulty.Score =
			user_live_difficulty.GetOtherUserLiveDifficulty(session, otherUserId, customProfile.VoltageLiveDifficultyId).MaxScore
	}
	if customProfile.CommboLiveDifficultyId != 0 {
		res.MaxComboLiveDifficulty.LiveDifficultyMasterId = generic.NewNullable(customProfile.CommboLiveDifficultyId)
		res.MaxComboLiveDifficulty.Score =
			user_live_difficulty.GetOtherUserLiveDifficulty(session, otherUserId, customProfile.CommboLiveDifficultyId).MaxCombo
	}
	return res
}
