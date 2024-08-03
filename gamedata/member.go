package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/generic"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Member struct {
	// from m_member
	Id          int32 `xorm:"'id' pk"`                           // member master id
	MemberGroup int32 `xorm:"'member_group'" enum:"MemberGroup"` // muse aqour niji
	SchoolGrade int32 `xorm:"school_grade"`

	// colors use rgba, 8 bits each
	// ThemeColor int `xorm:"'theme_color'"`
	// ThemeLightColor int `xorm:"'theme_light_color'"`
	// ThemeDarkColor int `xorm:"'theme_dark_color'"`
	// BackgroundUpperLeftColor int `xorm:"'background_upper_left_color'"`
	// BackgroundBottomRightColor int `xorm:"'background_bottom_right_color'"`

	// names and info are strings to dictionary_k
	Name string `xorm:"'name'"`
	// NameHiragana string `xorm:"'name_hiragana'"`
	// NameRomaji string `xorm:"'name_romaji'"`
	// Height string
	// Birthday string

	BirthMonth int32 `xorm:"birth_month"`
	BirthDay   int32 `xorm:"birth_day"`

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
	LoveLevelRewardIds []int32              `xorm:"-"` // 2 indexed for love level
	LoveLevelRewards   []([]client.Content) `xorm:"-"` // 2 indexed for love level

	// from m_member_unit_detail
	MemberUnit int32 `xorm:"-"` // subgroup

	// from m_member_init
	MemberInit MemberInit `xorm:"-"`

	// from m_member_login_bonus_birthday
	MemberLoginBonusBirthdays []MemberLoginBonusBirthday `xorm:"-"`

	// from s_event_member_name_asset
	MainNameTopAssetPath    client.TextureStruktur `xorm:"-"`
	MainNameBottomAssetPath client.TextureStruktur `xorm:"-"`
	SubNameTopAssetPath     client.TextureStruktur `xorm:"-"`
	SubNameBottomAssetPath  client.TextureStruktur `xorm:"-"`
}

type MemberInit struct {
	// from m_member_init
	// MemberMId int
	SuitMasterId        int32 `xorm:"'suit_m_id'"`
	CustomBackgroundMId int32 `xorm:"'custom_background_m_id'"`
	LoveLevel           int32 `xorm:"'love_level'"`
	LovePoint           int32 `xorm:"'love_point'"`
	LovePointLimit      int32 `xorm:"'love_point_limit'"`
}

func (member *Member) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {

	{
		type LoveLevelReward struct {
			Id        int32          `xorm:"pk 'id'"`
			LoveLevel int            `xorm:"'love_level'"`
			Content   client.Content `xorm:"extends"`
		}
		rewards := []LoveLevelReward{}
		err := masterdata_db.Table("m_member_love_level_reward").Where("member_m_id = ?", member.Id).OrderBy("love_level").Find(&rewards)
		utils.CheckErr(err)
		for i := int32(0); i <= gamedata.MemberLoveLevelCount; i++ {
			member.LoveLevelRewards = append(member.LoveLevelRewards, []client.Content{})
			member.LoveLevelRewardIds = append(member.LoveLevelRewardIds, 0)
		}
		for _, reward := range rewards {
			member.LoveLevelRewardIds[reward.LoveLevel] = reward.Id
			member.LoveLevelRewards[reward.LoveLevel] = append(member.LoveLevelRewards[reward.LoveLevel], reward.Content)
		}
	}

	{
		exist, err := masterdata_db.Table("m_member_init").Where("member_m_id = ?", member.Id).Get(&member.MemberInit)
		utils.CheckErrMustExist(err, exist)
	}

	{
		exist, err := masterdata_db.Table("m_member_unit_detail").Where("member_m_id = ?", member.Id).Cols("member_unit").Get(&member.MemberUnit)
		utils.CheckErrMustExist(err, exist)
	}

	member.Name = dictionary.Resolve(member.Name)
	// member.NameHiragana = dictionary.Resolve(member.NameHiragana)
	// member.NameRomaji = dictionary.Resolve(member.NameRomaji)
	// fmt.Println(member.Id, "\t", member.Name, "\t", member.NameHiragana, "\t", member.NameRomaji, "\t",
	// 	member.ThemeColor, "\t", member.ThemeLightColor, "\t", member.ThemeDarkColor, "\t",
	// 	member.BackgroundUpperLeftColor, "\t", member.BackgroundBottomRightColor)

	{
		asset := serverdata.EventMemberNameAsset{}
		exist, err := serverdata_db.Table("s_event_member_name_asset").Where("member_master_id = ?", member.Id).Get(&asset)
		utils.CheckErrMustExist(err, exist)

		member.MainNameTopAssetPath = client.TextureStruktur{
			V: generic.NewNullable(asset.MainNameTopAssetPath),
		}
		member.MainNameBottomAssetPath = client.TextureStruktur{
			V: generic.NewNullable(asset.MainNameBottomAssetPath),
		}
		member.SubNameTopAssetPath = client.TextureStruktur{
			V: generic.NewNullable(asset.SubNameTopAssetPath),
		}
		member.SubNameBottomAssetPath = client.TextureStruktur{
			V: generic.NewNullable(asset.SubNameBottomAssetPath),
		}
	}
}

func loadMember(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Member")
	gamedata.Member = make(map[int32]*Member)
	err := masterdata_db.Table("m_member").Find(&gamedata.Member)
	utils.CheckErr(err)
	gamedata.MemberByBirthday = make(map[int32]([]*Member))
	for _, member := range gamedata.Member {
		member.populate(gamedata, masterdata_db, serverdata_db, dictionary)
		gamedata.MemberByBirthday[member.BirthMonth*100+member.BirthDay] =
			append(gamedata.MemberByBirthday[member.BirthMonth*100+member.BirthDay], member)
	}
}

func init() {
	addLoadFunc(loadMember)
	addPrequisite(loadMember, loadMemberLoveLevel)
	addPrequisite(loadMember, loadMemberGroup)
}
