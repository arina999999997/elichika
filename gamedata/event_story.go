package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
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
	//   - this is the numbers for the first event, as found in recordings on youtube.
	//
	RequiredEventPoint       int32                    `xorm:"-"`
	StoryBannerThumbnailPath client.TextureStruktur   `xorm:"varchar(255) 'banner_thumbnail_path'"`
	StoryDetailThumbnailPath client.TextureStruktur   `xorm:"varchar(255) 'detail_thumbnail_path'"`
	Title                    client.LocalizedText     `xorm:"varchar(255) 'title'"`
	Description              client.LocalizedText     `xorm:"varchar(255) 'description'"`
	ScenarioScriptAssetPath  client.AdvScriptStruktur `xorm:"varchar(255) 'scenario_script_asset_path'"`
}

func (es *EventStory) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	es.Title.DotUnderText = dictionary.Resolve(es.Title.DotUnderText)
	es.Description.DotUnderText = dictionary.Resolve(es.Description.DotUnderText)
}

func (es *EventStory) GetEventMarathonStory() client.EventMarathonStory {
	requiredEventPointMarathon := []int32{-1, 0, 300, 2100, 4500, 8000, 14000, 25000}
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

func loadEventStory(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.EventStory = map[int32]*EventStory{}
	err := masterdata_db.Table("m_story_event_history_detail").Find(&gamedata.EventStory)
	utils.CheckErr(err)
	for _, story := range gamedata.EventStory {
		story.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadEventStory)
}
