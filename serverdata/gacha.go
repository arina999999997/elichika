package serverdata

import (
	"elichika/client"
	"elichika/config"
	"elichika/utils"

	"encoding/json"

	"xorm.io/xorm"
)

// different gacha can share groups
type GachaGroup struct { // s_gacha_group
	GroupMasterId int32 `xorm:"pk 'group_master_id'"`
	GroupWeight   int64 `xorm:"'group_weight'"`
}

type GachaCard struct { // s_gacha_card
	GroupMasterId int32 `xorm:"pk 'group_master_id'"`
	CardMasterId  int32 `xorm:"pk 'card_master_id'"`
}

// can be shared depending on impl
// GuaranteedCardSet is not stored, built by gamedata when loaded, if applicable
// - (static) CardSet can be used to specify almost if not all the guaranteed form official version had
// - More exotic form of guarantee should be built into the handler itself
// - If CardSetSQL is not empty, GuaranteedCardSet would contain the relevant cards Id
type GachaGuarantee struct { // s_gacha_guarantee
	GachaGuaranteeMasterId int32          `xorm:"pk 'gacha_guarantee_master_id'"` // unique id
	GuaranteeHandler       string         `xorm:"'handler'"`
	CardSetSQL             string         `xorm:"card_set_sql"`
	GuaranteedCardSet      map[int32]bool `xorm:"-"`
}

// TODO(gacha): Do this properly
type ServerGacha struct {
	GachaMasterId int32        `xorm:"pk"`
	ClientGacha   client.Gacha `xorm:"json"`
	// the groups (predefined) that the gacha can result in
	GachaGroups    []int32   `xorm:"json"`
	DrawGuarantees [][]int32 `xorm:"json"`
}

type GachaDrawSetupInfo struct {
	Guarantees []int32 `json:"guarantees"`
}
type GachaSetupInfo struct {
	GachaGroups []int32              `json:"gacha_groups"`
	GachaDraws  []GachaDrawSetupInfo `json:"gacha_draws"`
}

func InitGacha(session *xorm.Session) {
	// insert some relevant gacha group, gacha card, and gacha guarantee

	// this is the same for everything
	masterdata, err := xorm.NewEngine("sqlite", config.GlMasterdataPath+"masterdata.db")
	utils.CheckErr(err)
	// 9 groups for now:
	// (R, SR, UR) * (muse, aqours, niji)
	weight := make(map[int]int64)
	weight[10] = 85
	weight[20] = 10
	weight[30] = 5
	for rarity := 10; rarity <= 30; rarity += 10 {
		for school := 0; school < 3; school++ {
			groupMasterId := int32(rarity*10 + school)
			cardMasterIds := []int32{}
			err := masterdata.Table("m_card").Where("card_rarity_type = ? AND member_m_id / 100 == ?", rarity, school).
				Cols("id").Find(&cardMasterIds)
			utils.CheckErr(err)
			for _, cardMasterId := range cardMasterIds {
				_, err := session.Table("s_gacha_card").Insert(GachaCard{
					GroupMasterId: groupMasterId,
					CardMasterId:  cardMasterId,
				})
				utils.CheckErr(err)
			}
			session.Table("s_gacha_group").Insert(GachaGroup{
				GroupMasterId: groupMasterId,
				GroupWeight:   weight[rarity],
			})
		}
	}

	// gacha guarantee: new card
	session.Table("s_gacha_guarantee").Insert(GachaGuarantee{
		GachaGuaranteeMasterId: 0,
		GuaranteeHandler:       "guaranteed_new_card",
	})
	// gacha guarantee: UR card
	session.Table("s_gacha_guarantee").Insert(GachaGuarantee{
		GachaGuaranteeMasterId: 1,
		GuaranteeHandler:       "guaranteed_card_set",
		CardSetSQL:             "card_rarity_type = 30",
	})
	// gacha guarantee: SR+ card
	session.Table("s_gacha_guarantee").Insert(GachaGuarantee{
		GachaGuaranteeMasterId: 2,
		GuaranteeHandler:       "guaranteed_card_set",
		CardSetSQL:             "card_rarity_type >= 20",
	})
	// gacha guarantee: festival / party card
	session.Table("s_gacha_guarantee").Insert(GachaGuarantee{
		GachaGuaranteeMasterId: 3,
		GuaranteeHandler:       "guaranteed_card_set",
		CardSetSQL:             "passive_skill_slot == 2",
	})
}

func InsertGacha(session *xorm.Session, file string) {
	// insert gacha from json format, with some exceptions.
	gachaJsons := utils.ReadAllText(file)

	gachas := []client.Gacha{}
	err := json.Unmarshal([]byte(gachaJsons), &gachas)
	utils.CheckErr(err)

	gachaSetups := []GachaSetupInfo{}
	err = json.Unmarshal([]byte(gachaJsons), &gachaSetups)
	utils.CheckErr(err)

	for pos, gacha := range gachas {
		serverGacha := ServerGacha{
			GachaMasterId:  gacha.GachaMasterId,
			GachaGroups:    gachaSetups[pos].GachaGroups,
			DrawGuarantees: [][]int32{},
		}

		for i := range gacha.GachaDraws.Slice {
			// TODO(gacha): Remove this magic id and use a link or something
			gacha.GachaDraws.Slice[i].GachaDrawMasterId = gacha.GachaMasterId*10 + int32(i)
			serverGacha.DrawGuarantees = append(serverGacha.DrawGuarantees, gachaSetups[pos].GachaDraws[i].Guarantees)
		}
		serverGacha.ClientGacha = gacha
		_, err = session.Table("s_gacha").Insert(serverGacha)
		utils.CheckErr(err)
	}
}

func gachaInitializer(session *xorm.Session) {
	InitGacha(session)
	InsertGacha(session, config.ServerInitJsons+"gacha.json")
}

func init() {
	addTable("s_gacha_guarantee", GachaGuarantee{}, nil)
	addTable("s_gacha_group", GachaGroup{}, nil)
	addTable("s_gacha_card", GachaCard{}, nil)
	addTable("s_gacha", ServerGacha{}, gachaInitializer)
	// InitTable("s_gacha_guarantee", GachaGuarantee{}, overwrite)
	// InitTable("s_gacha", ServerGacha{}, overwrite)
	// InitTable("s_gacha_group", GachaGroup{}, overwrite)
	// InitTable("s_gacha_card", GachaCard{}, overwrite)
}
