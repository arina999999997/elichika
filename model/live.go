package model

type UserLiveDeck struct {
	UserID         int `xorm:"pk 'user_id'" json:"-"`
	UserLiveDeckID int `xorm:"pk 'user_live_deck_id'" json:"user_live_deck_id"`
	Name           struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // deck name
	CardMasterID1 int `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterID2 int `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterID3 int `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	CardMasterID4 int `xorm:"'card_master_id_4'" json:"card_master_id_4"`
	CardMasterID5 int `xorm:"'card_master_id_5'" json:"card_master_id_5"`
	CardMasterID6 int `xorm:"'card_master_id_6'" json:"card_master_id_6"`
	CardMasterID7 int `xorm:"'card_master_id_7'" json:"card_master_id_7"`
	CardMasterID8 int `xorm:"'card_master_id_8'" json:"card_master_id_8"`
	CardMasterID9 int `xorm:"'card_master_id_9'" json:"card_master_id_9"`
	SuitMasterID1 int `xorm:"'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterID2 int `xorm:"'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterID3 int `xorm:"'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterID4 int `xorm:"'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterID5 int `xorm:"'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterID6 int `xorm:"'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterID7 int `xorm:"'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterID8 int `xorm:"'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterID9 int `xorm:"'suit_master_id_9'" json:"suit_master_id_9"`
}

func (uld *UserLiveDeck) ID() int64 {
	return int64(uld.UserLiveDeckID)
}

type UserLiveParty struct {
	UserID         int `xorm:"pk 'user_id'" json:"-"`
	PartyID        int `xorm:"pk 'party_id'" json:"party_id"`
	UserLiveDeckID int `xorm:"'user_live_deck_id'" json:"user_live_deck_id"`
	Name           struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // deck name
	IconMasterID     int    `xorm:"'icon_master_id'" json:"icon_master_id"`
	CardMasterID1    int    `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterID2    int    `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterID3    int    `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	UserAccessoryID1 *int64 `xorm:"'user_accessory_id_1'" json:"user_accessory_id_1"` // null for empty
	UserAccessoryID2 *int64 `xorm:"'user_accessory_id_2'" json:"user_accessory_id_2"`
	UserAccessoryID3 *int64 `xorm:"'user_accessory_id_3'" json:"user_accessory_id_3"`
}

func (uld *UserLiveParty) ID() int64 {
	return int64(uld.PartyID)
}

// PartyName ...
type PartyName struct {
	DotUnderText string `json:"dot_under_text"`
}

// DeckSquadDict ...
type DeckSquadDict struct {
	CardMasterIDs    []int    `json:"card_master_ids"`
	UserAccessoryIDs []*int64 `json:"user_accessory_ids"`
}

// LiveDaily ...
type LiveDaily struct {
	LiveDailyMasterID      int `json:"live_daily_master_id" xorm:"id"`
	LiveMasterID           int `json:"live_master_id" xorm:"live_id"`
	EndAt                  int `json:"end_at"`
	RemainingPlayCount     int `json:"remaining_play_count"`
	RemainingRecoveryCount int `json:"remaining_recovery_count"`
}

// LiveStartReq ...
type LiveStartReq struct {
	LiveDifficultyID    int  `json:"live_difficulty_id"`
	DeckID              int  `json:"deck_id"`
	CellID              *int `json:"cell_id"`
	PartnerUserID       int  `json:"partner_user_id"`
	PartnerCardMasterID int  `json:"partner_card_master_id"`
	LpMagnification     int  `json:"lp_magnification"`
	IsAutoPlay          bool `json:"is_auto_play"`
	IsReferenceBook     bool `json:"is_reference_book"`
}

// LivePartnerInfo ...
type LivePartnerInfo UserBasicInfo

// guests before live start
type LiveStartLivePartner struct {
	UserID int `xorm:"'user_id' "json:"user_id"`
	Name   struct {
		DotUnderText string `xorm:"'name'" json:"dot_under_text"`
	} `xorm:"extends" json:"name"`
	Rank                int   `json:"rank"`
	LastLoginAt         int64 `json:"last_login_at"`
	CardByCategory      []any `xorm:"-" json:"card_by_category"`
	EmblemID            int   `xorm:"'emblem_id'" json:"emblem_id"`
	IsFriend            bool  `xorm:"-" json:"is_friend"`
	IntroductionMessage struct {
		DotUnderText string `xorm:"'message'" json:"dot_under_text"`
	} `xorm:"extends" json:"introduction_message"`
}

// the state of the song user is playing
// sent to user in /live/Start
// stored necessary info to recover full state in db
// each user can only have 1 live state stored in db
type LiveState struct {
	UserID    int   `xorm:"pk 'user_id'" json:"-"`
	LiveID    int64 `xorm:"'live_id'" json:"live_id"`
	LiveType  int   `json:"live_type"`
	DeckID    int   `xorm:"-" json:"deck_id"` // get from user status
	LiveStage struct {
		LiveDifficultyID int `json:"live_difficulty_id"` // get from user status
		// get from song db
		LiveNotes        []LiveNote        `json:"live_notes"`
		LiveWaveSettings []LiveWaveSetting `json:"live_wave_settings"`
		NoteGimmicks     []NoteGimmick     `json:"note_gimmicks"`
		StageGimmickDict []any             `json:"stage_gimmick_dict"`
	} `xorm:"-" json:"live_stage"`
	PartnerUserID   int             `xorm:"partner_user_id" json:"-"`
	LivePartnerCard PartnerCardInfo `xorm:"extends" json:"live_partner_card"`
	IsPartnerFriend bool            `json:"is_partner_friend"`
	CellID          *int            `xorm:"'cell_id' "json:"cell_id"`
	TowerLive       *int            `json:"tower_live"`
}

