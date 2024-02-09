package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/klab"
	"elichika/subsystem/user_accessory"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_live_deck"
	"elichika/subsystem/user_live_difficulty"
	"elichika/subsystem/user_live_party"
	"elichika/subsystem/user_member"
	"elichika/subsystem/user_profile"
	"elichika/subsystem/user_status"
	"elichika/subsystem/user_story_main"
	"elichika/subsystem/voltage_ranking"
	"elichika/userdata"

	"reflect"
)

func liveTypeManualHandler(session *userdata.Session, req request.FinishLiveRequest, live client.Live, startReq request.StartLiveRequest) response.FinishLiveResponse {
	gamedata := session.Gamedata
	liveDifficulty := gamedata.LiveDifficulty[session.UserStatus.LastLiveDifficultyId]

	userLiveDifficulty := user_live_difficulty.GetUserLiveDifficulty(session, session.UserStatus.LastLiveDifficultyId)
	userLiveDifficulty.IsNew = false
	userLiveDifficulty.IsAutoplay = startReq.IsAutoPlay

	resp := response.FinishLiveResponse{
		LiveResult: client.LiveResult{
			LiveDifficultyMasterId: session.UserStatus.LastLiveDifficultyId,
			LiveDeckId:             session.UserStatus.LatestLiveDeckId,
			Voltage:                req.LiveScore.CurrentScore,
			BeforeUserExp:          session.UserStatus.Exp,
			LiveFinishStatus:       req.LiveFinishStatus,
			LastBestVoltage:        userLiveDifficulty.MaxScore,
		},
		UserModelDiff: &session.UserModel,
	}

	isCenter := map[int32]bool{}

	for _, memberMapping := range liveDifficulty.Live.LiveMemberMapping {
		if memberMapping.IsCenter && (memberMapping.Position <= 9) {
			isCenter[int32(memberMapping.Position-1)] = true
		}
	}

	rewardCenterLovePoint := int32(0)
	if len(isCenter) != 0 {
		// liella songs have no center
		rewardCenterLovePoint = klab.CenterBondGainBasedOnBondGain(liveDifficulty.RewardBaseLovePoint) / int32(len(isCenter))
	}

	// build the last played deck
	lastPlayDeck := client.LastPlayLiveDifficultyDeck{
		LiveDifficultyId: resp.LiveResult.LiveDifficultyMasterId,
		Voltage:          req.LiveScore.CurrentScore,
		IsCleared:        req.LiveFinishStatus == enum.LiveFinishStatusSucceeded,
		RecordedAt:       session.Time.Unix(),
	}

	userLiveDeck := user_live_deck.GetUserLiveDeck(session, session.UserStatus.LatestLiveDeckId)
	for position := 1; position <= 9; position++ {
		cardMasterId := reflect.ValueOf(userLiveDeck).Field(1 + position).Interface().(generic.Nullable[int32]).Value
		suitMasterId := reflect.ValueOf(userLiveDeck).Field(1 + position + 9).Interface().(generic.Nullable[int32]).Value
		lastPlayDeck.CardWithSuitDict.Set(cardMasterId, suitMasterId)
	}
	liveParties := user_live_party.GetUserLivePartiesWithDeckId(session, session.UserStatus.LatestLiveDeckId)
	for _, liveParty := range liveParties {
		liveSquad := client.LiveSquad{}
		liveSquad.CardMasterIds.Append(liveParty.CardMasterId1.Value)
		liveSquad.CardMasterIds.Append(liveParty.CardMasterId2.Value)
		liveSquad.CardMasterIds.Append(liveParty.CardMasterId3.Value)
		liveSquad.UserAccessoryIds.Append(liveParty.UserAccessoryId1)
		liveSquad.UserAccessoryIds.Append(liveParty.UserAccessoryId2)
		liveSquad.UserAccessoryIds.Append(liveParty.UserAccessoryId3)
		lastPlayDeck.SquadDict.Set(liveParty.PartyId%10-1, liveSquad)
	}

	userLiveDifficulty.PlayCount++
	resp.LiveResult.LiveResultAchievements.Set(1, client.LiveResultAchievement{
		Position:          1,
		IsAlreadyAchieved: userLiveDifficulty.ClearedDifficultyAchievement1.HasValue,
	})
	resp.LiveResult.LiveResultAchievements.Set(2, client.LiveResultAchievement{
		Position:          2,
		IsAlreadyAchieved: userLiveDifficulty.ClearedDifficultyAchievement2.HasValue,
	})
	resp.LiveResult.LiveResultAchievements.Set(3, client.LiveResultAchievement{
		Position:          3,
		IsAlreadyAchieved: userLiveDifficulty.ClearedDifficultyAchievement3.HasValue,
	})

	if lastPlayDeck.IsCleared {
		// add story if it is a story mode
		if live.CellId.HasValue {
			user_story_main.InsertUserStoryMain(session, live.CellId.Value)
		}

		// update clear record
		userLiveDifficulty.ClearCount++
		if userLiveDifficulty.MaxScore < req.LiveScore.CurrentScore {
			userLiveDifficulty.MaxScore = req.LiveScore.CurrentScore
			// update voltage ranking
			userVoltageRanking := voltage_ranking.UserVoltageRanking{
				UserId:           session.UserId,
				LiveDifficultyId: userLiveDifficulty.LiveDifficultyId,
				VoltagePoint:     req.LiveScore.CurrentScore,
				DeckDetail: client.OtherUserDeckDetail{
					Deck: client.OtherUserDeck{
						Name: userLiveDeck.Name,
					},
				},
			}

			for _, liveParty := range liveParties {
				otherUserParty := client.OtherUserParty{
					Id: liveParty.PartyId,
				}
				otherUserParty.CardIds.Append(liveParty.CardMasterId1.Value)
				otherUserParty.CardIds.Append(liveParty.CardMasterId2.Value)
				otherUserParty.CardIds.Append(liveParty.CardMasterId3.Value)
				if liveParty.UserAccessoryId1.HasValue {
					otherUserParty.Accessories.Append(user_accessory.GetUserAccessory(session, liveParty.UserAccessoryId1.Value).ToOtherUserAccessory())
				}
				if liveParty.UserAccessoryId2.HasValue {
					otherUserParty.Accessories.Append(user_accessory.GetUserAccessory(session, liveParty.UserAccessoryId2.Value).ToOtherUserAccessory())
				}
				if liveParty.UserAccessoryId3.HasValue {
					otherUserParty.Accessories.Append(user_accessory.GetUserAccessory(session, liveParty.UserAccessoryId3.Value).ToOtherUserAccessory())
				}
				userVoltageRanking.DeckDetail.Deck.Parties.Append(otherUserParty)
			}
			for i := 1; i <= 9; i++ {
				cardMasterId := reflect.ValueOf(userLiveDeck).Field(1 + i).Interface().(generic.Nullable[int32]).Value
				suitMasterId := reflect.ValueOf(userLiveDeck).Field(1 + i + 9).Interface().(generic.Nullable[int32]).Value
				memberId := gamedata.Card[cardMasterId].Member.Id
				if !userVoltageRanking.DeckDetail.MemberLoveLevels.Has(memberId) {
					userVoltageRanking.DeckDetail.MemberLoveLevels.Set(memberId, user_member.GetMember(session, memberId).LoveLevel)
				}
				// no idea why the necessary stuff is like this, bad client code maybe
				otherUserCard := user_card.GetOtherUserCard(session, session.UserId, cardMasterId)
				for j := otherUserCard.AdditionalPassiveSkillIds.Size(); j < 9; j++ {
					otherUserCard.AdditionalPassiveSkillIds.Append(0) // this is the official server behaviour
				}
				otherUserCard.LoveLevel = *userVoltageRanking.DeckDetail.MemberLoveLevels.GetOnly(memberId)

				userVoltageRanking.DeckDetail.Deck.Cards.Append(otherUserCard)
				userVoltageRanking.DeckDetail.Deck.CardIds.Append(cardMasterId)
				userVoltageRanking.DeckDetail.Deck.SuitMasterIds.Append(suitMasterId)
			}

			voltage_ranking.SetVoltageRanking(session, userVoltageRanking)

		}
		if userLiveDifficulty.MaxCombo < req.LiveScore.HighestComboCount {
			userLiveDifficulty.MaxCombo = req.LiveScore.HighestComboCount
		}

		// award items
		for i, mission := range liveDifficulty.Missions {
			if (i == 0) || (int(req.LiveScore.CurrentScore) >= mission.TargetValue) {
				resp.LiveResult.LiveResultAchievements.Map[int32(i+1)].IsCurrentlyAchieved = true
				if !resp.LiveResult.LiveResultAchievements.Map[int32(i+1)].IsAlreadyAchieved { // new, add reward
					user_content.AddContent(session, mission.Reward)
					switch i {
					case 0:
						userLiveDifficulty.ClearedDifficultyAchievement1 = generic.NewNullable(int32(1))
					case 1:
						userLiveDifficulty.ClearedDifficultyAchievement2 = generic.NewNullable(int32(2))
					case 2:
						userLiveDifficulty.ClearedDifficultyAchievement3 = generic.NewNullable(int32(3))
					}
				}
			}
		}
		resp.LiveResult.GainUserExp = liveDifficulty.RewardUserExp
		user_status.AddUserExp(session, resp.LiveResult.GainUserExp)
	}

	memberRepresentativeCard := make(map[int32]int32)
	memberLoveGained := make(map[int32]int32)
	for i := range req.LiveScore.CardStatDict.Map {
		liveFinishCard := req.LiveScore.CardStatDict.Map[i]

		// calculate mvp
		if liveFinishCard.GotVoltage > resp.LiveResult.Mvp.GetVoltage {
			resp.LiveResult.Mvp.GetVoltage = liveFinishCard.GotVoltage
			resp.LiveResult.Mvp.CardMasterId = liveFinishCard.CardMasterId
			resp.LiveResult.Mvp.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			resp.LiveResult.Mvp.AppealCount = liveFinishCard.AppealCount
		}

		// update card stat and member bond if cleared
		if lastPlayDeck.IsCleared {

			addedLove := liveDifficulty.RewardBaseLovePoint
			if isCenter[i] {
				addedLove += rewardCenterLovePoint
			}

			userCardPlayCountStat := user_card.GetUserCardPlayCountStat(session, liveFinishCard.CardMasterId)
			userCardPlayCountStat.LiveJoinCount++
			userCardPlayCountStat.ActiveSkillPlayCount += liveFinishCard.SkillTriggeredCount
			user_card.UpdateUserCardPlayCountStat(session, userCardPlayCountStat)
			// update member love point
			memberMasterId := gamedata.Card[liveFinishCard.CardMasterId].Member.Id

			_, exist := memberRepresentativeCard[memberMasterId]
			// only use 1 card master id or an idol might be shown multiple times
			if !exist {
				memberRepresentativeCard[memberMasterId] = liveFinishCard.CardMasterId
			}
			memberLoveGained[memberMasterId] += int32(addedLove)
		}
	}
	// it's normal to show +0 on the bond screen if the person is already maxed
	// this is checked against (video) recording
	for _, i := range req.LiveScore.CardStatDict.OrderedKey {
		liveFinishCard := req.LiveScore.CardStatDict.Map[i]
		memberMasterId := gamedata.Card[liveFinishCard.CardMasterId].Member.Id
		if memberRepresentativeCard[memberMasterId] != i {
			continue
		}
		addedLove := user_member.AddMemberLovePoint(session, memberMasterId, memberLoveGained[memberMasterId])
		resp.LiveResult.MemberLoveStatuses.Set(liveFinishCard.CardMasterId, client.LiveResultMemberLoveStatus{
			RewardLovePoint: addedLove,
		})
	}

	resp.LiveResult.LiveResultAchievementStatus.ClearCount = userLiveDifficulty.ClearCount
	resp.LiveResult.LiveResultAchievementStatus.GotVoltage = req.LiveScore.CurrentScore
	resp.LiveResult.LiveResultAchievementStatus.RemainingStamina = req.LiveScore.RemainingStamina
	if live.LivePartnerCard.HasValue {
		resp.LiveResult.Partner = generic.NewNullable(user_profile.GetOtherUser(session, startReq.PartnerUserId))
	}

	user_live_difficulty.UpdateLiveDifficulty(session, userLiveDifficulty)
	user_live_difficulty.UpdateLastPlayLiveDifficultyDeck(session, lastPlayDeck)

	return resp
}
