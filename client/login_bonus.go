package client

// the struct in client also has plural
type LoginBonusRewards struct {
	Day                int       `json:"day"`
	Status             int       `json:"status"`
	ContentGrade       int       `json:"content_grade"`
	LoginBonusContents []Content `json:"login_bonus_contents"`
}

// used for extra login bonus event that use an illustration?
type IllustLoginBonus struct {
	LoginBonusId      int                 `json:"login_bonus_id"`
	LoginBonusRewards []LoginBonusRewards `json:"login_bonus_rewards"`
	BackgroundId      int                 `json:"background_id"`
	StartAt           int64               `json:"start_at"`
	EndAt             int64               `json:"end_at"`
}

// this is the normal login bonus
type NaviLoginBonus struct {
	LoginBonusId           int                 `json:"login_bonus_id"`
	LoginBonusRewards      []LoginBonusRewards `json:"login_bonus_rewards"`
	BackgroundId           int                 `json:"background_id"`
	WhiteboardTextureAsset *TextureStruktur    `json:"whiteboard_texture_asset"`
	StartAt                int64               `json:"start_at"`
	EndAt                  int64               `json:"end_at"`
	// these doesn't seem to do anything but they have to be present
	MaxPage     int `json:"max_page"`
	CurrentPage int `json:"current_page"`
}

type LoginBonusBirthDayMember struct {
	MemberMasterId int `json:"member_master_id"`
	SuitMasterId   int `json:"suit_master_id"`
}

type BootstrapLoginBonus struct {
	Event2DLoginBonuses    []IllustLoginBonus         `json:"event_2d_login_bonuses"`
	LoginBonuses           []NaviLoginBonus           `json:"login_bonuses"`
	Event3DLoginBonus      []NaviLoginBonus           `json:"event_3d_login_bonuses"`
	BeginnerLoginBonuses   []NaviLoginBonus           `json:"beginner_login_bonuses"`
	ComebackLoginBonuses   []IllustLoginBonus         `json:"comeback_login_bonuses"`
	BirthdayLoginBonuses   []NaviLoginBonus           `json:"birthday_login_bonuses"`
	BirthdayMember         []LoginBonusBirthDayMember `json:"birth_day_member"`
	NextLoginBonsReceiveAt int64                      `json:"next_login_bons_receive_at"` // this is correct
}

type LoginBonus struct {
	LoginBonusId            int `xorm:"pk"`
	LoginBonusType          int `xorm:"pk"`
	StartAt                 int64
	EndAt                   int64
	BackgroundId            int
	WhiteboardTextureAsset  *TextureStruktur `xorm:"varchar(40)"`
	LoginBonusHandler       string
	LoginBonusHandlerConfig string
}

type LoginBonusRewardDay struct {
	LoginBonusId int `xorm:"pk"`
	Day          int `xorm:"pk"`
	ContentGrade int
}

type LoginBonusRewardContent struct {
	LoginBonusId int
	Day          int
	Content      Content `xorm:"extends"`
}
