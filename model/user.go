package model

import (
	"elichika/client"
	"elichika/generic"
)

// this is not stored, constructed from main db
// partially loaded from u_info, then load from u_card
type UserBasicInfo struct {
	UserId int `xorm:"pk 'user_id'" json:"user_id"`
	Name   struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // player name
	Rank                  int   `json:"rank"` // rank
	LastPlayedAt          int64 `xorm:"'last_login_at'" json:"last_played_at"`
	RecommendCardMasterId int32 `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"` // featured / partner card

	RecommendCardLevel                  int  `xorm:"-" json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool `xorm:"-" json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool `xorm:"-" json:"is_recommend_card_all_training_activated"`

	EmblemId            int  `xorm:"'emblem_id' "json:"emblem_id"` // title
	IsNew               bool `xorm:"-" json:"is_new"`              // not sure what this thing is about, maybe new friend?
	IntroductionMessage struct {
		DotUnderText string `xorm:"message" json:"dot_under_text"`
	} `xorm:"extends" json:"introduction_message"` // introduction message
	FriendApprovedAt *int64 `xorm:"-" json:"friend_approved_at"`
	RequestStatus    int    `xorm:"-" json:"request_status"`
	IsRequestPending bool   `xorm:"-" json:"is_request_pending"`
}

type UserProfileLiveStats struct {
	LivePlayCount  [5]int `xorm:"'live_play_count'"`
	LiveClearCount [5]int `xorm:"'live_clear_count'"`
}

type UserProfileInfo struct {
	BasicInfo      UserBasicInfo `xorm:"extends" json:"basic_info"`
	TotalLovePoint int           `xorm:"-" json:"total_love_point"`
	LoveMembers    [3]struct {
		MemberMasterId int `json:"member_master_id"`
		LovePoint      int `json:"love_point"`
	} `xorm:"-" json:"love_members"`
	MemberGuildMemberMasterId int `xorm:"'member_guild_member_master_id'" json:"member_guild_member_master_id"`
}

type LivePartnerCard struct {
	LivePartnerCategoryMasterId int             `json:"live_partner_category_master_id"`
	PartnerCard                 PartnerCardInfo `json:"partner_card"`
}

type Profile struct {
	ProfileInfo UserProfileInfo `xorm:"extends" json:"profile_info"`
	GuestInfo   struct {
		LivePartnersCards []LivePartnerCard `json:"live_partner_cards"`
	} `xorm:"-" json:"guest_info"`
	PlayInfo struct {
		LivePlayCount          []int          `xorm:"-" json:"live_play_count"`
		LiveClearCount         []int          `xorm:"-" json:"live_clear_count"`
		JoinedLiveCardRanking  []CardPlayInfo `xorm:"-" json:"joined_live_card_ranking"`
		PlaySkillCardRanking   []CardPlayInfo `xorm:"-" json:"play_skill_card_ranking"`
		MaxScoreLiveDifficulty struct {
			LiveDifficultyMasterId int32 `xorm:"'max_score_live_difficulty_master_id'" json:"live_difficulty_master_id"`
			Score                  int32 `xorm:"'live_max_score'" json:"score"`
		} `xorm:"extends" json:"max_score_live_difficulty"`
		MaxComboLiveDifficulty struct {
			LiveDifficultyMasterId int32 `xorm:"'max_combo_live_difficulty_master_id'" json:"live_difficulty_master_id"`
			Score                  int32 `xorm:"'live_max_combo'" json:"score"`
		} `xorm:"extends" json:"max_combo_live_difficulty"`
	} `xorm:"extends" json:"play_info"`
	MemberInfo struct {
		UserMembers []MemberPublicInfo `json:"user_members"`
	} `xorm:"-" json:"member_info"`
}

func init() {

	type DbUser struct {
		client.UserStatus    `xorm:"extends"`
		UserProfileLiveStats `xorm:"extends"`
	}
	TableNameToInterface["u_info"] = generic.UserIdWrapper[DbUser]{}

}
