package userdata

import (
	"elichika/client"
)

// TODO(multiplayer): This only work for single player
func (session *Session) GetTowerRankingUser() client.TowerRankingUser {
	card := session.GetUserCard(session.UserStatus.RecommendCardMasterId)
	return client.TowerRankingUser{
		UserId:                 int32(session.UserId),
		UserName:               session.UserStatus.Name,
		UserRank:               session.UserStatus.Rank,
		CardMasterId:           card.CardMasterId,
		Level:                  card.Level,
		IsAwakening:            card.IsAwakening,
		IsAllTrainingActivated: card.IsAllTrainingActivated,
		EmblemMasterId:         session.UserStatus.EmblemId,
	}
}

func (session *Session) GetRankingUser() client.RankingUser {
	card := session.GetUserCard(session.UserStatus.RecommendCardMasterId)
	return client.RankingUser{
		UserId:                 int32(session.UserId),
		Name:                   session.UserStatus.Name,
		Rank:                   session.UserStatus.Rank,
		FavoriteCardMasterId:   card.CardMasterId,
		FavoriteCardLevel:      card.Level,
		IsAwakeningImage:       card.IsAwakening,
		IsAllTrainingActivated: card.IsAllTrainingActivated,
		EmblemId:               session.UserStatus.EmblemId,
	}
}
