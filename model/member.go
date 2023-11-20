package model

// UserCommunicationMemberDetailBadgeByID ...
type UserCommunicationMemberDetailBadgeByID struct {
	MemberMasterID     int  `json:"member_master_id"`
	IsStoryMemberBadge bool `json:"is_story_member_badge"`
	IsStorySideBadge   bool `json:"is_story_side_badge"`
	IsVoiceBadge       bool `json:"is_voice_badge"`
	IsThemeBadge       bool `json:"is_theme_badge"`
	IsCardBadge        bool `json:"is_card_badge"`
	IsMusicBadge       bool `json:"is_music_badge"`
}

// UserMember ...
type UserMember struct {
	UserID                   int  `xorm:"pk 'user_id'" json:"-"`
	MemberMasterID           int  `xorm:"pk 'member_master_id'" json:"member_master_id"`
	CustomBackgroundMasterID int  `xorm:"'custom_background_master_id'" json:"custom_background_master_id"`
	SuitMasterID             int  `xorm:"'suit_master_id'" json:"suit_master_id"`
	LovePoint                int  `json:"love_point"`
	LovePointLimit           int  `json:"love_point_limit"`
	LoveLevel                int  `json:"love_level"`
	ViewStatus               int  `json:"view_status"`
	IsNew                    bool `json:"is_new"`
	OwnedCardCount           int  `json:"-"`
	AllTrainingCardCount     int  `json:"-"`
}

func (um *UserMember) ID() int64 {
	return int64(um.MemberMasterID)
}

type MemberPublicInfo struct {
	MemberMasterID       int `xorm:"pk 'member_master_id'" json:"member_master_id"`
	LoveLevel            int `json:"love_level"`
	LovePointLimit       int `json:"love_point_limit"`
	OwnedCardCount       int `json:"owned_card_count"`
	AllTrainingCardCount int `json:"all_training_activated_card_count"`
}

// Bond board (Love Panel)
type UserMemberLovePanel struct {
	// love panel level = m_member_love_panel[id] / 1000
	// SELECT * FROM m_member_love_panel_cell WHERE id != (member_love_panel_master_id / 1000 - 1) * 10000 + (panel_index + 1) * 1000 + (member_love_panel_master_id % 1000); -> 0
	UserID                    int   `xorm:"pk <- 'user_id'" json:"-"`
	MemberID                  int   `xorm:"pk <- 'member_master_id'" json:"member_id"` // member_love_panel_master_id % 1000
	MemberLovePanelCellIDs    []int `xorm:"-" json:"member_love_panel_cell_ids"`
	LovePanelLevel            int   `json:"-"` // member_love_panel_master_id / 1000
	LovePanelLastLevelCellIDs []int `xorm:"'love_panel_last_level_cell_ids'" json:"-"`
	// there is no ambiguous representation
	// - When the last level is filled, if there is a next level then LovePanelLevel is increased, and LovePanelLastLevelCellIDs is cleared
	// - otherwise, LovePanelLevel stay the same, and LovePanelLastLevelCellIDs has 5 tiles.
}

func (x *UserMemberLovePanel) Fill() {
	x.MemberLovePanelCellIDs = []int{} // [] instead of null
	for l := 1; l < x.LovePanelLevel; l++ {
		for cell := 1000; cell <= 5000; cell += 1000 {
			x.MemberLovePanelCellIDs = append(x.MemberLovePanelCellIDs, (l-1)*10000+cell+x.MemberID)
		}
	}
	x.MemberLovePanelCellIDs = append(x.MemberLovePanelCellIDs, x.LovePanelLastLevelCellIDs...)
}

func (x *UserMemberLovePanel) LevelUp() {
	if len(x.LovePanelLastLevelCellIDs) != 5 {
		panic("incorrect level up")
	}
	x.LovePanelLastLevelCellIDs = []int{}
	x.LovePanelLevel++
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	type DbMember struct {
		UserMember          `xorm:"extends"`
		LovePanelLevel            int   `xorm:"'love_panel_level' default 1"`
		LovePanelLastLevelCellIds []int `xorm:"'love_panel_last_level_cell_ids' default '[]'"`
	}
	TableNameToInterface["u_member"] = DbMember{}
}