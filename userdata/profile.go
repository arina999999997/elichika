package userdata

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
)

func FetchDBProfile(userId int32, result interface{}) {
	exist, err := Engine.Table("u_info").Where("user_id = ?", userId).Get(result)
	utils.CheckErrMustExist(err, exist)
}

func FetchPartnerCards(otherUserId int) []client.UserCard {
	partnerCards := []client.UserCard{}
	err := Engine.Table("u_card").
		Where("user_id = ? AND live_partner_categories != 0", otherUserId).
		Find(&partnerCards)
	if err != nil {
		panic(err)
	}
	return partnerCards
}

func (session *Session) GetPartnerCardFromUserCard(card client.UserCard) model.PartnerCardInfo {

	memberId := session.Gamedata.Card[card.CardMasterId].Member.Id

	partnerCard := model.PartnerCardInfo{}

	jsonByte, err := json.Marshal(card)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonByte, &partnerCard)
	if err != nil {
		panic(err)
	}

	// TODO(friend): This doesn't work totally correctly because the card is isn't totally attached to the user anymore
	exist, err := Engine.Table("u_member").Where("user_id = ? AND member_master_id = ?", session.UserId, memberId).
		Cols("love_level").Get(&partnerCard.LoveLevel)
	utils.CheckErrMustExist(err, exist)

	partnerCard.PassiveSkillLevels = []int{}
	partnerCard.PassiveSkillLevels = append(partnerCard.PassiveSkillLevels, int(card.PassiveSkillALevel))
	partnerCard.PassiveSkillLevels = append(partnerCard.PassiveSkillLevels, int(card.PassiveSkillBLevel))
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, int(card.AdditionalPassiveSkill1Id))
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, int(card.AdditionalPassiveSkill2Id))
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, int(card.AdditionalPassiveSkill3Id))
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, int(card.AdditionalPassiveSkill4Id))
	partnerCard.MemberLovePanels = []int{} // must not be null

	// filling this for a card of self freeze the game
	// the displayed value still correct for own's card in the guest setup menu with an empty array
	// display value in getOtherUserCard is wrong, but if we fill in for own card then it also freeze
	// TODO: revisit after implmenting friends

	// lovePanel := client.MemberLovePanel{}
	// exist, err = Engine.Table("u_member").
	// 	Where("user_id = ? AND member_master_id = ?", card.UserId, memberId).Get(&lovePanel)
	// if err != nil {
	// 	panic(err)
	// }
	// if !exist {
	// 	panic("member doesn't exist")
	// }
	// partnerCard.MemberLovePanels = lovePanel.MemberLovePanelCellIds

	return partnerCard
}

func (session *Session) GetOtherUserProfileInformation(otherUserId int32) client.ProfileInfomation {
	// TODO(friend): Actually calculate the friend links
	otherUserStatus := client.UserStatus{}
	exist, err := Engine.Table("u_info").Where("user_id = ?", otherUserId).Get(&otherUserStatus)
	utils.CheckErrMustExist(err, exist)
	otherUserCard := session.GetOtherUserCard(otherUserId, otherUserStatus.RecommendCardMasterId)
	otherUser := client.OtherUser{
		UserId:                              otherUserId,
		Name:                                otherUserStatus.Name,
		Rank:                                otherUserStatus.Rank,
		LastPlayedAt:                        otherUserStatus.LastLoginAt, // this might not be correct
		RecommendCardMasterId:               otherUserStatus.RecommendCardMasterId,
		RecommendCardLevel:                  otherUserCard.Level,
		IsRecommendCardImageAwaken:          otherUserCard.IsAwakeningImage,
		IsRecommendCardAllTrainingActivated: otherUserCard.IsAllTrainingActivated,
		EmblemId:                            otherUserStatus.EmblemId,
		// IsNew: otherUserStatus.IsNew,
		IntroductionMessage: otherUserStatus.Message,
		// FriendApprovedAt: otherUserStatus.FriendApprovedAt,
		// RequestStatus: otherUserStatus.RequestStatus,
		// IsRequestPending: otherUserStatus.IsRequestPending,
	}
	profileInfomation := client.ProfileInfomation{
		BasicInfo:                 otherUser,
		MemberGuildMemberMasterId: otherUserStatus.MemberGuildMemberMasterId,
	}
	err = session.Db.Table("u_member").Where("user_id = ?", otherUserId).OrderBy("member_master_id").Find(&profileInfomation.LoveMembers.Slice)
	utils.CheckErr(err)
	for _, member := range profileInfomation.LoveMembers.Slice {
		profileInfomation.TotalLovePoint += member.LovePoint
	}
	return profileInfomation
}

func GetOtherUserCard(otherUserId, cardMasterId int32) client.UserCard {
	card := client.UserCard{}
	exist, err := Engine.Table("u_card").Where("user_id = ? AND card_master_id = ?", otherUserId, cardMasterId).
		Get(&card)
	utils.CheckErrMustExist(err, exist)
	return card
}

