package event

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/subsystem/user_event/marathon"
	"elichika/userdata"

	"fmt"
)

func GetLiveResultActiveEventMarathon(session *userdata.Session, liveDifficulty *gamedata.LiveDifficulty, score, deckBonusFactor, loopCount int32, useBoosterItem bool) generic.Nullable[client.LiveResultActiveEvent] {
	getBaseEventPointMarathon := func() int32 {
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
		// numbers from https://suyo.be/sifas/wiki/events/story/event-points
		// TODO(extra): the event point rewards has changed with time
		// so maybe we'd want to load this from a database instead
		switch liveDifficulty.ConsumedLP {
		case 9:
			return []int32{148, 141, 135, 128, 121}[id]
		case 10:
			return []int32{165, 157, 150, 142, 135}[id]
		case 12:
			return []int32{243, 234, 225, 216, 207}[id]
		case 13:
			return []int32{263, 253, 243, 234, 224}[id]
		case 15:
			return []int32{360, 348, 337, 326, 315}[id]
		case 16:
			return []int32{384, 372, 360, 348, 336}[id]
		case 20:
			return []int32{525, 516, 507, 498, 489}[id]
		default:
			panic(fmt.Sprint("not supported LP amount: ", liveDifficulty.ConsumedLP))
		}
		return 0
	}
	// the unit of measurement is /10000
	epPerClear := getBaseEventPointMarathon()
	baseFactor := int32(10000) // this is the base 100% + the bonus from items
	if useBoosterItem {
		baseFactor += 5000 // TODO(hardcoded): load from m_event_marathon_booster_item
	}

	// skip ticket doesn't give event point bonus for multiple songs like drops
	// so skiping xk is similar to skip k time
	basePointTotal := (baseFactor * epPerClear / 10000) * loopCount
	deckBonusPointTotal := (deckBonusFactor * epPerClear / 10000) * loopCount

	// this should be fine because this code will only run when event isn't nil
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	marathonEvent := session.Gamedata.EventMarathon[event.EventId]
	result := client.LiveResultActiveEvent{
		EventId:            event.EventId,
		EventType:          enum.EventType1Marathon,
		EventLogoAssetPath: marathonEvent.TopStatus.TitleImagePath,
		ReceivePoint: client.LiveResultActiveEventPoint{ // raw point rewards, plus the ticket(?)
			Point:      basePointTotal,
			BonusParam: baseFactor,
		},
		// TotalPoint: // filled by user_event/marathon
		BonusPoint: client.LiveResultActiveEventPoint{
			Point:      deckBonusPointTotal,
			BonusParam: deckBonusFactor,
		},
		// OpenedEventStory: // filled by user_event/marathon
		// LiveEventDropItemInfo: // filled by user_event/marathon
		// PointReward: // filled by user_event/marathon,
		IsStartLoopReward: false,
	}
	if useBoosterItem {
		marathon.RemoveBoosterItem(session, loopCount)
	}
	marathon.AddEventPoint(session, basePointTotal+deckBonusPointTotal, &result)
	return generic.NewNullable(result)
}
