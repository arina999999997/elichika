package client

import (
	"elichika/generic"
)

type OtherUser struct {
	UserId                              int32                   `json:"user_id"`
	Name                                LocalizedText           `json:"name"`
	Rank                                int32                   `json:"rank"`
	LastPlayedAt                        int64                   `json:"last_played_at"`
	RecommendCardMasterId               int32                   `json:"recommend_card_master_id"`
	RecommendCardLevel                  int32                   `json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool                    `json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool                    `json:"is_recommend_card_all_training_activated"`
	EmblemId                            int32                   `json:"emblem_id"`
	IsNew                               bool                    `json:"is_new"`
	IntroductionMessage                 LocalizedText           `json:"introduction_message"`
	FriendApprovedAt                    generic.Nullable[int64] `json:"friend_approved_at"`
	RequestStatus                       int32                   `json:"request_status" enum:""`
	IsRequestPending                    bool                    `json:"is_request_pending"`
}
