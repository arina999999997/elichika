package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetUserLiveParty(partyId int) model.UserLiveParty {
	liveParty := model.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND party_id = ?", session.UserStatus.UserId, partyId).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		panic("Party doesn't exist")
	}
	return liveParty
}

func (session *Session) GetUserLivePartiesWithDeckId(deckId int) []model.UserLiveParty {
	liveParties := []model.UserLiveParty{}
	err := session.Db.Table("u_live_party").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserStatus.UserId, deckId).
		OrderBy("party_id").Find(&liveParties)
	utils.CheckErr(err)
	return liveParties
}

func (session *Session) GetUserLivePartyWithDeckAndCardId(deckId int, cardId int) model.UserLiveParty {
	liveParty := model.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND user_live_deck_id = ? AND (card_master_id_1 = ? OR card_master_id_2 = ? OR card_master_id_3 = ?)",
			session.UserStatus.UserId, deckId, cardId, cardId, cardId).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		panic("Party doesn't exist")
	}
	return liveParty
}

func (session *Session) UpdateUserLiveParty(liveParty model.UserLiveParty) {
	session.UserLivePartyMapping.SetList(&session.UserModel.UserLivePartyById).Update(liveParty)
}

func livePartyFinalizer(session *Session) {
	for _, party := range session.UserModel.UserLivePartyById.Objects {
		affected, err := session.Db.Table("u_live_party").
			Where("user_id = ? AND party_id = ?", session.UserStatus.UserId, party.PartyId).AllCols().
			Update(party)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_live_party").Insert(party)
			utils.CheckErr(err)
		}
	}
}

func (session *Session) GetAllLivePartiesWithAccessory(accessoryId int64) []model.UserLiveParty {
	parties := []model.UserLiveParty{}
	err := session.Db.Table("u_live_party").
		Where("user_id = ? AND (user_accessory_id_1 = ? OR user_accessory_id_2 = ? OR user_accessory_id_3 = ? )",
			session.UserStatus.UserId, accessoryId, accessoryId, accessoryId).Find(&parties)
	utils.CheckErr(err)
	return parties
}

func (session *Session) InsertLiveParties(parties []model.UserLiveParty) {
	session.UserModel.UserLivePartyById.Objects = append(session.UserModel.UserLivePartyById.Objects, parties...)
}

func init() {
	addFinalizer(livePartyFinalizer)
	addGenericTableFieldPopulator("u_live_party", "UserLivePartyById")
}
