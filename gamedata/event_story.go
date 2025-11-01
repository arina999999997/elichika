package gamedata

import (
	"elichika/client"

	"elichika/utils"

	"xorm.io/xorm"
)

// both event marathon and event mining use this structure except for the id field
// TODO(extra) for now, these will be loaded from m_story_event_history_detail
// this means that the event story will still show up in both the story menu
// and if a new event is added, they have to be added to the story menu too

type EventStory struct {
	EventMasterId int32 `xorm:"'event_master_id'"`
	StoryEventId  int32 `xorm:"pk 'story_event_id'"`
	StoryNumber   int32 `xorm:"'story_number'"`
	// - TODO(event): dynamically load these
	// this is server sided, for now we use the arrays:
	// - [0, 300, 2100, 4500, 8000, 14000, 25000] for marathon event:
	//   - this is the numbers for the first event, as found in recordings on youtube
	RequiredEventPoint       int32                    `xorm:"-"`
	StoryBannerThumbnailPath client.TextureStruktur   `xorm:"varchar(255) 'banner_thumbnail_path'"`
	StoryDetailThumbnailPath client.TextureStruktur   `xorm:"varchar(255) 'detail_thumbnail_path'"`
	Title                    client.LocalizedText     `xorm:"varchar(255) 'title'"`
	Description              client.LocalizedText     `xorm:"varchar(255) 'description'"`
	ScenarioScriptAssetPath  client.AdvScriptStruktur `xorm:"varchar(255) 'scenario_script_asset_path'"`
}

func (es *EventStory) populate(gamedata *Gamedata) {
	es.Title.DotUnderText = gamedata.Dictionary.Resolve(es.Title.DotUnderText)
	es.Description.DotUnderText = gamedata.Dictionary.Resolve(es.Description.DotUnderText)
}

func (es *EventStory) GetEventMarathonStory() client.EventMarathonStory {
	// initially  this set of point was used, but then it was readjusted and kept to the new one
	// https://pasuggun.tistory.com/entry/%EC%8A%A4%EC%BF%A0%EC%8A%A4%ED%83%80-01-%EC%8B%9C%ED%81%AC%EB%A6%BF-%ED%8C%8C%ED%8B%B0-%ED%98%B8%EB%85%B8%EC%B9%B4-%EC%B9%98%EC%B9%B4-%EC%95%84%EC%9C%A0%EB%AC%B4
	// requiredEventPointMarathon := []int32{-1, 0, 300, 2100, 4500, 8000, 14000, 25000}
	// https://pasuggun.tistory.com/entry/%EC%8A%A4%EC%BF%A0%EC%8A%A4%ED%83%80-127-%EC%97%AC%EB%A6%84%EB%B0%A4%EC%97%90-%EC%8A%A4%EB%A9%B0%EB%93%9C%EB%8A%94-%EB%A7%88%EC%9D%8C-%EC%97%90%EB%A6%AC-%EC%97%A0%EB%A7%88-%EC%B9%B4%EB%82%9C
	// https://pasuggun.tistory.com/entry/%EC%8A%A4%EC%BF%A0%EC%8A%A4%ED%83%80-78-%ED%8A%B9%EB%B3%84%ED%95%9C-%EC%97%AC%EB%A6%84-%ED%8E%98%EC%8A%A4%ED%8B%B0%EB%B2%8C-%EC%9A%94%EC%9A%B0-%EC%9A%B0%EB%AF%B8-%EB%A6%AC%EC%BD%94
	requiredEventPointMarathon := []int32{-1, 0, 1000, 3000, 6000, 10000, 16000, 21000}
	// 15000 and then 21000 would have make more sense as a progression but whatever

	return client.EventMarathonStory{
		EventMarathonStoryId:     es.StoryEventId,
		StoryNumber:              es.StoryNumber,
		IsPrologue:               false, // always false
		RequiredEventPoint:       requiredEventPointMarathon[es.StoryNumber],
		StoryBannerThumbnailPath: es.StoryBannerThumbnailPath,
		StoryDetailThumbnailPath: es.StoryDetailThumbnailPath,
		Title:                    es.Title,
		Description:              es.Description,
		ScenarioScriptAssetPath:  es.ScenarioScriptAssetPath,
	}
}

func (es *EventStory) GetEventMiningStory() client.EventMiningStory {
	// https://pasuggun.tistory.com/entry/%EC%8A%A4%EC%BF%A0%EC%8A%A4%ED%83%80-07-%EB%B0%94%EB%8B%A4-%EC%9C%84-%EB%8C%80%EC%97%B4%EC%A0%84-%EC%9A%B0%EB%AF%B8-%EB%A7%88%ED%82%A4-%ED%95%98%EB%82%98%EB%A7%88%EB%A3%A8
	// https://pasuggun.tistory.com/entry/%EC%8A%A4%EC%BF%A0%EC%8A%A4%ED%83%80-161-%EC%8A%A4%EC%9C%84%ED%8A%B8-%EB%93%9C%EB%A6%AC%EB%B0%8D%E2%99%AA-%EB%8B%88%EC%BD%94-%EB%A6%B0-%EC%8B%9C%EC%A6%88%EC%BF%A0
	// baring exception when the WW server are catching up to the JP one, this is the standard amount
	requiredEventPointMining := []int32{-1, 0, 1500, 5000, 10000, 15500, 21500, 27500}
	return client.EventMiningStory{
		EventMiningStoryId:       es.StoryEventId,
		StoryNumber:              es.StoryNumber,
		IsPrologue:               false, // always false
		RequiredEventPoint:       requiredEventPointMining[es.StoryNumber],
		StoryBannerThumbnailPath: es.StoryBannerThumbnailPath,
		StoryDetailThumbnailPath: es.StoryDetailThumbnailPath,
		Title:                    es.Title,
		Description:              es.Description,
		ScenarioScriptAssetPath:  es.ScenarioScriptAssetPath,
	}
}

func loadEventStory(gamedata *Gamedata) {
	gamedata.EventStory = map[int32]*EventStory{}
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_story_event_history_detail").Find(&gamedata.EventStory)
	})
	utils.CheckErr(err)
	for _, story := range gamedata.EventStory {
		story.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadEventStory)
}