type UserLiveMvDeck struct {
	LiveMasterID     int  `json:"live_master_id"`
	MemberMasterID1  *int `json:"member_master_id_1"`
	MemberMasterID2  *int `json:"member_master_id_2"`
	MemberMasterID3  *int `json:"member_master_id_3"`
	MemberMasterID4  *int `json:"member_master_id_4"`
	MemberMasterID5  *int `json:"member_master_id_5"`
	MemberMasterID6  *int `json:"member_master_id_6"`
	MemberMasterID7  *int `json:"member_master_id_7"`
	MemberMasterID8  *int `json:"member_master_id_8"`
	MemberMasterID9  *int `json:"member_master_id_9"`
	MemberMasterID10 *int `json:"member_master_id_10"`
	MemberMasterID11 *int `json:"member_master_id_11"`
	MemberMasterID12 *int `json:"member_master_id_12"`
	SuitMasterID1    *int `json:"suit_master_id_1"`
	SuitMasterID2    *int `json:"suit_master_id_2"`
	SuitMasterID3    *int `json:"suit_master_id_3"`
	SuitMasterID4    *int `json:"suit_master_id_4"`
	SuitMasterID5    *int `json:"suit_master_id_5"`
	SuitMasterID6    *int `json:"suit_master_id_6"`
	SuitMasterID7    *int `json:"suit_master_id_7"`
	SuitMasterID8    *int `json:"suit_master_id_8"`
	SuitMasterID9    *int `json:"suit_master_id_9"`
	SuitMasterID10   *int `json:"suit_master_id_10"`
	SuitMasterID11   *int `json:"suit_master_id_11"`
	SuitMasterID12   *int `json:"suit_master_id_12"`
}

func (ulmd *UserLiveMvDeck) ID() int64 {
	return int64(ulmd.LiveMasterID)
}

// LiveStageInfo ...
type LiveStageInfo struct {
	LiveDifficultyID int               `json:"live_difficulty_id"`
	LiveNotes        []LiveNote        `json:"live_notes"`
	LiveWaveSettings []LiveWaveSetting `json:"live_wave_settings"`
	NoteGimmicks     []NoteGimmick     `json:"note_gimmicks"`
	StageGimmickDict []any             `json:"stage_gimmick_dict"`
}

// LiveNotes ...
type LiveNote struct {
	ID                  int `json:"id"`
	CallTime            int `json:"call_time"`
	NoteType            int `json:"note_type"`
	NotePosition        int `json:"note_position"`
	GimmickID           int `json:"gimmick_id"`
	NoteAction          int `json:"note_action"`
	WaveID              int `json:"wave_id"`
	NoteRandomDropColor int `json:"note_random_drop_color"`
	AutoJudgeType       int `json:"auto_judge_type"`
}

// LiveWaveSetting
type LiveWaveSetting struct {
	ID            int `json:"id"`
	WaveDamage    int `json:"wave_damage"`
	MissionType   int `json:"mission_type"`
	Arg1          int `json:"arg_1"`
	Arg2          int `json:"arg_2"`
	RewardVoltage int `json:"reward_voltage"`
}

// NoteGimmick
type NoteGimmick struct {
	UniqID          int `json:"uniq_id"`
	ID              int `json:"id"`
	NoteGimmickType int `json:"note_gimmick_type"`
	Arg1            int `json:"arg_1"`
	Arg2            int `json:"arg_2"`
	EffectMID       int `json:"effect_m_id"`
	IconType        int `json:"icon_type"`
}

// MemberLovePanels ...
type MemberLovePanels struct {
	MemberID               int   `json:"member_id"`
	MemberLovePanelCellIds []int `json:"member_love_panel_cell_ids"`
}

type LiveUpdatePlayListReq struct {
	LiveMasterID int  `json:"live_master_id"`
	GroupNum     int  `json:"group_num"`
	IsSet        bool `json:"is_set"`
}

type UserPlayListItem struct {
	UserID         int  `xorm:"pk 'user_id'" json:"-"`
	UserPlayListID int  `xorm:"pk 'user_play_list_id'" json:"user_play_list_id"`
	GroupNum       int  `xorm:"'group_num'" json:"group_num"` // UserPlayListID % 10
	LiveID         int  `xorm:"'live_id'" json:"live_id"`     // UserPlayListID / 10
	IsNull         bool `xorm:"-" json:"-"`
}

func (item UserPlayListItem) ID() int64 {
	return int64(item.UserPlayListID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_live_deck"] = UserLiveDeck{}
	TableNameToInterface["u_live_party"] = UserLiveParty{}
	TableNameToInterface["u_live_state"] = LiveState{}
	TableNameToInterface["u_play_list"] = UserPlayListItem{}
}