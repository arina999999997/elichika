package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Member struct {
	// from m_member
	ID          int `xorm:"'id' pk"`        // member master id
	MemberGroup int `xorm:"'member_group'"` // muse aqour niji
	SchoolGrade int `xorm:"school_grade"`

	// colors use rgba, 8 bits each
	// ThemeColor int `xorm:"'theme_color'"`
	// ThemeLightColor int
	// ThemeDarkColor int
	// BackgroundUpperLeftColor int
	// BackgroundBottomRightColor int

	// names and info are strings to dictionary_k
	// Name string
	// NameHiragana string
	// NameRomanji string
	// Height string
	// Birthday string

	BirthMonth int `xorm:"birth_month"`
	BirthDay   int `xorm:"birth_day"`

	// BloodType string
	// ZodiacSign string
	// StandingImageAssetPath string
	// Description string
	// AutographImageAssetPath string
	// MemberIconImageAssetPath string
	// MemberIconSmallImageAssetPath string
	// CharacterVoiceActor string
	// ThumbnailImageAssetPath string
	// BgmPath string
	// DisplayOrder int
	// IsProfileDarkColor bool
	// SmallMemberStillImageAssetPath string
	// MemberIconImageText string
	// SubscriptionPassBaseAssetPath string
	// DailyTheaterInlineImageAssetPath string
	// StandingThumbnailImageAssetPath string
	// StandingThumbnailBackgroundUpperColor int
	// StandingThumbnailBackgroundBottomColor int

	// from m_member_love_level_reward
	LoveLevelRewards []model.Content `xorm:"-"` // 2 indexed for love level

	// from m_member_unit_detail
	// Unit int // subgroup

	// from m_member_init
	MemberInit MemberInit `xorm:"-"`
}

type MemberInit struct {
	// from m_member_init
	// MemberMID int
	SuitMasterID        int `xorm:"'suit_m_id'"`
	CustomBackgroundMID int `xorm:"'custom_background_m_id'"`
	LoveLevel           int `xorm:"'love_level'"`
	LovePoint           int `xorm:"'love_point'"`
	LovePointLimit      int `xorm:"'love_point_limit'"`
}

func (member *Member) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {

	{
		err := masterdata_db.Table("m_member_love_level_reward").Where("member_m_id = ?", member.ID).OrderBy("love_level").Find(&member.LoveLevelRewards)
		utils.CheckErr(err)
		member.LoveLevelRewards = append([]model.Content{model.Content{}, model.Content{}}, member.LoveLevelRewards...)
	}

	{
		exists, err := masterdata_db.Table("m_member_init").Where("member_m_id = ?", member.ID).Get(&member.MemberInit)
		utils.CheckErrMustExist(err, exists)
	}
}

func loadMember(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Member")
	gamedata.Member = make(map[int]*Member)
	err := masterdata_db.Table("m_member").Find(&gamedata.Member)
	utils.CheckErr(err)
	for _, member := range gamedata.Member {
		member.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}
