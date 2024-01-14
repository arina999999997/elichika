package serverdata

import (
	"elichika/client"
	"elichika/config"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
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

func InitGacha(session *xorm.Session, args []string) {
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

func InsertGacha(session *xorm.Session, args []string) {
	if len(args) == 0 {
		fmt.Println("Invalid params:", args)
		return
	}
	// insert gacha from json format, with some exceptions.
	file := args[0]
	gachaJsons := utils.ReadAllText(file)
	gachas := []client.Gacha{}
	err := json.Unmarshal([]byte(gachaJsons), &gachas)
	utils.CheckErr(err)
	for pos, gacha := range gachas {
		serverGacha := ServerGacha{
			GachaMasterId:  gacha.GachaMasterId,
			GachaGroups:    []int32{},
			DrawGuarantees: [][]int32{},
		}
		bytes := []byte(gjson.Get(gachaJsons, fmt.Sprintf("%d.gacha_groups", pos)).String())
		for i := range gacha.GachaDraws.Slice {
			guarantee := []int32{}
			// TODO(gacha): Remove this magic id and use a link or something
			gacha.GachaDraws.Slice[i].GachaDrawMasterId = gacha.GachaMasterId*10 + int32(i)
			bytes := []byte(gjson.Get(gachaJsons, fmt.Sprintf("%d.gacha_draws.%d.guarantees", pos, i)).String())
			err := json.Unmarshal(bytes, &guarantee)
			utils.CheckErr(err)
			serverGacha.DrawGuarantees = append(serverGacha.DrawGuarantees, guarantee)
		}
		err := json.Unmarshal(bytes, &serverGacha.GachaGroups)
		utils.CheckErr(err)
		serverGacha.ClientGacha = gacha
		_, err = session.Table("s_gacha").Insert(serverGacha)
		utils.CheckErr(err)
	}
}

func GachaCli(session *xorm.Session, args []string) {
	if len(args) == 0 {
		fmt.Println("Invalid params:", args)
		return
	}
	switch args[0] {
	case "init":
		InitGacha(session, args[1:])
	case "insert":
		InsertGacha(session, args[1:])
	}
}
