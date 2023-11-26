package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetUserLiveParty(partyID int) model.UserLiveParty {
	liveParty := model.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND party_id = ?", session.UserStatus.UserID, partyID).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		panic("Party doesn't exist")
	}
	return liveParty
}

func (session *Session) GetUserLivePartiesWithDeckID(deckID int) []model.UserLiveParty {
	liveParties := []model.UserLiveParty{}
	err := session.Db.Table("u_live_party").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserStatus.UserID, deckID).
		OrderBy("party_id").Find(&liveParties)
	utils.CheckErr(err)
	return liveParties
}

func (session *Session) GetUserLivePartyWithDeckAndCardID(deckID int, cardID int) model.UserLiveParty {
	liveParty := model.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND user_live_deck_id = ? AND (card_master_id_1 = ? OR card_master_id_2 = ? OR card_master_id_3 = ?)",
			session.UserStatus.UserID, deckID, cardID, cardID, cardID).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		panic("Party doesn't exist")
	}
	return liveParty
}

func (session *Session) UpdateUserLiveParty(liveParty model.UserLiveParty) {
	session.UserLivePartyMapping.SetList(&session.UserModel.UserLivePartyByID).Update(liveParty)
}

func livePartyFinalizer(session *Session) {
	for _, party := range session.UserModel.UserLivePartyByID.Objects {
		affected, err := session.Db.Table("u_live_party").
			Where("user_id = ? AND party_id = ?", session.UserStatus.UserID, party.PartyID).AllCols().
			Update(party)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_live_party").Insert(party)
			utils.CheckErr(err)
		}
	}
}

func (session *Session) GetAllLivePartiesWithAccessory(accessoryID int64) []model.UserLiveParty {
	parties := []model.UserLiveParty{}
	err := session.Db.Table("u_live_party").
		Where("user_id = ? AND (user_accessory_id_1 = ? OR user_accessory_id_2 = ? OR user_accessory_id_3 = ? )",
			session.UserStatus.UserID, accessoryID, accessoryID, accessoryID).Find(&parties)
	utils.CheckErr(err)
	return parties
}

func (session *Session) InsertLiveParties(parties []model.UserLiveParty) {
	session.UserModel.UserLivePartyByID.Objects = append(session.UserModel.UserLivePartyByID.Objects, parties...)
}

func init() {
	addFinalizer(livePartyFinalizer)
	addGenericTableFieldPopulator("u_live_party", "UserLivePartyByID")
}
