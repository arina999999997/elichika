package userdata

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/utils"
)

// fetch a card, use the value in diff is present, otherwise fetch from db
func (session *Session) GetUserCard(cardMasterId int32) client.UserCard {
	ptr, exist := session.UserModel.UserCardByCardId.Get(cardMasterId)
	if exist {
		return *ptr
	}
	card := client.UserCard{}
	exist, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Get(&card)
	utils.CheckErr(err)

	if !exist {
		gamedata := session.Ctx.MustGet("gamedata").(*gamedata.Gamedata)
		masterCard := gamedata.Card[cardMasterId]
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

func (session *Session) UpdateUserCard(card client.UserCard) {
	session.UserModel.UserCardByCardId.Set(card.CardMasterId, card)
}

func cardFinalizer(session *Session) {
	for _, card := range session.UserModel.UserCardByCardId.Map {
		affected, err := session.Db.Table("u_card").
			Where("user_id = ? AND card_master_id = ?", session.UserId, card.CardMasterId).AllCols().Update(*card)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_card", *card)
		}
	}
}

// insert all the cards
func (session *Session) InsertCards(cards []client.UserCard) {
	for _, card := range cards {
		session.UserModel.UserCardByCardId.Set(card.CardMasterId, card)
	}
}

func init() {
	addFinalizer(cardFinalizer)
}
