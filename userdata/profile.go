package userdata

import (
	"elichika/client"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
)

func FetchDBProfile(userId int32, result interface{}) {
	exist, err := Engine.Table("u_status").Where("user_id = ?", userId).Get(result)
	utils.CheckErrMustExist(err, exist)
}

func FetchPartnerCards(otherUserId int32) []client.UserCard {
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
	exist, err := Engine.Table("u_status").Where("user_id = ?", otherUserId).Get(&otherUserStatus)
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

func (session *Session) GetOtherUser(otherUserId int32) client.OtherUser {
	otherUser := client.OtherUser{}
	FetchDBProfile(otherUserId, &otherUser)
	recommendCard := GetOtherUserCard(otherUserId, otherUser.RecommendCardMasterId)

	otherUser.RecommendCardLevel = recommendCard.Level
	otherUser.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	otherUser.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated

	// TODO(friend): not implemented
	// otherUser.FriendApprovedAt = new(int64)
	// *otherUser.FriendApprovedAt = 0
	// otherUser.RequestStatus = 3
	// otherUser.IsRequestPending = false
	return otherUser
}

func (session *Session) GetOtherUserLiveStats(otherUserId int32) model.UserProfileLiveStats {
	stats := model.UserProfileLiveStats{}
	_, err := Engine.Table("u_status").Where("user_id = ?", otherUserId).Get(&stats)
	utils.CheckErr(err)
	return stats
}

func (session *Session) GetUserLiveStats() model.UserProfileLiveStats {
	return session.GetOtherUserLiveStats(session.UserId)
}

func (session *Session) UpdateUserLiveStats(stats model.UserProfileLiveStats) {
	_, err := session.Db.Table("u_status").Where("user_id = ?", session.UserId).AllCols().Update(&stats)
	utils.CheckErr(err)
}

// fetch profile of another user, from session.UserId's perspective
// it's possible that otherUserId == session.UserId

func (session *Session) GetOtherUserSetProfile(otherUserId int32) client.UserSetProfile {
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
