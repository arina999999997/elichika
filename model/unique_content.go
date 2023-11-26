// content that can only be obtained once like story, voice, emblem, ...
// basically everything that is unlocked once then you have access to it from then on, but you can't modify it except for some small flag
// listed in the order of appearing in user model
package model

var (
	TableNameToInterface map[string]interface{}
)

type UserStoryMain struct {
	UserID            int `xorm:"pk 'user_id'" json:"-"`
	StoryMainMasterID int `xorm:"pk 'story_main_master_id'" json:"story_main_master_id"`
}

func (usm *UserStoryMain) ID() int64 {
	return int64(usm.StoryMainMasterID)
}

// for that one part in the main story where you select an idol
type UserStoryMainSelected struct {
	UserID          int `xorm:"pk 'user_id'" json:"-"`
	StoryMainCellID int `xorm:"pk 'story_main_cell_id'" json:"story_main_cell_id"`
	SelectedID      int `xorm:"'selected_id'" json:"selected_id"`
}

func (usms *UserStoryMainSelected) ID() int64 {
	return int64(usms.StoryMainCellID)
}

type UserVoice struct {
	UserID            int  `xorm:"pk 'user_id'" json:"-"`
	NaviVoiceMasterID int  `xorm:"pk 'navi_voice_master_id'" json:"navi_voice_master_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
}

func (uv *UserVoice) ID() int64 {
	return int64(uv.NaviVoiceMasterID)
}

type UserEmblem struct {
	UserID      int     `xorm:"pk 'user_id'" json:"-"`
	EmblemMID   int     `xorm:"pk 'emblem_m_id'" json:"emblem_m_id"`
	IsNew       bool    `xorm:"'is_new'" json:"is_new"`
	EmblemParam *string `xorm:"'emblem_param'" json:"emblem_param"`
	AcquiredAt  int64   `xorm:"'acquired_at'" json:"acquired_at"`
}

func (ue *UserEmblem) ID() int64 {
	return int64(ue.EmblemMID)
}

type UserCustomBackground struct {
	UserID                   int  `xorm:"pk 'user_id'" json:"-"`
	CustomBackgroundMasterID int  `xorm:"pk 'custom_background_master_id'" json:"custom_background_master_id"`
	IsNew                    bool `xorm:"'is_new'" json:"is_new"`
}

func (ucb *UserCustomBackground) ID() int64 {
	return int64(ucb.CustomBackgroundMasterID)
}

type UserStorySide struct {
	UserID            int   `xorm:"pk 'user_id'" json:"-"`
	StorySideMasterID int   `xorm:"pk 'story_side_master_id'" json:"story_side_master_id"`
	IsNew             bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt        int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (uss *UserStorySide) ID() int64 {
	return int64(uss.StorySideMasterID)
}

type UserStoryMember struct {
	UserID              int   `xorm:"pk 'user_id'" json:"-"`
	StoryMemberMasterID int   `xorm:"pk 'story_member_master_id'" json:"story_member_master_id"`
	IsNew               bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt          int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (usm *UserStoryMember) ID() int64 {
	return int64(usm.StoryMemberMasterID)
}

type UserStoryEventHistory struct {
	UserID       int `xorm:"pk 'user_id'" json:"-"`
	StoryEventID int `xorm:"pk 'story_event_id'" json:"story_event_id"`
}

func (useh *UserStoryEventHistory) ID() int64 {
	return int64(useh.StoryEventID)
}

type UserUnlockScene struct {
	UserID          int `xorm:"pk 'user_id'" json:"-"`
	UnlockSceneType int `xorm:"pk 'unlock_scene_type'" json:"unlock_scene_type"`
	Status          int `xorm:"'status'" json:"status"`
}

func (uus *UserUnlockScene) ID() int64 {
	return int64(uus.UnlockSceneType)
}

type UserSceneTips struct {
	UserID        int `xorm:"pk 'user_id'" json:"-"`
	SceneTipsType int `xorm:"pk 'scene_tips_type'" json:"scene_tips_type"`
}

func (ust *UserSceneTips) ID() int64 {
	return int64(ust.SceneTipsType)
}

type UserRuleDescription struct {
	UserID            int `xorm:"pk 'user_id'" json:"-"`
	RuleDescriptionID int `xorm:"pk 'rule_description_id'" json:"-"`
	DisplayStatus     int `xorm:"'display_status'" json:"display_status"`
}

func (urd *UserRuleDescription) ID() int64 {
	return int64(urd.RuleDescriptionID)
}
func (urd *UserRuleDescription) SetID(id int64) {
	urd.RuleDescriptionID = int(id)
}

type UserReferenceBook struct {
	UserID          int `xorm:"pk 'user_id'" json:"-"`
	ReferenceBookID int `xorm:"pk 'reference_book_id'" json:"reference_book_id"`
}

func (urb *UserReferenceBook) ID() int64 {
	return int64(urb.ReferenceBookID)
}

type UserStoryLinkage struct {
	UserID                   int `xorm:"pk 'user_id'" json:"-"`
	StoryLinkageCellMasterID int `xorm:"pk 'story_linkage_cell_master_id'" json:"story_linkage_cell_master_id"`
}

func (usl *UserStoryLinkage) ID() int64 {
	return int64(usl.StoryLinkageCellMasterID)
}

type UserStoryMainPartDigestMovie struct {
	UserID                int `xorm:"pk 'user_id'" json:"-"`
	StoryMainPartMasterID int `xorm:"pk 'story_main_part_master_id'" json:"story_main_part_master_id"`
}

func (usmpdm *UserStoryMainPartDigestMovie) ID() int64 {
	return int64(usmpdm.StoryMainPartMasterID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_story_main"] = UserStoryMain{}
	TableNameToInterface["u_story_main_selected"] = UserStoryMainSelected{}
	TableNameToInterface["u_voice"] = UserVoice{}
	TableNameToInterface["u_emblem"] = UserEmblem{}
	TableNameToInterface["u_custom_background"] = UserCustomBackground{}
	TableNameToInterface["u_story_side"] = UserStorySide{}
	TableNameToInterface["u_story_member"] = UserStoryMember{}
	TableNameToInterface["u_story_event_history"] = UserStoryEventHistory{}
	TableNameToInterface["u_unlock_scene"] = UserUnlockScene{}
	TableNameToInterface["u_scene_tips"] = UserSceneTips{}
	TableNameToInterface["u_rule_description"] = UserRuleDescription{}
	TableNameToInterface["u_reference_book"] = UserReferenceBook{}
	TableNameToInterface["u_story_linkage"] = UserStoryLinkage{}
	TableNameToInterface["u_story_main_part_digest_movie"] = UserStoryMainPartDigestMovie{}
}
