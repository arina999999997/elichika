package userdata

import (
	"elichika/client"
	"elichika/protocol/response"
)

func (session *Session) GetRankingUser() response.RankingUser {
	card := session.GetUserCard(session.UserStatus.RecommendCardMasterId)
	return response.RankingUser{
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
