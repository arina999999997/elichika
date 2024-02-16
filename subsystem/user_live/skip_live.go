package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/item"
	"elichika/klab"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_live_deck"
	"elichika/subsystem/user_member"
	"elichika/subsystem/user_status"
	"elichika/userdata"

	"reflect"
)

func SkipLive(session *userdata.Session, req request.SkipLiveRequest) response.SkipLiveResponse {
	gamedata := session.Gamedata

	user_content.RemoveContent(session, item.SkipTicket.Amount(req.TicketUseCount))

	session.UserStatus.LastLiveDifficultyId = req.LiveDifficultyMasterId
	liveDifficulty := gamedata.LiveDifficulty[req.LiveDifficultyMasterId]

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

	// if liveDifficulty.IsCountTarget { // counted toward target and profiles
	// 	// TODO(behavior): Check if this is counted toward card / clear usage and update that
	// }

	return resp
}