func (session *Session) GetOtherUserBasicProfile(otherUserId int32) model.UserBasicInfo {
	basicInfo := model.UserBasicInfo{}
	FetchDBProfile(otherUserId, &basicInfo)
	recommendCard := GetOtherUserCard(otherUserId, basicInfo.RecommendCardMasterId)

	basicInfo.RecommendCardLevel = int(recommendCard.Level)
	basicInfo.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	basicInfo.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated

	// friend system, not implemented
	basicInfo.FriendApprovedAt = new(int64)
	*basicInfo.FriendApprovedAt = 0
	basicInfo.RequestStatus = 3
	basicInfo.IsRequestPending = false
	return basicInfo
}

func (session *Session) GetOtherUserLiveStats(otherUserId int) model.UserProfileLiveStats {
	stats := model.UserProfileLiveStats{}
	_, err := Engine.Table("u_info").Where("user_id = ?", otherUserId).Get(&stats)
	utils.CheckErr(err)
	return stats
}

func (session *Session) GetUserLiveStats() model.UserProfileLiveStats {
	return session.GetOtherUserLiveStats(session.UserId)
}

func (session *Session) UpdateUserLiveStats(stats model.UserProfileLiveStats) {
	_, err := session.Db.Table("u_info").Where("user_id = ?", session.UserId).AllCols().Update(&stats)
	utils.CheckErr(err)
}

// fetch profile of another user, from session.UserId's perspective
// it's possible that otherUserId == session.UserId

func (session *Session) FetchProfile(otherUserId int32) response.UserProfileResponse {
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

	partnerCards := FetchPartnerCards(int(otherUserId))
	for _, card := range partnerCards {
		partnerCard := session.GetOtherUserCard(otherUserId, card.CardMasterId)
		for i := 1; i <= 7; i++ {
			if (card.LivePartnerCategories & (1 << i)) != 0 {
				resp.GuestInfo.LivePartnerCards.Slice[i-1].PartnerCard = partnerCard
			}
		}
	}

	// live clear stats
	liveStats := session.GetOtherUserLiveStats(int(otherUserId))
	for i, liveDifficultyType := range enum.LiveDifficultyTypes {
		resp.PlayInfo.LivePlayCount.Set(int32(liveDifficultyType), int32(liveStats.LivePlayCount[i]))
		resp.PlayInfo.LiveClearCount.Set(int32(liveDifficultyType), int32(liveStats.LiveClearCount[i]))
	}

	session.Db.Table("u_card").Where("user_id = ?", otherUserId).
		OrderBy("live_join_count DESC").Limit(3).Find(&resp.PlayInfo.JoinedLiveCardRanking.Slice)
	session.Db.Table("u_card").Where("user_id = ?", otherUserId).
		OrderBy("active_skill_play_count DESC").Limit(3).Find(&resp.PlayInfo.PlaySkillCardRanking.Slice)

	// custom profile
	customProfile := session.GetOtherUserSetProfile(int(otherUserId))
	if customProfile.VoltageLiveDifficultyId != 0 {
		resp.PlayInfo.MaxScoreLiveDifficulty.LiveDifficultyMasterId = generic.NewNullable(customProfile.VoltageLiveDifficultyId)
		resp.PlayInfo.MaxScoreLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(int(otherUserId), customProfile.VoltageLiveDifficultyId).MaxScore
	}
	if customProfile.CommboLiveDifficultyId != 0 {
		resp.PlayInfo.MaxComboLiveDifficulty.LiveDifficultyMasterId = generic.NewNullable(customProfile.CommboLiveDifficultyId)
		resp.PlayInfo.MaxComboLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(int(otherUserId), customProfile.CommboLiveDifficultyId).MaxCombo
	}

	// can get from members[] to save sql
	session.Db.Table("u_member").Where("user_id = ?", otherUserId).Find(&resp.MemberInfo.UserMembers.Slice)
	return resp
}

func (session *Session) GetOtherUserSetProfile(otherUserId int) client.UserSetProfile {
	p := client.UserSetProfile{}
	_, err := session.Db.Table("u_set_profile").Where("user_id = ?", otherUserId).Get(&p)
	utils.CheckErr(err)
	return p
}

func (session *Session) GetUserSetProfile() client.UserSetProfile {
	return session.GetOtherUserSetProfile(session.UserId)
}

// doesn't need to return delta patch or submit at the start because we would need to fetch profile everytime we need this thing
func (session *Session) SetUserSetProfile(userSetProfile client.UserSetProfile) {
	affected, err := session.Db.Table("u_set_profile").Where("user_id = ?", session.UserId).
		AllCols().Update(&userSetProfile)
	utils.CheckErr(err)
	if affected == 0 {
		// need to insert
		genericDatabaseInsert(session, "u_set_profile", userSetProfile)
	}
}

func userSetProfileFinalizer(session *Session) {
	for _, userSetProfile := range session.UserModel.UserSetProfileById.Map {
		affected, err := session.Db.Table("u_set_profile").Where("user_id = ?",
			session.UserId).AllCols().Update(*userSetProfile)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_set_profile", *userSetProfile)
		}
	}
}

func init() {
	addFinalizer(userSetProfileFinalizer)
}
