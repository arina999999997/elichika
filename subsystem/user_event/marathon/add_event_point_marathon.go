package marathon

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_story_event_history"
	"elichika/userdata"
)

// result is already partially filled
func AddEventPoint(session *userdata.Session, gainedPoint int32, result *client.LiveResultActiveEvent) {
	event := session.Gamedata.EventActive.GetEventMarathon()
	userEventMarathon := GetUserEventMarathon(session)

	ranking := GetRanking(session.Db, event.EventId)
	ranking.AddScore(session.UserId, gainedPoint)

	beforePoint := userEventMarathon.EventPoint
	afterPoint := beforePoint + gainedPoint
	// this list is in reverse order from higher point required to lower point
	for _, story := range event.TopStatus.StoryStatus.Stories.Slice {
		if story.StoryNumber <= userEventMarathon.ReadStoryNumber { // done for sure
			break
		}
		if (story.RequiredEventPoint > beforePoint) && (story.RequiredEventPoint <= afterPoint) {
			if userEventMarathon.OpenedStoryNumber < story.StoryNumber {
				userEventMarathon.OpenedStoryNumber = story.StoryNumber
			}
			eventStory := session.Gamedata.EventStory[story.EventMarathonStoryId]
			result.OpenedEventStory = generic.NewNullable(client.EventResultOpenedNewStory{
				Title:                 eventStory.Title.DotUnderText,
				PreviewImageAssetPath: eventStory.StoryDetailThumbnailPath,
			})

			// permanantly unlock this in the story history too
			// official server would do this when the event disappear
			user_story_event_history.UnlockEventStory(session, story.EventMarathonStoryId)
		}
	}
	userEventMarathon.EventPoint = afterPoint
	session.UserModel.UserEventMarathonByEventMasterId.Set(userEventMarathon.EventMasterId, userEventMarathon)
	result.TotalPoint = client.LiveResultActiveEventPoint{
		Point: afterPoint,
		// BonusParam: 10000, // this field is unused
	}

	activeEventPointReward := client.LiveResultActiveEventPointReward{}
	for _, pointReward := range event.TopStatus.EventMarathonPointRewardMasterRows.Slice {
		if (pointReward.RequiredPoint <= afterPoint) && (pointReward.RequiredPoint > beforePoint) {
			for _, content := range session.Gamedata.EventMarathonReward[pointReward.RewardGroupId] {
				user_present.AddPresentWithDuration(session, client.PresentItem{
					Content:          *content,
					PresentRouteType: enum.PresentRouteTypeEventMarathonPointReward,
					PresentRouteId:   generic.NewNullable(event.EventId),
				}, user_present.Duration30Days)
				activeEventPointReward.GettedPointRewards.Append(client.EventMarathonPointReward{
					RewardContent:     *content,
					RequiredPoint:     pointReward.RequiredPoint,
					IsStartLoopReward: false,
				})
			}
		} else if pointReward.RequiredPoint > afterPoint {
			activeEventPointReward.NextPointReward = generic.NewNullable(client.EventMarathonPointReward{
				RewardContent:     *session.Gamedata.EventMarathonReward[pointReward.RewardGroupId][0],
				RequiredPoint:     pointReward.RequiredPoint,
				IsStartLoopReward: false,
			})
			break
		}
	}
	result.PointReward = generic.NewNullable(activeEventPointReward)
}
