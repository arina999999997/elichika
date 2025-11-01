package serverdata

import (
	"elichika/client"
	"elichika/config"
	// "elichika/enum"
	"elichika/log"
	"elichika/parser"
	"elichika/utils"

	// "encoding/json"
	// "fmt"
	"os"
	// "path/filepath"
	// "strings"

	"xorm.io/xorm"
)

type EventMining struct {
	EventId             int32  `xorm:"pk 'event_id'" json:"event_id"`
	TitleImagePath      string `xorm:"'title_image_path'" json:"title_image_path"`
	BackgroundImagePath string `xorm:"'background_image_path'" json:"background_image_path"`
	BgmAssetPath        string `xorm:"'bgm_asset_path'" json:"bgm_asset_path"`
	GachaMasterId       int32  `xorm:"'gacha_master_id'" json:"gacha_master_id"`
	// Equal EventId - 11000
	// EventPointMasterId            int32             `xorm:"'event_point_master_id'" json:"event_point_master_id"`
	EventCompetitionLiveIds []int32 `xorm:"'event_competition_live_ids'" json:"event_competition_live_ids"`
	// this is interpreted as the amount of score for each song
	EventCompetitionSelectionAmount int32 `xorm:"'event_competition_selection_amount'" json:"event_competition_selection_amount"`
	// trade master id is not stored here, it should fetch from m_event_mining_trade, but we just default to event_id instead
	// TradeMasterId int32 `xorm:"'trade_master_id'" json:"trade_master_id"` // usually is just event id
	RuleDescriptionPagesAssetPath []string `xorm:"-" json:"rule_description_pages_asset_path"`
	// the topic reward for voltage ranking just show an image
	VoltageRankingTopicRewardAssetPath string `xorm:"'voltage_ranking_topic_reward_asset_path'" json:"voltage_ranking_topic_reward_asset_path"`
	BannerImagePath                    string `xorm:"'banner_image_path'" json:"banner_image_path"`
}

type EventMiningRankingReward struct {
	EventId                int32  `xorm:"pk 'event_id'"`
	RankingRewardMasterId  int32  `xorm:"pk 'ranking_reward_master_id'"`
	UpperRank              int32  `xorm:"'upper_rank'"`
	LowerRank              *int32 `xorm:"'lower_rank'"`
	RewardGroupId          int32  `xorm:"'reward_group_id'"`
	RankingResultPrizeType int32  `xorm:"'ranking_result_prize_type'" enum:"EventRankingResultPrizeType"`
	// id convention: RewardGroupId = RankingRewardMasterId
	// RankingRewardMasterId is EventId * 10000 + id
	// where id is from 1000 to 1999 for point ranking
	// and id is from 2000 to 2999 for voltage ranking
}

type EventMiningReward struct {
	// event id is not necessary because this reference reward_group_id directly
	// but it's still present to load the data easier
	EventId       int32          `xorm:"pk 'event_id'"`
	RewardGroupId int32          `xorm:"pk 'reward_group_id'"`
	RewardContent client.Content `xorm:"extends"`
	DisplayOrder  int32          `xorm:"pk 'display_order'"`
}

type EventMiningRuleDescriptionPage struct {
	EventId int32 `xorm:"pk 'event_id'"`
	Page    int32 `xorm:"pk 'page'"`
	// Title          client.LocalizedText `xorm:"'title'"`
	ImageAssetPath string `xorm:"'image_asset_path'"`
}

// following are types that describe the image or animation that show up on event top
// the placing algorithm is as follow:
// - the screen space is divided into square grids of 5 columns, the amount of row is dynamic
// - each image can be small or big:
//   - small image take up 1 x 1 cells
//   - big image take up 2 x 2 cells
// - each image has priority, the image with higher priority go first
// - for each image, search for the first empty cell and put it there:
//   - a cell is empty if it is not used by any image, including the extra cells of big one
//   - this process can put a big cell on the 5th columns, in which case it go out of bound and cause the first cell of next row
// - repeat until done:
//   - not sure if there's a limit on the amount of row
//   - some row on the last cell can be unfilled
// each object also have a is_popup attribute:
// - setting this to true allow clicking on it to zoom in and like / unlike it
// - setting it to false disable that
// - there's also popup_comment that supposedly go with it, but it doesn't seems to have any effect
// reading an event story would unlock a new image, with the last clear unlocking 2 images
// more filter image might also be added, presumablly to make complete rows
// comparing JP network data and data from GL blog post / video:
// - they have the same cells and animation, but different subcell
// - but every players on the same server seems to have the same subcell
// - we then make the assumption that the subcell is chosen randomly (by the server machine), and is shared among the player:
//   - if they had been manually chosen manually (randomly or not), there'd be no reason to have to change it for localisation
// - another theory is that the value is decided by how many like each character had, which could explain the difference between servers
// we will just manually make a panel and use it

