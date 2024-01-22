package profile

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func FetchProfile(session *userdata.Session, otherUserId int32) response.UserProfileResponse {
	resp := response.UserProfileResponse{
		ProfileInfo: session.GetOtherUserProfileInformation(otherUserId),
	}

	// need to return this in order, so make the array then write to it
	// this also prevent the game from freezing if 2 cards or more have the same bit
	// TODO(refactor, guest): Rewrite this once we get there
	for i := int32(1); i <= 7; i++ {
		resp.GuestInfo.LivePartnerCards.Append(client.ProfileLivePartnerCard{
			LivePartnerCategoryMasterId: i,
		})
	}

	partnerCards := userdata.FetchPartnerCards(otherUserId)
	for _, card := range partnerCards {
		partnerCard := session.GetOtherUserCard(otherUserId, card.CardMasterId)
		for i := 1; i <= 7; i++ {
			if (card.LivePartnerCategories & (1 << i)) != 0 {
				resp.GuestInfo.LivePartnerCards.Slice[i-1].PartnerCard = partnerCard
			}
		}
	}

	// live clear stats
	liveStats := session.GetOtherUserLiveStats(otherUserId)
	for i, liveDifficultyType := range enum.LiveDifficultyTypes {
		resp.PlayInfo.LivePlayCount.Set(int32(liveDifficultyType), int32(liveStats.LivePlayCount[i]))
		resp.PlayInfo.LiveClearCount.Set(int32(liveDifficultyType), int32(liveStats.LiveClearCount[i]))
	}

	session.Db.Table("u_card").Where("user_id = ?", otherUserId).
		OrderBy("live_join_count DESC").Limit(3).Find(&resp.PlayInfo.JoinedLiveCardRanking.Slice)
	session.Db.Table("u_card").Where("user_id = ?", otherUserId).
		OrderBy("active_skill_play_count DESC").Limit(3).Find(&resp.PlayInfo.PlaySkillCardRanking.Slice)

	// custom profile
	customProfile := session.GetOtherUserSetProfile(otherUserId)
	if customProfile.VoltageLiveDifficultyId != 0 {
		resp.PlayInfo.MaxScoreLiveDifficulty.LiveDifficultyMasterId = generic.NewNullable(customProfile.VoltageLiveDifficultyId)
		resp.PlayInfo.MaxScoreLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(otherUserId, customProfile.VoltageLiveDifficultyId).MaxScore
	}
	if customProfile.CommboLiveDifficultyId != 0 {
		resp.PlayInfo.MaxComboLiveDifficulty.LiveDifficultyMasterId = generic.NewNullable(customProfile.CommboLiveDifficultyId)
		resp.PlayInfo.MaxComboLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(otherUserId, customProfile.CommboLiveDifficultyId).MaxCombo
	}

	cards := []client.UserCard{}
	err := session.Db.Table("u_card").Where("user_id = ?", otherUserId).Find(&cards)
	utils.CheckErr(err)

	ownedCardCount := map[int32]int32{}
	allTrainingActivatedCardCount := map[int32]int32{}
	for _, card := range cards {
		memberId := session.Gamedata.Card[card.CardMasterId].Member.Id
		ownedCardCount[memberId]++
		if card.IsAllTrainingActivated {
			allTrainingActivatedCardCount[memberId]++
		}
	}

	members := []client.UserMember{}
	err = session.Db.Table("u_member").Where("user_id = ?", otherUserId).OrderBy("member_master_id").Find(&members)
	utils.CheckErr(err)

	for _, member := range members {
		resp.MemberInfo.UserMembers.Append(client.ProfileUserMember{
			MemberMasterId:                member.MemberMasterId,
			LoveLevel:                     member.LoveLevel,
			LovePointLimit:                member.LovePointLimit,
			OwnedCardCount:                ownedCardCount[member.MemberMasterId],
			AllTrainingActivatedCardCount: allTrainingActivatedCardCount[member.MemberMasterId],
		})
	}

	utils.CheckErr(err)
	return resp
}
