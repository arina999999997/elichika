package event

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/log"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_event/mining"
	"elichika/userdata"

	"fmt"
)

// Handle a live result and return the relevant data
// Convention: loop count = 0 means manual live (calculated for top score)
// loop count > 0 means skip ticket was used
func HandleLiveResultActiveEventMining(session *userdata.Session, liveDifficulty *gamedata.LiveDifficulty, score, deckBonusFactor, loopCount int32) generic.Nullable[client.LiveResultActiveEvent] {
	isManualClear := loopCount == 0
	if isManualClear {
		loopCount = 1
	}
	getBaseEventPointMining := func() int32 {
		id := 0
		if score < liveDifficulty.EvaluationSScore {
			id++
		}
		if score < liveDifficulty.EvaluationAScore {
			id++
		}
		if score < liveDifficulty.EvaluationBScore {
			id++
		}
		if score < liveDifficulty.EvaluationCScore {
			id++
		}
		// event point is based entirely on LP spent
		// numbers from https://suyo.be/sifas/wiki/events/itex/event-points
		// TODO(extra): the event point rewards has changed with time
		// so maybe we'd want to load this from a database instead
		switch liveDifficulty.ConsumedLP {
		case 9:
			return []int32{247, 236, 225, 213, 202}[id]
		case 10:
			return []int32{275, 262, 250, 237, 225}[id]
		case 12:
			return []int32{405, 390, 375, 360, 345}[id]
		case 13:
			return []int32{438, 422, 406, 390, 373}[id]
		case 15:
			return []int32{600, 581, 562, 543, 525}[id]
		case 16:
			return []int32{640, 620, 600, 580, 560}[id]
		case 20:
			return []int32{875, 860, 845, 830, 815}[id]
		default:
			log.Panic(fmt.Sprint("not supported LP amount: ", liveDifficulty.ConsumedLP))
		}
		return 0
	}
	getBaseEventCurrencyMining := func() int32 {
		id := 0
		if score < liveDifficulty.EvaluationSScore {
			id++
		}
		if score < liveDifficulty.EvaluationAScore {
			id++
		}
		if score < liveDifficulty.EvaluationBScore {
			id++
		}
		if score < liveDifficulty.EvaluationCScore {
			id++
		}
		// https://suyo.be/sifas/wiki/events/itex/event-points
		switch liveDifficulty.LiveDifficultyType {
		case enum.LiveDifficultyTypeNormal: // beginner
			return []int32{135, 120, 105, 90, 75}[id]
		case enum.LiveDifficultyTypeHard: // intermediate
			return []int32{225, 210, 195, 180, 165}[id]
		case enum.LiveDifficultyTypeExpert: // advanced
			return []int32{345, 330, 315, 300, 285}[id]
		case enum.LiveDifficultyTypeExpertPlus: // expert
			fallthrough
		case enum.LiveDifficultyTypeExpertPlusPlus: // challenge
			fallthrough
		case enum.LiveDifficultyTypeMaster: // expert
			return []int32{525, 510, 495, 480, 465}[id]
		default:
			log.Panic(fmt.Sprint("not supported LiveDifficultyType: ", liveDifficulty.LiveDifficultyType))
		}
		return 0
	}
	// bonus apply to the currency drop, not to the event point, so this is just a simply multiply
	epPointTotal := getBaseEventPointMining() * loopCount

	// this should be fine because this code will only run when event isn't nil
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	eventMining := session.Gamedata.EventMining[event.EventId]
	result := client.LiveResultActiveEvent{
		EventId:            event.EventId,
		EventType:          enum.EventType1Mining,
		EventLogoAssetPath: eventMining.TopStatus.TitleImagePath,
		ReceivePoint: client.LiveResultActiveEventPoint{
			Point:      epPointTotal,
			BonusParam: 10000, // note that this is set to 10000 = 100% because the bonus text is not shown for event mining
		},
		// TotalPoint: // filled by user_event/mining
		// BonusPoint: // only used by event marathon
		// OpenedEventStory: // filled by user_event/mining
		LiveEventDropItemInfo: generic.NewNullable(client.LiveEventDropItemInfo{
			BonusRate:          deckBonusFactor,
			IsSendToPresentBox: false,
		}),
		// PointReward: // only used by event marathon
		IsStartLoopReward: false,
	}
	itemPerClear := getBaseEventCurrencyMining() * deckBonusFactor / 10000
	itemDropTotal := itemPerClear * loopCount
	itemDropTotalBase := getBaseEventCurrencyMining() * loopCount
	result.LiveEventDropItemInfo.Value.LiveEventDropContents.Append(client.LiveEventDropContents{})
	result.LiveEventDropItemInfo.Value.LiveEventDropContents.Slice[0].StandardDrops.Append(client.Content{
		ContentType:   enum.ContentTypeExchangeEventPoint,
		ContentId:     eventMining.TopStatus.EventPointMasterId,
		ContentAmount: itemDropTotalBase,
	})
	result.LiveEventDropItemInfo.Value.LiveEventDropContents.Slice[0].BonusDrops.Append(client.Content{
		ContentType:   enum.ContentTypeExchangeEventPoint,
		ContentId:     eventMining.TopStatus.EventPointMasterId,
		ContentAmount: itemDropTotal - itemDropTotalBase,
	})
	user_content.AddContent(session, result.LiveEventDropItemInfo.Value.LiveEventDropContents.Slice[0].StandardDrops.Slice[0])

	mining.AddEventPoint(session, epPointTotal, &result)
	if isManualClear {
		mining.TrackClearVoltage(session, liveDifficulty, score)
	}

	return generic.NewNullable(result)
}