type EventMiningTopStillCell struct {
	EventId int32 `xorm:"pk 'event_id'"`
	// id convention: event_id * 100 + cell id
	// cell id is from 16 to 26:
	// - this is because 1 to 15 is used for subcell
	// - 16 and 17 is for the 2 always unlocked item
	// - 18 - 26 are for items unlocked after story:
	// - 7 smaller image and 2 bigger image
	// - they will form a 5 x 3 grid at the end
	ThumbnailCellId int32 `xorm:"pk 'thumbnail_cell_id'"`
	AddStoryNumber  int32 `xorm:"'add_story_number'"`
	// bigger priority will show up first
	Priority                     int32   `xorm:"'priority'"`
	ImageThumbnailAssetPath      string  `xorm:"'image_thumbnail_asset_path'"`
	IsPopup                      bool    `xorm:"'is_popup'"`
	PopupImageThumbnailAssetPath string  `xorm:"'popup_image_thumbnail_asset_path'"`
	PopupComment                 *string `xorm:"'popup_comment'"`
	IsBig                        bool    `xorm:"'is_big'"`
}

// similar to EventMiningTopStill but can't be big, and can't be unlocked after story
type EventMiningTopStillSubCell struct {
	EventId    int32 `xorm:"pk 'event_id'"`
	PanelSetId int32 `xorm:"'panel_set_id'"`
	// id convention: event_id * 10000 + 1000(panel_set_id - 1) + cell id
	// where cell id = 1 to amount of cell in the panel set
	ThumbnailCellId int32 `xorm:"pk 'thumbnail_cell_id'"`
	// priority and id are not correlated in official server
	Priority                     int32  `xorm:"'priority'"`
	ImageThumbnailAssetPath      string `xorm:"'image_thumbnail_asset_path'"`
	IsPopup                      bool   `xorm:"'is_popup'"`
	PopupImageThumbnailAssetPath string `xorm:"'popup_image_thumbnail_asset_path'"`
}

// animated cell in event mining, these were used initially  but by the end they gave up doing it
// it seems like the game no longer has the ability to actually do this:
// - this is tested with item inside m_movie and with some recovered early asset
// we will give up on this for now, and add it later if we can get animation to work
type EventMiningTopAnimationCell struct {
}

// reproduced typos
type EventMiningBonusPopupOrderCardMater struct {
	EventId      int32 `xorm:"'event_id'"`
	CardMatserId int32 `xorm:"'card_matser_id'"`
	DisplayLine  int32 `xorm:"'display_line'"`
	DisplayOrder int32 `xorm:"'display_order'"`
	IsGacha      bool  `xorm:"'is_gacha'"`
}

// trading aspect:
// - event trade is separated from normal trade
// - when fetching trade, event trade will be checked for and added in if necessary
// - unlike normal trade, event trade also have limit that go away when the event end and reset the next time it come around
// Note that all the info about trade are available inside m_trade and m_trade_product, we only need to store:
// - what the product are
// - how much they cost
// - what is the buying limit
// TODO(trade): Reimplement normal trade system to follow this convention
type EventMiningTradeProduct struct {
	// event is can be calculated from product id, but we do this for easy loading
	EventId      int32          `xorm:"'event_id'"`
	ProductId    int32          `xorm:"'product_id'"`
	SourceAmount int32          `xorm:"'source_amount'"`
	StockAmount  *int32         `xorm:"'stock_amount'"`
	Content      client.Content `xorm:"extends"`
}

