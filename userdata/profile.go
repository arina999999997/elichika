package userdata

import (
	"elichika/enum"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
)

func FetchDBProfile(userID int, result interface{}) {
	exists, err := Engine.Table("u_info").Where("user_id = ?", userID).Get(result)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("user doesn't exist")
	}
}

func FetchPartnerCards(otherUserID int) []model.UserCard {
	partnerCards := []model.UserCard{}
	err := Engine.Table("u_card").
		Where("user_id = ? AND live_partner_categories != 0", otherUserID).
		Find(&partnerCards)
	if err != nil {
		panic(err)
	}
	return partnerCards
}

func (session *Session) GetPartnerCardFromUserCard(card model.UserCard) model.PartnerCardInfo {

	memberID := session.Gamedata.Card[card.CardMasterID].Member.ID

	partnerCard := model.PartnerCardInfo{}

	jsonByte, err := json.Marshal(card)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonByte, &partnerCard)
	if err != nil {
		panic(err)
	}

	exists, err := Engine.Table("u_member").Where("user_id = ? AND member_master_id = ?", card.UserID, memberID).
		Cols("love_level").Get(&partnerCard.LoveLevel)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("member doesn't exist")
	}

	partnerCard.PassiveSkillLevels = []int{}
	partnerCard.PassiveSkillLevels = append(partnerCard.PassiveSkillLevels, card.PassiveSkillALevel)
	partnerCard.PassiveSkillLevels = append(partnerCard.PassiveSkillLevels, card.PassiveSkillBLevel)
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, card.AdditionalPassiveSkill1ID)
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, card.AdditionalPassiveSkill2ID)
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, card.AdditionalPassiveSkill3ID)
	partnerCard.AdditionalPassiveSkillIds = append(partnerCard.AdditionalPassiveSkillIds, card.AdditionalPassiveSkill4ID)
	partnerCard.MemberLovePanels = []int{} // must not be null

	// filling this for a card of self freeze the game
	// the displayed value still correct for own's card in the guest setup menu with an empty array
	// display value in getOtherUserCard is wrong, but if we fill in for own card then it also freeze
	// TODO: revisit after implmenting friends

	// lovePanel := model.UserMemberLovePanel{}
	// exists, err = Engine.Table("u_member").
	// 	Where("user_id = ? AND member_master_id = ?", card.UserID, memberId).Get(&lovePanel)
	// if err != nil {
	// 	panic(err)
	// }
	// if !exists {
	// 	panic("member doesn't exist")
	// }
	// partnerCard.MemberLovePanels = lovePanel.MemberLovePanelCellIDs

	return partnerCard
}

func GetOtherUserCard(otherUserID, cardMasterID int) model.UserCard {
	card := model.UserCard{}
	exists, err := Engine.Table("u_card").Where("user_id = ? AND card_master_id = ?", otherUserID, cardMasterID).
		Get(&card)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("user card doesn't exist")
	}
	return card
}

func (session *Session) GetOtherUserBasicProfile(otherUserID int) model.UserBasicInfo {
	basicInfo := model.UserBasicInfo{}
	FetchDBProfile(otherUserID, &basicInfo)
	recommendCard := GetOtherUserCard(otherUserID, basicInfo.RecommendCardMasterID)

	basicInfo.RecommendCardLevel = recommendCard.Level
	basicInfo.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	basicInfo.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated

	// friend system, not implemented
	basicInfo.FriendApprovedAt = new(int64)
	*basicInfo.FriendApprovedAt = 0
	basicInfo.RequestStatus = 3
	basicInfo.IsRequestPending = false
	return basicInfo
}

func (session *Session) GetOtherUserLiveStats(otherUserID int) model.UserProfileLiveStats {
	stats := model.UserProfileLiveStats{}
	_, err := Engine.Table("u_info").Where("user_id = ?", otherUserID).Get(&stats)
	utils.CheckErr(err)
	return stats
}

func (session *Session) GetUserLiveStats() model.UserProfileLiveStats {
	return session.GetOtherUserLiveStats(session.UserStatus.UserID)
}

func (session *Session) UpdateUserLiveStats(stats model.UserProfileLiveStats) {
	_, err := session.Db.Table("u_info").Where("user_id = ?", session.UserStatus.UserID).AllCols().Update(&stats)
	utils.CheckErr(err)
}

