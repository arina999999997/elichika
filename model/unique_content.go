// content that can only be obtained once like story, voice, emblem, ...
// basically everything that is unlocked once then you have access to it from then on, but you can't modify it except for some small flag
// listed in the order of appearing in user model
package model

var (
	TableNameToInterface map[string]interface{}
)

type UserStoryMain struct {
	UserId            int `xorm:"pk 'user_id'" json:"-"`
	StoryMainMasterId int `xorm:"pk 'story_main_master_id'" json:"story_main_master_id"`
}

func (usm *UserStoryMain) Id() int64 {
	return int64(usm.StoryMainMasterId)
}

// for that one part in the main story where you select an idol
type UserStoryMainSelected struct {
	UserId          int `xorm:"pk 'user_id'" json:"-"`
	StoryMainCellId int `xorm:"pk 'story_main_cell_id'" json:"story_main_cell_id"`
	SelectedId      int `xorm:"'selected_id'" json:"selected_id"`
}

func (usms *UserStoryMainSelected) Id() int64 {
	return int64(usms.StoryMainCellId)
}

type UserVoice struct {
	UserId            int  `xorm:"pk 'user_id'" json:"-"`
	NaviVoiceMasterId int  `xorm:"pk 'navi_voice_master_id'" json:"navi_voice_master_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
}

func (uv *UserVoice) Id() int64 {
	return int64(uv.NaviVoiceMasterId)
}

type UserEmblem struct {
	UserId      int     `xorm:"pk 'user_id'" json:"-"`
	EmblemMId   int     `xorm:"pk 'emblem_m_id'" json:"emblem_m_id"`
	IsNew       bool    `xorm:"'is_new'" json:"is_new"`
	EmblemParam *string `xorm:"'emblem_param'" json:"emblem_param"`
	AcquiredAt  int64   `xorm:"'acquired_at'" json:"acquired_at"`
}

func (ue *UserEmblem) Id() int64 {
	return int64(ue.EmblemMId)
}

type UserCustomBackground struct {
	UserId                   int  `xorm:"pk 'user_id'" json:"-"`
	CustomBackgroundMasterId int  `xorm:"pk 'custom_background_master_id'" json:"custom_background_master_id"`
	IsNew                    bool `xorm:"'is_new'" json:"is_new"`
}

func (ucb *UserCustomBackground) Id() int64 {
	return int64(ucb.CustomBackgroundMasterId)
}

type UserStorySide struct {
	UserId            int   `xorm:"pk 'user_id'" json:"-"`
	StorySideMasterId int   `xorm:"pk 'story_side_master_id'" json:"story_side_master_id"`
	IsNew             bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt        int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (uss *UserStorySide) Id() int64 {
	return int64(uss.StorySideMasterId)
}

type UserStoryMember struct {
	UserId              int   `xorm:"pk 'user_id'" json:"-"`
	StoryMemberMasterId int   `xorm:"pk 'story_member_master_id'" json:"story_member_master_id"`
	IsNew               bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt          int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (usm *UserStoryMember) Id() int64 {
	return int64(usm.StoryMemberMasterId)
}

type UserStoryEventHistory struct {
	UserId       int `xorm:"pk 'user_id'" json:"-"`
	StoryEventId int `xorm:"pk 'story_event_id'" json:"story_event_id"`
}

func (useh *UserStoryEventHistory) Id() int64 {
	return int64(useh.StoryEventId)
}

type UserUnlockScene struct {
	UserId          int `xorm:"pk 'user_id'" json:"-"`
	UnlockSceneType int `xorm:"pk 'unlock_scene_type'" json:"unlock_scene_type"`
	Status          int `xorm:"'status'" json:"status"`
}

func (uus *UserUnlockScene) Id() int64 {
	return int64(uus.UnlockSceneType)
}

type UserSceneTips struct {
	UserId        int `xorm:"pk 'user_id'" json:"-"`
	SceneTipsType int `xorm:"pk 'scene_tips_type'" json:"scene_tips_type"`
}

func (ust *UserSceneTips) Id() int64 {
	return int64(ust.SceneTipsType)
}

type UserRuleDescription struct {
	UserId            int `xorm:"pk 'user_id'" json:"-"`
	RuleDescriptionId int `xorm:"pk 'rule_description_id'" json:"-"`
	DisplayStatus     int `xorm:"'display_status'" json:"display_status"`
}

func (urd *UserRuleDescription) Id() int64 {
	return int64(urd.RuleDescriptionId)
}
func (urd *UserRuleDescription) SetId(id int64) {
	urd.RuleDescriptionId = int(id)
}

type UserReferenceBook struct {
	UserId          int `xorm:"pk 'user_id'" json:"-"`
	ReferenceBookId int `xorm:"pk 'reference_book_id'" json:"reference_book_id"`
}

func (urb *UserReferenceBook) Id() int64 {
	return int64(urb.ReferenceBookId)
}

type UserStoryLinkage struct {
	UserId                   int `xorm:"pk 'user_id'" json:"-"`
	StoryLinkageCellMasterId int `xorm:"pk 'story_linkage_cell_master_id'" json:"story_linkage_cell_master_id"`
}

func (usl *UserStoryLinkage) Id() int64 {
	return int64(usl.StoryLinkageCellMasterId)
}

type UserStoryMainPartDigestMovie struct {
	UserId                int `xorm:"pk 'user_id'" json:"-"`
	StoryMainPartMasterId int `xorm:"pk 'story_main_part_master_id'" json:"story_main_part_master_id"`
}

func (usmpdm *UserStoryMainPartDigestMovie) Id() int64 {
	return int64(usmpdm.StoryMainPartMasterId)
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
