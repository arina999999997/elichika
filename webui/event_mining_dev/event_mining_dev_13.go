//go:build dev

package event_mining_dev

import (
	"elichika/gamedata"
	// "elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/log"
	"elichika/router"
	"elichika/utils"
	// "elichika/generic"
	"elichika/webui/form/image_form"

	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

// TODO: early event has the gem star before the cards, so add a phase array for that whenever necessary
var finalTradePhaseArray = []string{"ur", "sr", "background", "star_gem", "memento_ticket", "school_idol_wish", "accessories", "macaron", "member_memorial", "manual_seed", "school_idol_badge", "radiance", "live_aid", "training_aid", "gold_exp"}

var tradePhaseArray = finalTradePhaseArray // switch based on event id
func EventMiningDev13GET(ctx *gin.Context) {
	// maybe setup a phase array
	phase := ctx.DefaultQuery("phase", tradePhaseArray[0])
	form := image_form.ContentForm{
		TradeToggleButton: true,
	}
	sourceContentId := TopStatus.EventId - 11000
	gamedata := gamedata.GamedataByLocale["en"]
	// each phase has different format that depends on the phase
	switch phase {
	case "ur":
		urCardMasterId := TopStatus.EventPointRankingTopicRewardInfo.Slice[0].RewardContent.ContentId
		form.FormId = "event_trade_form_ur"
		form.DataLabel = "Event Trade Data UR (notation: stock,price|stock,price|stock,price|...)"
		form.DataId = "event_trade_data_ur"
		option := "1,34500|1,400000|1,600000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, urCardMasterId, 0, 1, sourceContentId, 34500))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, urCardMasterId, 0, 1, sourceContentId, 400000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, urCardMasterId, 0, 1, sourceContentId, 600000))
		form.NewRow()
		option = "1,36000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, urCardMasterId, 0, 1, sourceContentId, 36000))
		form.NewRow()
	case "sr":
		srCardMasterId1 := TopStatus.EventPointRankingTopicRewardInfo.Slice[1].RewardContent.ContentId
		srCardMasterId2 := TopStatus.EventPointRankingTopicRewardInfo.Slice[2].RewardContent.ContentId
		form.FormId = "event_trade_form_sr"
		form.DataLabel = "Event Trade Data SR (notation: card_master_id,stock,price|card_master_id,stock,price|card_master_id,stock,price|...)"
		form.DataId = "event_trade_data_sr"
		option := fmt.Sprintf("%d,1,7200|%d,1,7200|%d,2,12000|%d,3,70000", srCardMasterId2, srCardMasterId1, srCardMasterId1, srCardMasterId1)
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId2, 0, 1, sourceContentId, 7200))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId1, 0, 1, sourceContentId, 7200))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId1, 0, 2, sourceContentId, 12000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId1, 0, 3, sourceContentId, 70000))
		form.NewRow()
		option = fmt.Sprintf("%d,1,7200|%d,1,7200|%d,2,12000|%d,3,70000", srCardMasterId1, srCardMasterId2, srCardMasterId2, srCardMasterId2)
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId1, 0, 1, sourceContentId, 7200))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId2, 0, 1, sourceContentId, 7200))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId2, 0, 2, sourceContentId, 12000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCard, srCardMasterId2, 0, 3, sourceContentId, 70000))
		form.NewRow()
	case "background":
		form.FormId = "event_trade_form_background"
		form.DataLabel = "Event Trade Data background (notation: background_master_id, empty to stop, -background_master_id to remove existing one)"
		form.DataId = "event_trade_data_background"
		used := map[int32]bool{}
		for _, product := range tradeProducts {
			if product.ContentType == enum.ContentTypeCustomBackground {
				used[product.ContentId] = true
				form.AddImage(fmt.Sprint(-product.ContentId), image_form.GetTradeItemHTML(enum.ContentTypeCustomBackground, product.ContentId, 0, 1, sourceContentId, 4500))
			}
		}
		form.NewRow()
		form.NewRow()
		ids := []int32{}
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err := session.Table("m_custom_background").Cols("id").OrderBy("display_order").Find(&ids)
			utils.CheckErr(err)
		})
		for _, id := range ids {
			if used[id] {
				continue
			}
			form.AddImage(fmt.Sprint(id), image_form.GetTradeItemHTML(enum.ContentTypeCustomBackground, id, 0, 1, sourceContentId, 4500))
		}
	case "star_gem":
		form.FormId = "event_trade_form_star_gem"
		form.DataLabel = "Event Trade Data Star Gem (notation: amount,stock,price|amount,stock,price)"
		form.DataId = "event_trade_data_star_gem"
		option := "100,1,10500"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeSnsCoin, 0, 100, 1, sourceContentId, 10500))
		form.NewRow()
		option = "10,6,1350|100,1,12150"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeSnsCoin, 0, 10, 1, sourceContentId, 1350))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeSnsCoin, 0, 100, 1, sourceContentId, 12150))
		form.NewRow()
	case "memento_ticket":
		form.FormId = "event_trade_form_memento_ticket"
		form.DataLabel = "Event Trade Data Mememto Ticket (notation: amount,stock,price)"
		form.DataId = "event_trade_data_memento_ticket"
		option := "1,5,2000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeExchangeEventPoint, 21010, 1, 5, sourceContentId, 2000))
		form.NewRow()
	case "school_idol_wish":
		form.FormId = "event_trade_form_school_idol_wish"
		form.DataLabel = "Event Trade Data School idol wish (notation: amount,stock,price)"
		form.DataId = "event_trade_data_school_idol_wish"
		option := "1,30,10000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGradeUpper, 1805, 1, 30, sourceContentId, 10000))
		form.NewRow()
	case "accessories":
		form.FormId = "event_trade_form_accessory"
		form.DataLabel = "Event Trade Data Accessory (notation: accessory_id)"
		form.DataId = "event_trade_data_accessory"
		accessoryIds := []int32{}
		for id, accessory := range gamedata.Accessory {
			if accessory.RarityType != 30 {
				continue
			}
			accessoryIds = append(accessoryIds, id)
		}
		sort.Slice(accessoryIds, func(i, j int) bool {
			A := gamedata.Accessory[accessoryIds[i]]
			B := gamedata.Accessory[accessoryIds[j]]
			return A.Id < B.Id
		})
		for i, id := range accessoryIds {
			form.AddImage(fmt.Sprint(id), image_form.GetTradeItemHTML(enum.ContentTypeAccessory, id, 0, 1, sourceContentId, 3000))
			form.AddImage(fmt.Sprint(id), image_form.GetTradeItemHTML(enum.ContentTypeAccessory, id, 0, 5, sourceContentId, 5000))
			if i%3 == 2 {
				form.NewRow()
			}
		}
	case "macaron":
		form.FormId = "event_trade_form_macaron"
		form.DataLabel = "Event Trade Data Maracon (notation: amount,stock_for_normal,stock_for_sr,stock_for_ur (price stays the same, accounting for cheaper resource for shorter events))"
		form.DataId = "event_trade_data_macaron"
		option := "25,500,500,500"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1900, 25, 500, sourceContentId, 150))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1901, 25, 500, sourceContentId, 300))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1902, 25, 500, sourceContentId, 450))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1960, 25, 500, sourceContentId, 200))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1961, 25, 500, sourceContentId, 350))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1962, 25, 500, sourceContentId, 500))
		form.NewRow()
		option = "10,100,80,80"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1900, 10, 100, sourceContentId, 150))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1901, 10, 80, sourceContentId, 300))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1902, 10, 80, sourceContentId, 450))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1960, 10, 100, sourceContentId, 200))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1961, 10, 80, sourceContentId, 350))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1962, 10, 80, sourceContentId, 500))
		form.NewRow()
	case "member_memorial":
		form.FormId = "event_trade_form_member_memorial"
		form.DataLabel = "Event Trade Data Member Memorial (notation: amount,stock,price)"
		form.DataId = "event_trade_data_member_memorial"
		option := "100,5,2000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 8107, 100, 5, sourceContentId, 2000))
		form.NewRow()
		option = "5,1,120"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 8107, 5, 1, sourceContentId, 120))
	case "manual_seed":
		form.FormId = "event_trade_form_manual_seed"
		form.DataLabel = "Event Trade Data Manual and Seed (notation: amount,stock,has_ur_version_or_not (price stay the same))"
		form.DataId = "event_trade_data_manual_seed"
		option := "5,350,1"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12101, 5, 350, sourceContentId, 150))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12102, 5, 350, sourceContentId, 250))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12103, 5, 350, sourceContentId, 350))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12104, 5, 350, sourceContentId, 450))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13101, 5, 350, sourceContentId, 150))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13102, 5, 350, sourceContentId, 250))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13103, 5, 350, sourceContentId, 350))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13104, 5, 350, sourceContentId, 450))
		form.NewRow()
		option = "1,90,0"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12101, 1, 90, sourceContentId, 150))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12102, 1, 90, sourceContentId, 250))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 12103, 1, 90, sourceContentId, 350))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13101, 1, 90, sourceContentId, 150))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13102, 1, 90, sourceContentId, 250))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 13103, 1, 90, sourceContentId, 350))
	case "school_idol_badge":
		form.FormId = "event_trade_form_school_badge"
		form.DataLabel = "Event Trade Data School Idol Badge (notation: amount,stock,price)"
		form.DataId = "event_trade_data_school_badge"
		option := "1,10,500"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1700, 1, 10, sourceContentId, 500))
		form.NewRow()
		option = "1,10,1200"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeTrainingMaterial, 1700, 1, 10, sourceContentId, 1200))
		form.NewRow()
	case "radiance":
		form.FormId = "event_trade_form_school_idol_radiance"
		form.DataLabel = "Event Trade Data School idol radiance (notation: amount,stock,price)"
		form.DataId = "event_trade_data_school_idol_radiance"
		option := "1,30,3000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGradeUpper, 1800, 1, 30, sourceContentId, 3000))
		form.NewRow()
		option = "1,3,7200"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGradeUpper, 1800, 1, 3, sourceContentId, 7200))
	case "live_aid": // skip ticket and candies
		form.FormId = "event_trade_form_live_aid"
		form.DataLabel = "Event Trade Data Skip ticket and LP candies (notation: amount,stock,price|amount,stock,price|amount,stock,price for skip ticket, 50lp candy, 100lp candy)"
		form.DataId = "event_trade_data_live_aid"
		option := "10,25,1000|1,5,700|1,3,1400"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeLiveSkipTicket, 16001, 10, 25, sourceContentId, 1000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeRecoveryLp, 1300, 1, 5, sourceContentId, 700))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeRecoveryLp, 1301, 1, 3, sourceContentId, 1400))
		form.NewRow()
		option = "5,6,700|1,2,700|1,2,1400"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeLiveSkipTicket, 16001, 5, 6, sourceContentId, 700))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeRecoveryLp, 1300, 1, 2, sourceContentId, 700))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeRecoveryLp, 1301, 1, 2, sourceContentId, 1400))
		form.NewRow()
	case "training_aid":
		form.FormId = "event_trade_form_training_aid"
		form.DataLabel = "Event Trade Data Insight Pin and Lucky Charm (notation: stock,price|stock,price (amount is always 1))"
		form.DataId = "event_trade_data_training_aid"
		option := "5,3000|5,3000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeLessonEnhancingItem, 1400, 1, 5, sourceContentId, 3000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeLessonEnhancingItem, 1500, 1, 5, sourceContentId, 3000))
		form.NewRow()
		option = "20,3600|20,3600"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeLessonEnhancingItem, 1400, 1, 20, sourceContentId, 3600))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeLessonEnhancingItem, 1500, 1, 20, sourceContentId, 3600))
		form.NewRow()
	case "gold_exp":
		form.FormId = "event_trade_form_gold_exp"
		form.DataLabel = "Event Trade Data Gold and Exp (notation: amount,stock,price for first gold = first exp, the other ones are fixed)"
		form.DataId = "event_trade_data_gold_exp"
		option := "100000,10,5000"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGameMoney, 0, 100000, 10, sourceContentId, 5000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCardExp, 0, 100000, 10, sourceContentId, 5000))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGameMoney, 0, 10, 1000, sourceContentId, 10))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCardExp, 0, 10, 1000, sourceContentId, 10))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGameMoney, 0, 1, 1000, sourceContentId, 1))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCardExp, 0, 1, 1000, sourceContentId, 1))
		form.NewRow()
		option = "5000,50,450"
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGameMoney, 0, 5000, 10, sourceContentId, 450))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCardExp, 0, 5000, 10, sourceContentId, 450))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGameMoney, 0, 10, 1000, sourceContentId, 10))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCardExp, 0, 10, 1000, sourceContentId, 10))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeGameMoney, 0, 1, 1000, sourceContentId, 1))
		form.AddImage(option, image_form.GetTradeItemHTML(enum.ContentTypeCardExp, 0, 1, 1000, sourceContentId, 1))
		form.NewRow()
	default:
		log.Panic("unknown phase: " + phase)
	}
	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev13POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	phase := ctx.DefaultQuery("phase", tradePhaseArray[0])
	if phase == tradePhaseArray[0] {
		// discard the existing list
		tradeProducts = nil
	}
	keepCurrentPhase := false // whether to move to the next phase, used for phase with multiple steps
	// parse the input and append it
	// the input have the general structure of <trade_item>|<trade_item>|<trade_item>|..., representing multiple trade item in order
	// each trade_item are comma separated values that will represent:
	// - content_type
	// - content_id
	// - content_amount
	// - stock (how many time can player buy)
	// - price (price per purchase)
	// the values must be written in order
	// finally, use can pass in prefilled item of any of the type:
	// - if it is < 0 then it should be parsed
	// - otherwise it's a prefilled value
	parseTradeProducts := func(contentType, contentId, contentAmount, stock, price int32, data string) {
		products := strings.Split(data, "|")
		for _, product := range products {
			fields := strings.Split(product, ",")
			tradeProduct := EventMiningTradeProduct{
				ContentType:   contentType,
				ContentId:     contentId,
				ContentAmount: contentAmount,
				Stock:         stock,
				Price:         price,
			}
			fieldId := 0
			if contentType < 0 {
				value, err := strconv.Atoi(fields[fieldId])
				utils.CheckErr(err)
				tradeProduct.ContentType = int32(value)
				fieldId++
			}
			if contentId < 0 {
				value, err := strconv.Atoi(fields[fieldId])
				utils.CheckErr(err)
				tradeProduct.ContentId = int32(value)
				fieldId++
			}
			if contentAmount < 0 {
				value, err := strconv.Atoi(fields[fieldId])
				utils.CheckErr(err)
				tradeProduct.ContentAmount = int32(value)
				fieldId++
			}
			if stock < 0 {
				value, err := strconv.Atoi(fields[fieldId])
				utils.CheckErr(err)
				tradeProduct.Stock = int32(value)
				fieldId++
			}
			if price < 0 {
				value, err := strconv.Atoi(fields[fieldId])
				utils.CheckErr(err)
				tradeProduct.Price = int32(value)
				fieldId++
			}
			if fieldId != len(fields) {
				log.Panic("extra field when parsing trade product: " + product)
			}
			tradeProducts = append(tradeProducts, tradeProduct)
		}
	}

	switch phase {
	case "ur":
		urCardMasterId := TopStatus.EventPointRankingTopicRewardInfo.Slice[0].RewardContent.ContentId
		data := form.Value["event_trade_data_ur"][0]
		parseTradeProducts(enum.ContentTypeCard, urCardMasterId, 1, -1, -1, data)
	case "sr":
		data := form.Value["event_trade_data_sr"][0]
		parseTradeProducts(enum.ContentTypeCard, -1, 1, -1, -1, data)
	case "background":
		data := form.Value["event_trade_data_background"][0]
		if data != "" {
			keepCurrentPhase = true
			if data[0] == '-' {
				id, err := strconv.Atoi(data)
				utils.CheckErr(err)
				id = -id
				for i := range tradeProducts {
					if tradeProducts[i].ContentType != enum.ContentTypeCustomBackground {
						continue
					}
					if tradeProducts[i].ContentId == int32(id) {
						tradeProducts = append(tradeProducts[:i], tradeProducts[i+1:]...)
						break
					}
				}
			} else {
				parseTradeProducts(enum.ContentTypeCustomBackground, -1, 1, 1, 4500, data)
			}
		}
	case "star_gem":
		data := form.Value["event_trade_data_star_gem"][0]
		parseTradeProducts(enum.ContentTypeSnsCoin, 0, -1, -1, -1, data)
	case "memento_ticket":
		data := form.Value["event_trade_data_memento_ticket"][0]
		parseTradeProducts(enum.ContentTypeExchangeEventPoint, 21010, -1, -1, -1, data)
	case "school_idol_wish":
		data := form.Value["event_trade_data_school_idol_wish"][0]
		parseTradeProducts(enum.ContentTypeGradeUpper, 1805, -1, -1, -1, data)
	case "accessories":
		data := form.Value["event_trade_data_accessory"][0]
		parseTradeProducts(enum.ContentTypeAccessory, -1, 1, 1, 3000, data)
		parseTradeProducts(enum.ContentTypeAccessory, -1, 1, 5, 5000, data)
	case "macaron":
		data := form.Value["event_trade_data_macaron"][0]
		fields := strings.Split(data, ",")
		amount, err := strconv.Atoi(fields[0])
		utils.CheckErr(err)
		normal_stock, err := strconv.Atoi(fields[1])
		utils.CheckErr(err)
		sr_stock, err := strconv.Atoi(fields[2])
		utils.CheckErr(err)
		ur_stock, err := strconv.Atoi(fields[3])
		utils.CheckErr(err)
		for colour := 0; colour <= 5; colour++ {
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     int32(1900 + colour*10 + 0),
				ContentAmount: int32(amount),
				Stock:         int32(normal_stock),
				Price:         150,
			})
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     int32(1900 + colour*10 + 1),
				ContentAmount: int32(amount),
				Stock:         int32(sr_stock),
				Price:         300,
			})
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     int32(1900 + colour*10 + 2),
				ContentAmount: int32(amount),
				Stock:         int32(ur_stock),
				Price:         450,
			})
		}
		for colour := 6; colour <= 7; colour++ {
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     int32(1900 + colour*10 + 0),
				ContentAmount: int32(amount),
				Stock:         int32(normal_stock),
				Price:         200,
			})
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     int32(1900 + colour*10 + 1),
				ContentAmount: int32(amount),
				Stock:         int32(sr_stock),
				Price:         350,
			})
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     int32(1900 + colour*10 + 2),
				ContentAmount: int32(amount),
				Stock:         int32(ur_stock),
				Price:         500,
			})
		}
	case "member_memorial":
		data := form.Value["event_trade_data_member_memorial"][0]
		fields := strings.Split(data, ",")
		amount, err := strconv.Atoi(fields[0])
		utils.CheckErr(err)
		stock, err := strconv.Atoi(fields[1])
		utils.CheckErr(err)
		price, err := strconv.Atoi(fields[2])
		utils.CheckErr(err)
		memberIds := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 101, 102, 103, 104, 105, 106, 107, 108, 109, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212}
		for _, memberId := range memberIds {
			tradeProducts = append(tradeProducts, EventMiningTradeProduct{
				ContentType:   enum.ContentTypeTrainingMaterial,
				ContentId:     8000 + memberId,
				ContentAmount: int32(amount),
				Stock:         int32(stock),
				Price:         int32(price),
			})
		}
	case "manual_seed":
		data := form.Value["event_trade_data_manual_seed"][0]
		fields := strings.Split(data, ",")
		amount, err := strconv.Atoi(fields[0])
		utils.CheckErr(err)
		stock, err := strconv.Atoi(fields[1])
		utils.CheckErr(err)
		hasUr, err := strconv.Atoi(fields[2])
		utils.CheckErr(err)
		for itemType := int32(12); itemType <= 13; itemType++ {
			for colour := int32(1); colour <= 4; colour++ {
				for tier := int32(1); tier <= 3+int32(hasUr); tier++ {
					contentId := itemType*1000 + colour*100 + tier
					price := 50 + tier*100
					tradeProducts = append(tradeProducts, EventMiningTradeProduct{
						ContentType:   enum.ContentTypeTrainingMaterial,
						ContentId:     contentId,
						ContentAmount: int32(amount),
						Stock:         int32(stock),
						Price:         price,
					})
				}
			}
		}
	case "school_idol_badge":
		data := form.Value["event_trade_data_school_badge"][0]
		parseTradeProducts(enum.ContentTypeTrainingMaterial, 1700, -1, -1, -1, data)
	case "radiance":
		data := form.Value["event_trade_data_school_idol_radiance"][0]
		parseTradeProducts(enum.ContentTypeGradeUpper, 1800, -1, -1, -1, data)
	case "live_aid": // skip ticket and candies
		data := form.Value["event_trade_data_live_aid"][0]
		fields := strings.Split(data, "|")
		parseTradeProducts(enum.ContentTypeLiveSkipTicket, 16001, -1, -1, -1, fields[0])
		parseTradeProducts(enum.ContentTypeRecoveryLp, 1300, -1, -1, -1, fields[1])
		parseTradeProducts(enum.ContentTypeRecoveryLp, 1301, -1, -1, -1, fields[2])
	case "training_aid":
		data := form.Value["event_trade_data_training_aid"][0]
		fields := strings.Split(data, "|")
		parseTradeProducts(enum.ContentTypeLessonEnhancingItem, 1400, 1, -1, -1, fields[0])
		parseTradeProducts(enum.ContentTypeLessonEnhancingItem, 1500, 1, -1, -1, fields[1])
	case "gold_exp":
		data := form.Value["event_trade_data_gold_exp"][0]
		parseTradeProducts(enum.ContentTypeGameMoney, 1200, -1, -1, -1, data)
		parseTradeProducts(enum.ContentTypeCardExp, 1100, -1, -1, -1, data)

		tradeProducts = append(tradeProducts, EventMiningTradeProduct{
			ContentType:   enum.ContentTypeGameMoney,
			ContentId:     1200,
			ContentAmount: 10,
			Stock:         1000,
			Price:         10,
		})
		tradeProducts = append(tradeProducts, EventMiningTradeProduct{
			ContentType:   enum.ContentTypeCardExp,
			ContentId:     1100,
			ContentAmount: 10,
			Stock:         1000,
			Price:         10,
		})
		tradeProducts = append(tradeProducts, EventMiningTradeProduct{
			ContentType:   enum.ContentTypeGameMoney,
			ContentId:     1200,
			ContentAmount: 1,
			Stock:         1000,
			Price:         1,
		})
		tradeProducts = append(tradeProducts, EventMiningTradeProduct{
			ContentType:   enum.ContentTypeCardExp,
			ContentId:     1100,
			ContentAmount: 1,
			Stock:         1000,
			Price:         1,
		})
	default:
		log.Panic("unknown phase: " + phase)
	}
	phaseIndex := 0
	for ; tradePhaseArray[phaseIndex] != phase; phaseIndex++ {
	}
	nextPhaseIndex := phaseIndex + 1
	if keepCurrentPhase {
		nextPhaseIndex = phaseIndex
	}
	if nextPhaseIndex == len(tradePhaseArray) { // finish building trade, finalise
		db := gamedata.Instance.MasterdataDb
		db.Do(func(session *xorm.Session) {
			count, err := session.Table("m_trade_product").Where("trade_master_id = ?", TopStatus.EventId).Count()
			utils.CheckErr(err)
			if int(count) != len(tradeProducts) {
				log.Printf("Warning: amount of actual trade product is different from database (actual: %s, database: %s)\n", len(tradeProducts), count)
			}
		})
		ctx.Header("Location", "/webui/event_mining_dev/final")
	} else {
		ctx.Header("Location", fmt.Sprintf("/webui/event_mining_dev/13?phase=%s", tradePhaseArray[nextPhaseIndex]))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/13", EventMiningDev13GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/13", EventMiningDev13POST)
	}
}
