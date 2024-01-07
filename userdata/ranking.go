package userdata

import (
	"elichika/client"
	"elichika/protocol/response"
)

func (session *Session) GetRankingUser() response.RankingUser {
	card := session.GetUserCard(int(session.UserStatus.RecommendCardMasterId))
	return response.RankingUser{
		UserId: session.UserStatus.UserId,
		UserName: client.LocalizedText{
			DotUnderText: session.UserStatus.Name.DotUnderText,
		},
		UserRank:               int(session.UserStatus.Rank),
		CardMasterId:           card.CardMasterId,
		Level:                  card.Level,
		IsAwakening:            card.IsAwakening,
		IsAllTrainingActivated: card.IsAllTrainingActivated,
		EmblemMasterId:         int(session.UserStatus.EmblemId),
	}
}
