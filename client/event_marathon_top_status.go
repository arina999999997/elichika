package client

import (
	"elichika/generic"
)

type EventMarathonTopStatus struct {
	EventId                            int32                                           `json:"event_id"`
	IsFirstAccess                      bool                                            `json:"is_first_access"`
	StartAt                            int64                                           `json:"start_at"`
	EndAt                              int64                                           `json:"end_at"`
	ResultAt                           int64                                           `json:"result_at"`
	ExpiredAt                          int64                                           `json:"expired_at"`
	TitleImagePath                     TextureStruktur                                 `json:"title_image_path"`
	BackgroundImagePath                TextureStruktur                                 `json:"background_image_path"`
	UserRankingStatus                  EventMarathonUserRanking                        `json:"user_ranking_status"`
	BoardStatus                        EventMarathonBoard                              `json:"board_status"`
	StoryStatus                        EventMarathonStoryStatus                        `json:"story_status"`
	EventTotalTopicRewardInfo          generic.List[EventTopicReward]                  `json:"event_total_topic_reward_info"`
	EventRankingTopicRewardInfo        generic.List[EventTopicReward]                  `json:"event_ranking_topic_reward_info"`
	EventMarathonPointRewardMasterRows generic.List[EventMarathonPointRewardMasterRow] `json:"event_marathon_point_reward_master_rows"`
	// unused
	MaxLoopCount int32 `json:"max_loop_count"`
	// unused
	EventMarathonPointLoopRewardMasterRows      generic.List[EventMarathonPointRewardMasterRow]          `json:"event_marathon_point_loop_reward_master_rows"`
	EventMarathonRankingRewardMasterRows        generic.List[EventMarathonRankingRewardMasterRow]        `json:"event_marathon_ranking_reward_master_rows"`
	EventMarathonRewardMasterRows               generic.List[EventMarathonRewardMasterRow]               `json:"event_marathon_reward_master_rows"`
	EventMarathonRuleDescriptionPageMasterRows  generic.List[EventMarathonRuleDescriptionPageMasterRow]  `json:"event_marathon_rule_description_page_master_rows"`
	BgmAssetPath                                SoundStruktur                                            `json:"bgm_asset_path"`
	EventMarathonBonusPopupOrderCardMaterRows   generic.List[EventMarathonBonusPopupOrderCardMaterRow]   `json:"event_marathon_bonus_popup_order_card_mater_rows"`
	EventMarathonBonusPopupOrderMemberMaterRows generic.List[EventMarathonBonusPopupOrderMemberMaterRow] `json:"event_marathon_bonus_popup_order_member_mater_rows"`
	GachaMasterId                               int32                                                    `json:"gacha_master_id"`
}