// fetch profile of another user, from session.UserStatus.UserID's perspective
// it's possible that otherUserID == session.UserStatus.UserID
func (session *Session) FetchProfile(otherUserID int) model.Profile {
	profile := model.Profile{}

	exists, err := session.Db.Table("u_info").Where("user_id = ?", otherUserID).Get(&profile)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("user doesn't exist")
	}

	// recommend card
	recommendCard := GetOtherUserCard(otherUserID, profile.ProfileInfo.BasicInfo.RecommendCardMasterID)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("card doesn't exist")
	}
	profile.ProfileInfo.BasicInfo.RecommendCardLevel = recommendCard.Level
	profile.ProfileInfo.BasicInfo.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	profile.ProfileInfo.BasicInfo.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated

	// friend system, not implemented
	profile.ProfileInfo.BasicInfo.FriendApprovedAt = new(int64)
	*profile.ProfileInfo.BasicInfo.FriendApprovedAt = 0
	profile.ProfileInfo.BasicInfo.RequestStatus = 3
	profile.ProfileInfo.BasicInfo.IsRequestPending = false

	// other user's members
	members := []model.UserMember{}
	err = session.Db.Table("u_member").Where("user_id = ?", otherUserID).OrderBy("love_point DESC").Find(&members)
	if err != nil {
		panic(err)
	}
	profile.ProfileInfo.TotalLovePoint = 0
	for _, member := range members {
		profile.ProfileInfo.TotalLovePoint += member.LovePoint
	}
	for i := 0; i < 3; i++ {
		profile.ProfileInfo.LoveMembers[i].MemberMasterID = members[i].MemberMasterID
		profile.ProfileInfo.LoveMembers[i].LovePoint = members[i].LovePoint
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

	partnerCards := FetchPartnerCards(otherUserID)
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
	liveStats := session.GetOtherUserLiveStats(otherUserID)
	for i, liveDifficultyType := range enum.LiveDifficultyTypes {
		profile.PlayInfo.LivePlayCount = append(profile.PlayInfo.LivePlayCount, liveDifficultyType)
		profile.PlayInfo.LivePlayCount = append(profile.PlayInfo.LivePlayCount, liveStats.LivePlayCount[i])
		profile.PlayInfo.LiveClearCount = append(profile.PlayInfo.LiveClearCount, liveDifficultyType)
		profile.PlayInfo.LiveClearCount = append(profile.PlayInfo.LiveClearCount, liveStats.LiveClearCount[i])
	}

	session.Db.Table("u_card").Where("user_id = ?", otherUserID).
		OrderBy("live_join_count DESC").Limit(3).Find(&profile.PlayInfo.JoinedLiveCardRanking)
	session.Db.Table("u_card").Where("user_id = ?", otherUserID).
		OrderBy("active_skill_play_count DESC").Limit(3).Find(&profile.PlayInfo.PlaySkillCardRanking)

	// custom profile
	customProfile := GetOtherUserSetProfile(otherUserID)
	if customProfile.VoltageLiveDifficultyID != 0 {
		profile.PlayInfo.MaxScoreLiveDifficulty.LiveDifficultyMasterID = customProfile.VoltageLiveDifficultyID
		profile.PlayInfo.MaxScoreLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(otherUserID, customProfile.VoltageLiveDifficultyID).MaxScore
	}
	if customProfile.ComboLiveDifficultyID != 0 {
		profile.PlayInfo.MaxComboLiveDifficulty.LiveDifficultyMasterID = customProfile.ComboLiveDifficultyID
		profile.PlayInfo.MaxComboLiveDifficulty.Score =
			session.GetOtherUserLiveDifficulty(otherUserID, customProfile.ComboLiveDifficultyID).MaxCombo
	}

	// can get from members[] to save sql
	session.Db.Table("u_member").Where("user_id = ?", otherUserID).Find(&profile.MemberInfo.UserMembers)
	return profile
}

func GetOtherUserSetProfile(otherUserID int) model.UserSetProfile {
	p := model.UserSetProfile{}
	exists, err := Engine.Table("u_custom_set_profile").Where("user_id = ?", otherUserID).Get(&p)
	utils.CheckErr(err)
	if !exists {
		p.UserID = otherUserID
	}
	return p
}

func (session *Session) GetUserSetProfile() model.UserSetProfile {
	return GetOtherUserSetProfile(session.UserStatus.UserID)
}

// doesn't need to return delta patch or submit at the start because we would need to fetch profile everytime we need this thing
func (session *Session) SetUserSetProfile(p model.UserSetProfile) {
	affected, err := session.Db.Table("u_custom_set_profile").Where("user_id = ?", session.UserStatus.UserID).
		AllCols().Update(&p)
	utils.CheckErr(err)
	if affected == 0 {
		// need to insert
		_, err = session.Db.Table("u_custom_set_profile").Insert(&p)
		utils.CheckErr(err)
	}
}
