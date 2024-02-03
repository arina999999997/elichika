package user_card

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

// fetch a card, use the value in diff is present, otherwise fetch from db
func GetUserCard(session *userdata.Session, cardMasterId int32) client.UserCard {
	ptr, exist := session.UserModel.UserCardByCardId.Get(cardMasterId)
	if exist {
		return *ptr
	}
	card := client.UserCard{}
	exist, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Get(&card)
	utils.CheckErr(err)

	if !exist {
		masterCard := session.Gamedata.Card[cardMasterId]
		card = client.UserCard{
			CardMasterId:        cardMasterId,
			Level:               1,
			MaxFreePassiveSkill: masterCard.PassiveSkillSlot,
			Grade:               -1, // check this for new card
			ActiveSkillLevel:    1,
			PassiveSkillALevel:  1,
			PassiveSkillBLevel:  1,
			PassiveSkillCLevel:  1,
			AcquiredAt:          int32(session.Time.Unix()),
			IsNew:               true,
		}
	}
	return card
}
