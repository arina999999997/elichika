package userdata

import (
	"elichika/client"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"
)

func (session *Session) GetOtherUserLiveDifficulty(otherUserId int, liveDifficultyId int32) client.UserLiveDifficulty {
	userLiveDifficulty := client.UserLiveDifficulty{}
	exist, err := session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND live_difficulty_id = ?", otherUserId, liveDifficultyId).
		Get(&userLiveDifficulty)
	if err != nil {
		panic(err)
	}
	if !exist {
		// userLiveDifficulty.UserId = otherUserId
		userLiveDifficulty.LiveDifficultyId = liveDifficultyId
		userLiveDifficulty.EnableAutoplay = true
		userLiveDifficulty.IsNew = true
	}
	return userLiveDifficulty
}

func (session *Session) GetLiveDifficulty(liveDifficultyId int32) client.UserLiveDifficulty {
	return session.GetOtherUserLiveDifficulty(session.UserId, liveDifficultyId)
}

func (session *Session) GetAllLiveDifficulties() []client.UserLiveDifficulty {
	records := []client.UserLiveDifficulty{}
	err := session.Db.Table("u_live_difficulty").Where("user_id = ?", session.UserId).
		Find(&records)
	utils.CheckErr(err)
	return records
}

func (session *Session) UpdateLiveDifficulty(userLiveDifficulty client.UserLiveDifficulty) {
	session.UserModel.UserLiveDifficultyByDifficultyId.Set(userLiveDifficulty.LiveDifficultyId, userLiveDifficulty)
}

func liveDifficultyFinalizer(session *Session) {
	for _, userLiveDifficulty := range session.UserModel.UserLiveDifficultyByDifficultyId.Map {
		updated, err := session.Db.Table("u_live_difficulty").
			Where("user_id = ? AND live_difficulty_id = ?", session.UserId, userLiveDifficulty.LiveDifficultyId).
			AllCols().Update(*userLiveDifficulty)
		utils.CheckErr(err)
		if updated == 0 {
			genericDatabaseInsert(session, "u_live_difficulty", *userLiveDifficulty)
		}
	}

}

func (session *Session) GetLastPlayLiveDifficultyDeck(liveDifficultyId int) *model.LastPlayLiveDifficultyDeck {
	lastPlayDeck := model.LastPlayLiveDifficultyDeck{}
	exist, err := session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND live_difficulty_id = ?", session.UserId, liveDifficultyId).
		Get(&lastPlayDeck)
	utils.CheckErr(err)
	if !exist {
		return nil
	}
	return &lastPlayDeck
}

func (session *Session) BuildLastPlayLiveDifficultyDeck(deckId, liveDifficultyId int) model.LastPlayLiveDifficultyDeck {
	lastPlayDeck := model.LastPlayLiveDifficultyDeck{
		LiveDifficultyId: liveDifficultyId,
		Voltage:          0,     // filled by handler
		IsCleared:        false, // filled by handler
		RecordedAt:       session.Time.Unix()}
	lastPlayDeck.CardWithSuitDict = make([]int, 18)
	userLiveDeck := session.GetUserLiveDeck(deckId)
	userLiveDeckJson, _ := json.Marshal(userLiveDeck)

	gjson.Parse(string(userLiveDeckJson)).ForEach(func(key, value gjson.Result) bool {
		k := key.String()
		if strings.Contains(k, "_master_id_") {
			id := int(k[len(k)-1] - '0')
			if strings.Contains(k, "card_master_id") {
				lastPlayDeck.CardWithSuitDict[id*2-2] = int(value.Int())
			} else {
				lastPlayDeck.CardWithSuitDict[id*2-1] = int(value.Int())
			}
		}
		return true
	})

	liveParties := session.GetUserLivePartiesWithDeckId(deckId)
	for i, party := range liveParties {
		squad := model.DeckSquadDict{}
		partyJson, _ := json.Marshal(party)
		gjson.Parse(string(partyJson)).ForEach(func(key, value gjson.Result) bool {
			k := key.String()
			if strings.Contains(k, "card_master_id_") {
				squad.CardMasterIds = append(squad.CardMasterIds, int32(value.Int()))
			} else if strings.Contains(k, "user_accessory_id_") {
				ptr := new(int64)
				*ptr = value.Int()
				squad.UserAccessoryIds = append(squad.UserAccessoryIds, ptr)
			}
			return true
		})
		lastPlayDeck.SquadDict = append(lastPlayDeck.SquadDict, i)
		lastPlayDeck.SquadDict = append(lastPlayDeck.SquadDict, squad)
	}

	return lastPlayDeck
}

func (session *Session) SetLastPlayLiveDifficultyDeck(deck model.LastPlayLiveDifficultyDeck) {
	// TODO: maybe this can be put in finalizer
	// always call after inserting the actual live play, so we can just update
	_, err := session.Db.Table("u_live_difficulty").Where("user_id = ? and live_difficulty_id = ?", session.UserId, deck.LiveDifficultyId).
		AllCols().Update(&deck)
	utils.CheckErr(err)
}

func init() {
	addFinalizer(liveDifficultyFinalizer)
}
