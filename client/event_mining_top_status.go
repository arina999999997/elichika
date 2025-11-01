package client

import (
	"elichika/generic"
)

type EventMiningTopStatus struct {
	EventId                                   int32                                                 `json:"event_id"`
	StartAt                                   int64                                                 `json:"start_at"`
	EndAt                                     int64                                                 `json:"end_at"`
	ResultAt                                  int64                                                 `json:"result_at"`
	ExpiredAt                                 int64                                                 `json:"expired_at"`
	TitleImagePath                            TextureStruktur                                       `json:"title_image_path"`
	UserRankingStatus                         EventMiningUserRanking                                `json:"user_ranking_status"`
	StoryStatus                               EventMiningStoryStatus                                `json:"story_status"`
	EventMiningPointRankingRewardMasterRows   generic.List[EventMiningRankingRewardMasterRow]       `json:"event_mining_point_ranking_reward_master_rows"`
	EventMiningVoltageRankingRewardMasterRows generic.List[EventMiningRankingRewardMasterRow]       `json:"event_mining_voltage_ranking_reward_master_rows"`
	EventMiningRewardMasterRows               generic.List[EventMiningRewardMasterRow]              `json:"event_mining_reward_master_rows"`
	EventMiningRuleDescriptionPageMasterRows  generic.List[EventMiningRuleDescriptionPageMasterRow] `json:"event_mining_rule_description_page_master_rows"`
	EventPointRankingTopicRewardInfo          generic.List[EventMiningTopicReward]                  `json:"event_point_ranking_topic_reward_info"`
	EventVoltageRankingTopicRewardInfo        generic.List[EventMiningTopicReward]                  `json:"event_voltage_ranking_topic_reward_info"`
	BackgroundImagePath                       TextureStruktur                                       `json:"background_image_path"`
	BgmAssetPath                              SoundStruktur                                         `json:"bgm_asset_path"`
	EventMiningTopCellStateList               generic.List[EventMiningTopCellState]                 `json:"event_mining_top_cell_state_list"`
	EventMiningTopLikeMemerRows               generic.List[EventMiningTopLikeMemerRow]              `json:"event_mining_top_like_memer_rows"`
	EventMiningTopStillCellMasterRows         generic.List[EventMiningTopStillCellMasterRow]        `json:"event_mining_top_still_cell_master_rows"`
	EventMiningTopStillSubCellMasterRows      generic.List[EventMiningTopStillSubCellMasterRow]     `json:"event_mining_top_still_sub_cell_master_rows"`
	EventMiningTopAnimationCellMasterRows     generic.List[EventMiningTopAnimationCellMasterRow]    `json:"event_mining_top_animation_cell_master_rows"`
	EventMiningBonusPopupOrderCardMaterRows   generic.List[EventMiningBonusPopupOrderCardMaterRow]  `json:"event_mining_bonus_popup_order_card_mater_rows"`
	GachaMasterId                             int32                                                 `json:"gacha_master_id"`
	EventPointMasterId                        int32                                                 `json:"event_point_master_id"`
	IsAddNewPanel                             bool                                                  `json:"is_add_new_panel"`
	EventMiningCompetitionMasterRows          generic.List[EventMiningCompetitionMasterRow]         `json:"event_mining_competition_master_rows"`
	EventMiningMusicBestScoreList             generic.List[EventMiningMusicBestScore]               `json:"event_mining_music_best_score_list"`
	SelectionAmount                           int32                                                 `json:"selection_amount"`
	TradeMasterId                             generic.Nullable[int32]                               `json:"trade_master_id"`
}
