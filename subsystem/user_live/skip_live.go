package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/klab"
	"elichika/subsystem/event"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_live_deck"
	"elichika/subsystem/user_member"
	"elichika/subsystem/user_member_guild"
	"elichika/subsystem/user_mission"
	"elichika/subsystem/user_status"
	"elichika/userdata"

	"reflect"
)

func SkipLive(session *userdata.Session, req request.SkipLiveRequest) response.SkipLiveResponse {
	gamedata := session.Gamedata

	if config.Conf.ResourceConfig().ConsumeSkipTicket {
		user_content.RemoveContent(session, item.SkipTicket.Amount(req.TicketUseCount))
	}

	session.UserStatus.LastLiveDifficultyId = req.LiveDifficultyMasterId
	liveDifficulty := gamedata.LiveDifficulty[req.LiveDifficultyMasterId]

	if config.Conf.ResourceConfig().ConsumeLp {
		user_status.AddUserLp(session, -liveDifficulty.ConsumedLP*req.TicketUseCount)
	}

	// daily limit
	if liveDifficulty.UnlockPattern == enum.LiveUnlockPatternDaily && config.Conf.ResourceConfig().ConsumeDailyLiveLimit {
		liveDailyMasterId := GetLiveDailyMasterId(session, liveDifficulty.Live.LiveId)
		if liveDailyMasterId != nil {
			// this could happen if the user started the song before today
			// if that's the case, we don't need to track the play anymore
			userLiveDaily := GetUserLiveDaily(session, *liveDailyMasterId)
			userLiveDaily.RemainingPlayCount -= req.TicketUseCount
			UpdateUserLiveDaily(session, userLiveDaily)
		}
	}

	resp := response.SkipLiveResponse{
		SkipLiveResult: client.SkipLiveResult{
			LiveDifficultyMasterId: req.LiveDifficultyMasterId,
			LiveDeckId:             req.DeckId,
			GainUserExp:            liveDifficulty.RewardUserExp * req.TicketUseCount,
		},
		UserModelDiff: &session.UserModel,
	}

	isCenter := map[int32]bool{}
	for _, memberMapping := range liveDifficulty.Live.LiveMemberMapping {
		if memberMapping.IsCenter && (memberMapping.Position <= 9) {
			isCenter[memberMapping.Position-1] = true
		}
	}
	rewardCenterLovePoint := int32(0)
	if len(isCenter) != 0 {
		// liella songs have no center
		rewardCenterLovePoint = klab.CenterBondGainBasedOnBondGain(liveDifficulty.RewardBaseLovePoint) * req.TicketUseCount / int32(len(isCenter))
	}

	for i := int32(1); i <= req.TicketUseCount; i++ {
		standardDrops, isRewardAccessoryInPresentBox := getLiveStandardDrops(session, nil, liveDifficulty)
		resp.SkipLiveResult.IsRewardAccessoryInPresentBox = resp.SkipLiveResult.IsRewardAccessoryInPresentBox || isRewardAccessoryInPresentBox

		additionalDrops, isRewardAccessoryInPresentBox := getSkipAdditionalDrops(session, (i%2 == 0), liveDifficulty)
		resp.SkipLiveResult.IsRewardAccessoryInPresentBox = resp.SkipLiveResult.IsRewardAccessoryInPresentBox || isRewardAccessoryInPresentBox

		resp.SkipLiveResult.Drops.Append(client.LiveResultContentPack{
			StandardDrops:   standardDrops,
			AdditionalDrops: additionalDrops,
		})
	}

	user_status.AddUserExp(session, resp.SkipLiveResult.GainUserExp)

	deck := user_live_deck.GetUserLiveDeck(session, req.DeckId)
	cardMasterIds := []int32{}
	for i := 1; i <= 9; i++ {
		cardMasterIds = append(cardMasterIds, reflect.ValueOf(deck).Field(1+i).Interface().(generic.Nullable[int32]).Value)
	}

	memberRepresentativeCard := make(map[int32]int32)
	memberLoveGained := make(map[int32]int32)
	for i, cardMasterId := range cardMasterIds {
		addedLove := liveDifficulty.RewardBaseLovePoint * req.TicketUseCount
		if isCenter[int32(i)] {
			addedLove += rewardCenterLovePoint
		}
		memberMasterId := gamedata.Card[cardMasterId].Member.Id

		_, exist := memberRepresentativeCard[memberMasterId]
		// only use 1 card master id or an idol might be shown multiple times
		if !exist {
			memberRepresentativeCard[memberMasterId] = cardMasterId
		}
		memberLoveGained[memberMasterId] += addedLove
	}
	// it's normal to show +0 on the bond screen if the person is already maxed
	// this is checked against (video) recording
	for _, cardMasterId := range cardMasterIds {
		memberMasterId := gamedata.Card[cardMasterId].Member.Id
		if memberRepresentativeCard[memberMasterId] != cardMasterId {
			continue
		}
		addedLove := user_member.AddMemberLovePoint(session, memberMasterId, memberLoveGained[memberMasterId])
		resp.SkipLiveResult.MemberLoveStatuses.Set(cardMasterId, client.LiveResultMemberLoveStatus{
			RewardLovePoint: addedLove,
		})
	}

	// member guild
	memberGuildMemberMasterId := session.UserStatus.MemberGuildMemberMasterId
	if memberGuildMemberMasterId.HasValue {
		loveGained, exist := memberLoveGained[session.UserStatus.MemberGuildMemberMasterId.Value]
		if exist && (user_member_guild.IsMemberGuildRankingPeriod(session)) {
			lovePointAdded := user_member_guild.AddLovePoint(session, loveGained)
			resp.SkipLiveResult.LiveResultMemberGuild = generic.NewNullable(client.LiveResultMemberGuild{
				MemberGuildId:       user_member_guild.GetCurrentMemberGuildId(session),
				ReceiveLovePoint:    lovePointAdded,
				ReceiveVoltagePoint: 0,
				TotalPoint:          user_member_guild.GetCurrentUserMemberGuildTotalPoint(session),
			})
		}
	}

	// events
	activeEvent := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if (activeEvent != nil) && (activeEvent.ExpiredAt > session.Time.Unix()) {
		if activeEvent.EventType == enum.EventType1Marathon {
			if req.LiveEventMarathonStatus.HasValue {
				totalDeckBonus := int32(0)
				marathonEvent := session.Gamedata.EventMarathon[activeEvent.EventId]
				for _, cardMasterId := range cardMasterIds {
					userCard := user_card.GetUserCard(session, cardMasterId)
					memberId := session.Gamedata.Card[cardMasterId].Member.Id
					bonusArray, exist := marathonEvent.CardBonus[cardMasterId]
					if exist {
						totalDeckBonus += bonusArray[userCard.Grade]
					}
					totalDeckBonus += marathonEvent.MemberBonus[memberId]
				}
				resp.SkipLiveResult.ActiveEventResult = event.GetLiveResultActiveEventMarathon(session,
					liveDifficulty, liveDifficulty.EvaluationSScore, totalDeckBonus, req.TicketUseCount, req.LiveEventMarathonStatus.Value.IsUseEventMarathonBooster)
			}
		} else {
			panic("event type not supported")
		}
	}

	// if liveDifficulty.IsCountTarget { // counted toward target and profiles
	// 	// TODO(behavior): Check if this is counted toward card / clear usage and update that
	// }

	// mission stuff
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountClearedLive,
		&liveDifficulty.Live.LiveId, nil, user_mission.AddProgressHandler, req.TicketUseCount)
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountPlayLive,
		&liveDifficulty.Live.LiveId, nil, user_mission.AddProgressHandler, req.TicketUseCount)
	if liveDifficulty.UnlockPattern == enum.LiveUnlockPatternDaily {
		user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountPlayLiveDailyMusic,
			&liveDifficulty.Live.LiveId, nil, user_mission.AddProgressHandler, req.TicketUseCount)
		user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountClearedLiveDailyMusic,
			&liveDifficulty.Live.LiveId, nil, user_mission.AddProgressHandler, req.TicketUseCount)
	}
	return resp
}
