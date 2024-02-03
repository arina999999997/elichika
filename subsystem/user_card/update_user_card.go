package user_card

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserCard(session *userdata.Session, card client.UserCard) {
	session.UserModel.UserCardByCardId.Set(card.CardMasterId, card)
}