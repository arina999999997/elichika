package userdata

import (
	"elichika/gamedata"
	"elichika/model"
	"elichika/utils"

	"fmt"
	"time"
)

// fetch a card, use the value in diff is present, otherwise fetch from db
func (session *Session) GetUserCard(cardMasterID int) model.UserCard {
	pos, exist := session.UserCardMapping.Map[int64(cardMasterID)]
	if exist {
		return session.UserModel.UserCardByCardID.Objects[pos]
	}
	card := model.UserCard{}
	exists, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", session.UserStatus.UserID, cardMasterID).Get(&card)
	utils.CheckErr(err)

	if !exists {
		gamedata := session.Ctx.MustGet("gamedata").(*gamedata.Gamedata)
		masterCard := gamedata.Card[cardMasterID]
		card = model.UserCard{
			UserID:              session.UserStatus.UserID,
			CardMasterID:        cardMasterID,
			Level:               1,
			MaxFreePassiveSkill: masterCard.PassiveSkillSlot,
			Grade:               -1, // check this for new card
			ActiveSkillLevel:    1,
			PassiveSkillALevel:  1,
			PassiveSkillBLevel:  1,
			PassiveSkillCLevel:  1,
			AcquiredAt:          time.Now().Unix(),
			IsNew:               true,
		}
	}
	return card
}

func (session *Session) UpdateUserCard(card model.UserCard) {
	session.UserCardMapping.SetList(&session.UserModel.UserCardByCardID).Update(card)
}

func cardFinalizer(session *Session) {
	for _, card := range session.UserModel.UserCardByCardID.Objects {
		affected, err := session.Db.Table("u_card").
			Where("user_id = ? AND card_master_id = ?", session.UserStatus.UserID, card.CardMasterID).AllCols().Update(card)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_card").Insert(card)
			utils.CheckErr(err)
		}
	}
}

// insert all the cards
func (session *Session) InsertCards(cards []model.UserCard) {
	affected, err := session.Db.Table("u_card").Insert(&cards)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted ", affected, " cards")
}

func init() {
	addFinalizer(cardFinalizer)
	addGenericTableFieldPopulator("u_card", "UserCardByCardID")
}
