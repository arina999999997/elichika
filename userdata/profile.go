package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
)

func FetchDBProfile(userId int, result interface{}) {
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

	// lovePanel := model.UserMemberLovePanel{}
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

func GetOtherUserCard(otherUserId int, cardMasterId int32) client.UserCard {
	card := client.UserCard{}
	exist, err := Engine.Table("u_card").Where("user_id = ? AND card_master_id = ?", otherUserId, cardMasterId).
		Get(&card)
	utils.CheckErrMustExist(err, exist)
	return card
}

func (session *Session) GetOtherUserBasicProfile(otherUserId int) model.UserBasicInfo {
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
func (session *Session) FetchProfile(otherUserId int) model.Profile {
	profile := model.Profile{}

	exist, err := session.Db.Table("u_info").Where("user_id = ?", otherUserId).Get(&profile)
	utils.CheckErrMustExist(err, exist)

	// recommend card
	recommendCard := GetOtherUserCard(otherUserId, profile.ProfileInfo.BasicInfo.RecommendCardMasterId)

	profile.ProfileInfo.BasicInfo.RecommendCardLevel = int(recommendCard.Level)
	profile.ProfileInfo.BasicInfo.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	profile.ProfileInfo.BasicInfo.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated

	// friend system, not implemented
	profile.ProfileInfo.BasicInfo.FriendApprovedAt = new(int64)
	*profile.ProfileInfo.BasicInfo.FriendApprovedAt = 0
	profile.ProfileInfo.BasicInfo.RequestStatus = 3
	profile.ProfileInfo.BasicInfo.IsRequestPending = false

	// other user's members
	members := []client.UserMember{}
	err = session.Db.Table("u_member").Where("user_id = ?", otherUserId).OrderBy("love_point DESC").Find(&members)
	utils.CheckErr(err)
	profile.ProfileInfo.TotalLovePoint = 0
	for _, member := range members {
		profile.ProfileInfo.TotalLovePoint += int(member.LovePoint)
	}
	for i := 0; i < 3; i++ {
		profile.ProfileInfo.LoveMembers[i].MemberMasterId = int(members[i].MemberMasterId)
		profile.ProfileInfo.LoveMembers[i].LovePoint = int(members[i].LovePoint)
	}

	// need to return this in order, so make the array then write to it
	// this also prevent the game from freezing if 2 cards or more have the same bit
	for i := 0; i < 7; i++ {
		profile.GuestInfo.LivePartnersCards = append(profile.GuestInfo.LivePartnersCards, model.LivePartnerCard{})
		profile.GuestInfo.LivePartnersCards[i].LivePartnerCategoryMasterId = i + 1
		profile.GuestInfo.LivePartnersCards[i].PartnerCard.MemberLovePanels = []int{}
		profile.GuestInfo.LivePartnersCards[i].PartnerCard.PassiveSkillLevels = []int{}
		profile.GuestInfo.LivePartnersCards[i].PartnerCard.AdditionalPassiveSkillIds = []int{}
	}

	partnerCards := FetchPartnerCards(otherUserId)
	for _, card := range partnerCards {
		partnerCard := session.GetPartnerCardFromUserCard(card)
		livePartner := model.LivePartnerCard{}
		livePartner.PartnerCard = partnerCard
		for i := 1; i <= 7; i++ {
			if (card.LivePartnerCategories & (1 << i)) != 0 {
				profile.GuestInfo.LivePartnersCards[i-1].PartnerCard = partnerCard
			}
		}
	}

	// live clear stats
	liveStats := session.GetOtherUserLiveStats(otherUserId)
	for i, liveDifficultyType := range enum.LiveDifficultyTypes {
		profile.PlayInfo.LivePlayCount = append(profile.PlayInfo.LivePlayCount, liveDifficultyType)
		profile.PlayInfo.LivePlayCount = append(profile.PlayInfo.LivePlayCount, liveStats.LivePlayCount[i])
		profile.PlayInfo.LiveClearCount = append(profile.PlayInfo.LiveClearCount, liveDifficultyType)
		profile.PlayInfo.LiveClearCount = append(profile.PlayInfo.LiveClearCount, liveStats.LiveClearCount[i])
	}

	session.Db.Table("u_card").Where("user_id = ?", otherUserId).
		OrderBy("live_join_count DESC").Limit(3).Find(&profile.PlayInfo.JoinedLiveCardRanking)
	session.Db.Table("u_card").Where("user_id = ?", otherUserId).
		OrderBy("active_skill_play_count DESC").Limit(3).Find(&profile.PlayInfo.PlaySkillCardRanking)

	// custom profile
	customProfile := session.GetOtherUserSetProfile(otherUserId)
	if customProfile.VoltageLiveDifficultyId != 0 {
		profile.PlayInfo.MaxScoreLiveDifficulty.LiveDifficultyMasterId = customProfile.VoltageLiveDifficultyId
		profile.PlayInfo.MaxScoreLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(otherUserId, customProfile.VoltageLiveDifficultyId).MaxScore
	}
	if customProfile.CommboLiveDifficultyId != 0 {
		profile.PlayInfo.MaxComboLiveDifficulty.LiveDifficultyMasterId = customProfile.CommboLiveDifficultyId
		profile.PlayInfo.MaxComboLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(otherUserId, customProfile.CommboLiveDifficultyId).MaxCombo
	}

	// can get from members[] to save sql
	session.Db.Table("u_member").Where("user_id = ?", otherUserId).Find(&profile.MemberInfo.UserMembers)
	return profile
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
	addGenericTableFieldPopulator("u_set_profile", "UserSetProfileById")
}
