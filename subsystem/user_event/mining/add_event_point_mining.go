package mining

import (
	"elichika/client"
	"elichika/generic"
	"elichika/subsystem/user_story_event_history"
	"elichika/userdata"
)

// result is already partially filled
func AddEventPoint(session *userdata.Session, gainedPoint int32, result *client.LiveResultActiveEvent) {
	event := session.Gamedata.EventActive.GetEventMining()
	userEventMining := GetUserEventMining(session)

	ranking := GetPointRanking(session.Db, event.EventId)
	ranking.AddScore(session.UserId, gainedPoint)

	beforePoint := userEventMining.EventPoint
	afterPoint := beforePoint + gainedPoint
	// this list is in reverse order from higher point required to lower point
	for _, story := range event.TopStatus.StoryStatus.Stories.Slice {
		if story.StoryNumber <= userEventMining.ReadStoryNumber { // done for sure
			break
		}
		if (story.RequiredEventPoint > beforePoint) && (story.RequiredEventPoint <= afterPoint) {
			if userEventMining.OpenedStoryNumber < story.StoryNumber {
				userEventMining.OpenedStoryNumber = story.StoryNumber
			}
			eventStory := session.Gamedata.EventStory[story.EventMiningStoryId]
			result.OpenedEventStory = generic.NewNullable(client.EventResultOpenedNewStory{
				Title:                 eventStory.Title.DotUnderText,
				PreviewImageAssetPath: eventStory.StoryDetailThumbnailPath,
			})

			// permanantly unlock this in the story history too
			// official server would do this when the event disappear
			user_story_event_history.UnlockEventStory(session, story.EventMiningStoryId)
		}
	}
	userEventMining.EventPoint = afterPoint
	session.UserModel.UserEventMiningByEventMasterId.Set(userEventMining.EventMasterId, userEventMining)
	result.TotalPoint = client.LiveResultActiveEventPoint{
		Point: afterPoint,
		// BonusParam: 10000, // this field is unused
	}
}
