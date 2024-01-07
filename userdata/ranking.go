package userdata

import (
	"elichika/model"
	"elichika/protocol/response"
)

func (session *Session) GetRankingUser() response.RankingUser {
	card := session.GetUserCard(session.UserStatus.RecommendCardMasterId)
	return response.RankingUser{
		UserId: session.UserStatus.UserId,
		UserName: model.LocalizedText{
			DotUnderText: session.UserStatus.Name.DotUnderText,
		},
		UserRank:               session.UserStatus.Rank,
		CardMasterId:           card.CardMasterId,
		Level:                  card.Level,
		IsAwakening:            card.IsAwakening,
		IsAllTrainingActivated: card.IsAllTrainingActivated,
		EmblemMasterId:         session.UserStatus.EmblemId,
	}
}
