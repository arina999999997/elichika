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
	Id          int `xorm:"'id' pk"`        // member master id
	MemberGroup int `xorm:"'member_group'"` // muse aqour niji
	SchoolGrade int `xorm:"school_grade"`

	// colors use rgba, 8 bits each
	// ThemeColor int `xorm:"'theme_color'"`
	// ThemeLightColor int `xorm:"'theme_light_color'"`
	// ThemeDarkColor int `xorm:"'theme_dark_color'"`
	// BackgroundUpperLeftColor int `xorm:"'background_upper_left_color'"`
	// BackgroundBottomRightColor int `xorm:"'background_bottom_right_color'"`

	// names and info are strings to dictionary_k
	// Name string `xorm:"'name'"`
	// NameHiragana string `xorm:"'name_hiragana'"`
	// NameRomaji string `xorm:"'name_romaji'"`
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
	LoveLevelRewards []([]model.Content) `xorm:"-"` // 2 indexed for love level

	// from m_member_unit_detail
	// Unit int // subgroup

	// from m_member_init
	MemberInit MemberInit `xorm:"-"`

	// from m_member_login_bonus_birthday
	MemberLoginBonusBirthdays []MemberLoginBonusBirthday `xorm:"-"`
}

type MemberInit struct {
	// from m_member_init
	// MemberMId int
	SuitMasterId        int `xorm:"'suit_m_id'"`
	CustomBackgroundMId int `xorm:"'custom_background_m_id'"`
	LoveLevel           int `xorm:"'love_level'"`
	LovePoint           int `xorm:"'love_point'"`
	LovePointLimit      int `xorm:"'love_point_limit'"`
}

func (member *Member) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {

	{
		type LoveLevelReward struct {
			LoveLevel int           `xorm:"'love_level'"`
			Content   model.Content `xorm:"extends"`
		}
		rewards := []LoveLevelReward{}
		err := masterdata_db.Table("m_member_love_level_reward").Where("member_m_id = ?", member.Id).Find(&rewards)
		utils.CheckErr(err)
		for i := 0; i <= gamedata.MemberLoveLevelCount; i++ {
			member.LoveLevelRewards = append(member.LoveLevelRewards, []model.Content{})
		}
		for _, reward := range rewards {
			member.LoveLevelRewards[reward.LoveLevel] = append(member.LoveLevelRewards[reward.LoveLevel], reward.Content)
		}
	}

	{
		exist, err := masterdata_db.Table("m_member_init").Where("member_m_id = ?", member.Id).Get(&member.MemberInit)
		utils.CheckErrMustExist(err, exist)
	}

	// member.Name = dictionary.Resolve(member.Name)
	// member.NameHiragana = dictionary.Resolve(member.NameHiragana)
	// member.NameRomaji = dictionary.Resolve(member.NameRomaji)
	// fmt.Println(member.Id, "\t", member.Name, "\t", member.NameHiragana, "\t", member.NameRomaji, "\t",
	// 	member.ThemeColor, "\t", member.ThemeLightColor, "\t", member.ThemeDarkColor, "\t",
	// 	member.BackgroundUpperLeftColor, "\t", member.BackgroundBottomRightColor)
}

func loadMember(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Member")
	gamedata.Member = make(map[int]*Member)
	err := masterdata_db.Table("m_member").Find(&gamedata.Member)
	utils.CheckErr(err)
	gamedata.MemberByBirthday = make(map[int]([]*Member))
	for _, member := range gamedata.Member {
		member.populate(gamedata, masterdata_db, serverdata_db, dictionary)
		gamedata.MemberByBirthday[member.BirthMonth*100+member.BirthDay] =
			append(gamedata.MemberByBirthday[member.BirthMonth*100+member.BirthDay], member)
	}
}

func init() {
	addLoadFunc(loadMember)
	addPrequisite(loadMember, loadMemberLoveLevel)
}
