package serverdata

import (
	"elichika/config"
	"elichika/model"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"xorm.io/xorm"
)

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
			groupMasterId := rarity*10 + school
			cardMasterIds := []int{}
			err := masterdata.Table("m_card").Where("card_rarity_type = ? AND member_m_id / 100 == ?", rarity, school).
				Cols("id").Find(&cardMasterIds)
			utils.CheckErr(err)
			for _, cardMasterId := range cardMasterIds {
				_, err := session.Table("s_gacha_card").Insert(model.GachaCard{
					GroupMasterId: groupMasterId,
					CardMasterId:  cardMasterId,
				})
				utils.CheckErr(err)
			}
			session.Table("s_gacha_group").Insert(model.GachaGroup{
				GroupMasterId: groupMasterId,
				GroupWeight:   weight[rarity],
			})
		}
	}

	// gacha guarantee: new card
	session.Table("s_gacha_guarantee").Insert(model.GachaGuarantee{
		GachaGuaranteeMasterId: 0,
		GuaranteeHandler:       "guaranteed_new_card",
	})
	// gacha guarantee: UR card
	session.Table("s_gacha_guarantee").Insert(model.GachaGuarantee{
		GachaGuaranteeMasterId: 1,
		GuaranteeHandler:       "guaranteed_card_set",
		CardSetSQL:             "card_rarity_type = 30",
	})
	// gacha guarantee: SR+ card
	session.Table("s_gacha_guarantee").Insert(model.GachaGuarantee{
		GachaGuaranteeMasterId: 2,
		GuaranteeHandler:       "guaranteed_card_set",
		CardSetSQL:             "card_rarity_type >= 20",
	})
	// gacha guarantee: festival / party card
	session.Table("s_gacha_guarantee").Insert(model.GachaGuarantee{
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
	gachas := []model.Gacha{}
	gachaJsons := utils.ReadAllText(file)

	err := json.Unmarshal([]byte(gachaJsons), &gachas)
	utils.CheckErr(err)
	for pos, gacha := range gachas {
		for i, appeal := range gacha.GachaAppeals {
			appeal.GachaAppealMasterId = gacha.GachaMasterId*10 + i
			gacha.DbGachaAppeals = append(gacha.DbGachaAppeals, appeal.GachaAppealMasterId)
			_, err := session.Table("s_gacha_appeal").Insert(appeal)
			utils.CheckErr(err)
		}
		for i, draw := range gacha.GachaDraws {
			draw.GachaMasterId = gacha.GachaMasterId
			gacha.DbGachaDraws = append(gacha.DbGachaDraws, draw.GachaDrawMasterId)
			gjson.Get(gachaJsons, fmt.Sprintf("%d.gacha_draws.%d.guarantees", pos, i)).ForEach(
				func(_, value gjson.Result) bool {
					draw.Guarantees = append(draw.Guarantees, int(value.Int()))
					return true
				})
			_, err := session.Table("s_gacha_draw").Insert(draw)
			utils.CheckErr(err)
		}
		gjson.Get(gachaJsons, fmt.Sprintf("%d.gacha_groups", pos)).ForEach(
			func(_, value gjson.Result) bool {
				gacha.DbGachaGroups = append(gacha.DbGachaGroups, int(value.Int()))
				return true
			})

		_, err := session.Table("s_gacha").Insert(gacha)
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
