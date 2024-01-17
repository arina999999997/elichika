package userdata

import (
	"elichika/client"
	"elichika/protocol/response"
)

func (session *Session) GetTowerRankingUser() response.TowerRankingUser {
	card := session.GetUserCard(session.UserStatus.RecommendCardMasterId)
	return response.TowerRankingUser{
		UserId: session.UserId,
		UserName: client.LocalizedText{
			DotUnderText: session.UserStatus.Name.DotUnderText,
		},
		UserRank:               int(session.UserStatus.Rank),
		CardMasterId:           int(card.CardMasterId),
		Level:                  int(card.Level),
		IsAwakening:            card.IsAwakening,
		IsAllTrainingActivated: card.IsAllTrainingActivated,
		EmblemMasterId:         int(session.UserStatus.EmblemId),
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
