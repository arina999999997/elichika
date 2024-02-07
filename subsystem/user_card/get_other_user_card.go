package user_card

import (
	"elichika/client"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserCard(session *userdata.Session, otherUserId, cardMasterId int32) client.OtherUserCard {
	userCard := client.UserCard{}
	exist, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", otherUserId, cardMasterId).Get(&userCard)
	utils.CheckErr(err)
	if !exist {
		panic("other user card doesn't exist")
	}
	otherUserCard := client.OtherUserCard{
		CardMasterId: userCard.CardMasterId,
		Level:        userCard.Level,
		Grade:        userCard.Grade,
		// LoveLevel: session.Gamedata.LoveLevelFromLovePoint(userCard.LovePoint),
		LoveLevel:              1, // love level is calculated from card love point, not member love point, and thus it's always 1
		IsAwakening:            userCard.IsAwakening,
		IsAwakeningImage:       userCard.IsAwakeningImage,
		IsAllTrainingActivated: userCard.IsAllTrainingActivated,
		ActiveSkillLevel:       userCard.ActiveSkillLevel,
		// PassiveSkillLevels:
		// AdditionalPassiveSkillIds:
		MaxFreePassiveSkill: userCard.MaxFreePassiveSkill,
		TrainingStamina:     userCard.TrainingLife,
		TrainingAppeal:      userCard.TrainingAttack,
		TrainingTechnique:   userCard.TrainingDexterity,
		// MemberLovePanels:
	}
	otherUserCard.PassiveSkillLevels.Append(userCard.PassiveSkillALevel)
	otherUserCard.PassiveSkillLevels.Append(userCard.PassiveSkillBLevel)
	// otherUserCard.PassiveSkillLevels.Append(userCard.PassiveSkillCLevel)
	if otherUserCard.MaxFreePassiveSkill >= 1 {
		otherUserCard.AdditionalPassiveSkillIds.Append(userCard.AdditionalPassiveSkill1Id)
	}
	if otherUserCard.MaxFreePassiveSkill >= 2 {
		otherUserCard.AdditionalPassiveSkillIds.Append(userCard.AdditionalPassiveSkill2Id)
	}
	if otherUserCard.MaxFreePassiveSkill >= 3 {
		otherUserCard.AdditionalPassiveSkillIds.Append(userCard.AdditionalPassiveSkill3Id)
	}
	if otherUserCard.MaxFreePassiveSkill >= 4 {
		otherUserCard.AdditionalPassiveSkillIds.Append(userCard.AdditionalPassiveSkill4Id)
	}
	memberId := session.Gamedata.Card[cardMasterId].Member.Id
	otherUserCard.MemberLovePanels.Append(user_member.GetOtherUserMemberLovePanel(session, otherUserId, memberId))
	return otherUserCard
}