func init() {
	addTable("s_event_mining", EventMining{}, initEventMining)
	addTable("s_event_mining_point_ranking_topic_reward", EventTopicReward{}, nil)
	addTable("s_event_mining_reward", EventMiningReward{}, nil)
	addTable("s_event_mining_point_ranking_reward", EventMiningRankingReward{}, nil)
	addTable("s_event_mining_voltage_ranking_reward", EventMiningRankingReward{}, nil)
	addTable("s_event_mining_trade_product", EventMiningTradeProduct{}, nil)
	addTable("s_event_mining_top_still_cell", EventMiningTopStillCell{}, nil)
	addTable("s_event_mining_top_still_sub_cell", EventMiningTopStillSubCell{}, nil)
	// addTable("s_event_mining_top_animation_cell", EventMiningTopStillAnimationCell{}, nil)
	addTable("s_event_mining_bonus_popup_order_card_mater", EventMiningBonusPopupOrderCardMater{}, nil)
	addTable("s_event_mining_rule_description_page", EventMiningRuleDescriptionPage{}, nil)

}

func initEventMining(session *xorm.Session) {
	// each event is one directory with the following files:
	// - main.json: the main struction of the event (EventMining)
	// - bonus_popup_order.csv: the order the bonus popup is shown
	// - ... (just figure it out)

	entries, err := os.ReadDir(config.AssetPath + "event/mining")
	utils.CheckErr(err)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		path := config.AssetPath + "event/mining/" + entry.Name() + "/"
		log.Println("Parsing event mining: ", path)
		eventMining := EventMining{}
		parser.ParseJson(path+"main.json", &eventMining)
		_, err = session.Table("s_event_mining").Insert(eventMining)
		utils.CheckErr(err)
		log.Println(eventMining)

		ruleDescriptionPages := []EventMiningRuleDescriptionPage{}
		for i, assetPath := range eventMining.RuleDescriptionPagesAssetPath {
			ruleDescriptionPages = append(ruleDescriptionPages, EventMiningRuleDescriptionPage{
				EventId:        eventMining.EventId,
				Page:           int32(i + 1),
				ImageAssetPath: assetPath,
			})
		}

		_, err = session.Table("s_event_mining_rule_description_page").Insert(ruleDescriptionPages)
		utils.CheckErr(err)

		pointRankingTopicRewards := []EventTopicReward{}
		parser.ParseCsv(path+"point_ranking_topic_reward.csv", &pointRankingTopicRewards, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range pointRankingTopicRewards {
			pointRankingTopicRewards[i].EventId = eventMining.EventId
		}
		_, err = session.Table("s_event_mining_point_ranking_topic_reward").Insert(pointRankingTopicRewards)
		utils.CheckErr(err)

		bonusPopupOrders := []EventMiningBonusPopupOrderCardMater{}
		parser.ParseCsv(path+"bonus_popup_order.csv", &bonusPopupOrders, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range bonusPopupOrders {
			bonusPopupOrders[i].EventId = eventMining.EventId
		}
		_, err = session.Table("s_event_mining_bonus_popup_order_card_mater").Insert(bonusPopupOrders)
		utils.CheckErr(err)

		// rewards
		eventMiningPointRankingRewards := []EventMiningRankingReward{}
		eventMiningVoltageRankingRewards := []EventMiningRankingReward{}
		eventMiningRewards := []EventMiningReward{}

		for loop := 1; loop <= 2; loop++ {
			rankingRewards := []struct {
				UpperRank              int32
				ContentType            int32
				ContentId              int32
				ContentAmount          int32
				RankingResultPrizeType int32
			}{}
			if loop == 1 {
				parser.ParseCsv(path+"point_ranking_reward.csv", &rankingRewards, &parser.CsvContext{
					HasHeader: true,
				})
			} else {
				parser.ParseCsv(path+"voltage_ranking_reward.csv", &rankingRewards, &parser.CsvContext{
					HasHeader: true,
				})
			}
			n := len(rankingRewards)
			if n > 999 {
				log.Panic("Can't have more than 999 rewards due to id convention")
			}
			i := 0
			rankingOrder := int32(0)
			for i < n {
				upperRank := rankingRewards[i].UpperRank
				j := i
				for ; (j+1 < n) && (rankingRewards[j+1].UpperRank == upperRank); j++ {
				}

				if (i > 0) && (rankingRewards[i].UpperRank <= rankingRewards[i-1].UpperRank) {
					log.Panic("point ranking reward need to be sorted")
				}

				rankingOrder++
				rankingRewardMasterId := eventMining.EventId*10000 + 1000*int32(loop) + rankingOrder
				eventMiningRankingReward := EventMiningRankingReward{
					EventId:                eventMining.EventId,
					RankingRewardMasterId:  rankingRewardMasterId,
					UpperRank:              upperRank,
					RewardGroupId:          rankingRewardMasterId,
					RankingResultPrizeType: rankingRewards[i].RankingResultPrizeType,
				}
				if j+1 < n {
					eventMiningRankingReward.LowerRank = new(int32)
					*eventMiningRankingReward.LowerRank = rankingRewards[j+1].UpperRank - 1
				}
				if loop == 1 {
					eventMiningPointRankingRewards = append(eventMiningPointRankingRewards, eventMiningRankingReward)
				} else {
					eventMiningVoltageRankingRewards = append(eventMiningVoltageRankingRewards, eventMiningRankingReward)
				}
				for k := i; k <= j; k++ {
					if rankingRewards[k].RankingResultPrizeType != rankingRewards[i].RankingResultPrizeType {
						log.Panic("RankingResultPrizeType changed")
					}
					eventMiningRewards = append(eventMiningRewards, EventMiningReward{
						EventId:       eventMining.EventId,
						RewardGroupId: rankingRewardMasterId,
						RewardContent: client.Content{
							ContentType:   rankingRewards[k].ContentType,
							ContentId:     rankingRewards[k].ContentId,
							ContentAmount: rankingRewards[k].ContentAmount,
						},
						// official server seems to do this using the reward type
						// the order should be the higher on the left and the lower on the right
						DisplayOrder: int32(j - k + 1),
					})
				}
				i = j + 1
			}
		}

		_, err = session.Table("s_event_mining_point_ranking_reward").Insert(eventMiningPointRankingRewards)
		utils.CheckErr(err)
		_, err = session.Table("s_event_mining_voltage_ranking_reward").Insert(eventMiningVoltageRankingRewards)
		utils.CheckErr(err)
		_, err = session.Table("s_event_mining_reward").Insert(eventMiningRewards)
		utils.CheckErr(err)

		// stills
		topStillCells := []EventMiningTopStillCell{}
		parser.ParseCsv(path+"top_still_cell.csv", &topStillCells, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range topStillCells {
			topStillCells[i].EventId = eventMining.EventId
		}
		_, err := session.Table("s_event_mining_top_still_cell").Insert(topStillCells)
		utils.CheckErr(err)
		topStillSubCells := []EventMiningTopStillSubCell{}
		parser.ParseCsv(path+"top_still_sub_cell.csv", &topStillSubCells, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range topStillSubCells {
			topStillSubCells[i].EventId = eventMining.EventId
		}
		_, err = session.Table("s_event_mining_top_still_sub_cell").Insert(topStillSubCells)
		utils.CheckErr(err)
		// topAnimationCells := []...{}
		//
		// trades
		tradeProducts := []struct {
			ProductId     int32
			SourceAmount  int32
			StockAmount   *int32
			ContentType   int32
			ContentId     int32
			ContentAmount int32
		}{}
		parser.ParseCsv(path+"trade_product.csv", &tradeProducts, &parser.CsvContext{
			HasHeader: true,
		})
		tradeProductInserts := []EventMiningTradeProduct{}
		n := len(tradeProducts)
		for i := 0; i < n; i++ {
			tradeProductInserts = append(tradeProductInserts, EventMiningTradeProduct{
				EventId:      eventMining.EventId,
				ProductId:    tradeProducts[i].ProductId,
				SourceAmount: tradeProducts[i].SourceAmount,
				StockAmount:  tradeProducts[i].StockAmount,
				Content: client.Content{
					ContentType:   tradeProducts[i].ContentType,
					ContentId:     tradeProducts[i].ContentId,
					ContentAmount: tradeProducts[i].ContentAmount,
				},
			})
		}
		_, err = session.Table("s_event_mining_trade_product").Insert(tradeProductInserts)
		utils.CheckErr(err)
	}
}
