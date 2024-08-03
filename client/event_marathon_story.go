package client

type EventMarathonStory struct {
	EventMarathonStoryId     int32             `json:"event_marathon_story_id"`
	StoryNumber              int32             `json:"story_number"`
	IsPrologue               bool              `json:"is_prologue"`
	RequiredEventPoint       int32             `json:"required_event_point"`
	StoryBannerThumbnailPath TextureStruktur   `json:"story_banner_thumbnail_path"`
	StoryDetailThumbnailPath TextureStruktur   `json:"story_detail_thumbnail_path"`
	Title                    LocalizedText     `json:"title"`
	Description              LocalizedText     `json:"description"`
	ScenarioScriptAssetPath  AdvScriptStruktur `json:"scenario_script_asset_path"`
}
