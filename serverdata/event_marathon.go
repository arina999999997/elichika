package serverdata

import (
	"elichika/client"
	"elichika/config"
	// "elichika/enum"
	"elichika/parser"
	"elichika/utils"

	// "encoding/json"
	"fmt"
	"os"
	// "path/filepath"
	// "strings"

	"xorm.io/xorm"
)

// this is the storage of event in the database, while gamedata.EventMarathon is the processing structure
// TODO(tech): note that due to xorm limitation, TextureStruktur and SoundStruktur can't be used right now.
type EventMarathon struct {
	EventId             int32   `xorm:"pk 'event_id'" json:"event_id"`
	BoosterItemId       int32   `xorm:"'booster_item_id'" json:"booster_item_id"`
	TitleImagePath      *string `xorm:"'title_image_path'" json:"title_image_path"`
	BackgroundImagePath *string `xorm:"'background_image_path'" json:"background_image_path"`
	BoardBaseImagePath  *string `xorm:"'board_base_image_path'" json:"board_base_image_path"`
	BoardDecoImagePath  *string `xorm:"'board_deco_image_path'" json:"board_deco_image_path"`
	BgmAssetPath        *string `xorm:"'bgm_asset_path'" json:"bgm_asset_path"`
	GachaMasterId       int32   `xorm:"'gacha_master_id'" json:"gacha_master_id"`
}

type EventMarathonBoardThing struct {
	EventId                        int32  `xorm:"'event_id'"`
	EventMarathonBoardPositionType int32  `xorm:"'event_marathon_board_position_type'"`
	Position                       int32  `xorm:"'position'"` // physical position on the board
	AddStoryNumber                 int32  `xorm:"'add_story_number'"`
	Priority                       int32  `xorm:"'priority'"` // the order of the item
	ImageThumbnailAssetPath        string `xorm:"'image_thumbnail_asset_path'"`
}

type EventMarathonPointReward struct {
	EventId       int32 `xorm:"pk 'event_id'"`
	RequiredPoint int32 `xorm:"pk 'required_point'"`
	RewardGroupId int32 `xorm:"pk 'reward_group_id'"`
	// id convention: RewardGroupId is EventId * 10000 + id
	// where id is from 1 to 999
}

type EventMarathonRankingReward struct {
	EventId                int32  `xorm:"pk 'event_id'"`
	RankingRewardMasterId  int32  `xorm:"pk 'ranking_reward_master_id'"`
	UpperRank              int32  `xorm:"'upper_rank'"`
	LowerRank              *int32 `xorm:"'lower_rank'"`
	RewardGroupId          int32  `xorm:"'reward_group_id'"`
	RankingResultPrizeType int32  `xorm:"'ranking_result_prize_type'" enum:"EventRankingResultPrizeType"`
	// id convention: RewardGroupId = RankingRewardMasterId
	// RankingRewardMasterId is EventId * 10000 + id
	// where id is from 1000 to 1999
}

type EventMarathonReward struct {
	// event id is not necessary because this reference reward_group_id directly
	// but it's still present to load the data easier
	EventId       int32          `xorm:"pk 'event_id'"`
	RewardGroupId int32          `xorm:"pk 'reward_group_id'"`
	RewardContent client.Content `xorm:"extends"`
	DisplayOrder  int32          `xorm:"pk 'display_order'"`
}

type EventMarathonRuleDescriptionPage struct {
	EventId        int32                `xorm:"pk 'event_id'"`
	Page           int32                `xorm:"pk 'page'"`
	Title          client.LocalizedText `xorm:"'title'"`
	ImageAssetPath string               `xorm:"'image_asset_path'"`
}

// reproduced typos
type EventMarathonBonusPopupOrderCardMater struct {
	EventId      int32 `xorm:"'event_id'"`
	CardMatserId int32 `xorm:"'card_matser_id'"`
	DisplayLine  int32 `xorm:"'display_line'"`
	DisplayOrder int32 `xorm:"'display_order'"`
	IsGacha      bool  `xorm:"'is_gacha'"`
}

func init() {
	addTable("s_event_marathon", EventMarathon{}, initEventMarathon)
	addTable("s_event_marathon_board_thing", EventMarathonBoardThing{}, nil)
	addTable("s_event_marathon_total_topic_reward", EventTopicReward{}, nil)
	addTable("s_event_marathon_ranking_topic_reward", EventTopicReward{}, nil)
	addTable("s_event_marathon_point_reward", EventMarathonPointReward{}, nil)
	addTable("s_event_marathon_ranking_reward", EventMarathonRankingReward{}, nil)
	addTable("s_event_marathon_reward", EventMarathonReward{}, nil)
	addTable("s_event_marathon_rule_description_page", EventMarathonRuleDescriptionPage{}, nil)
	addTable("s_event_marathon_bonus_popup_order_card_mater", EventMarathonBonusPopupOrderCardMater{}, nil)
}

func initEventMarathon(session *xorm.Session) {
	// each event is one directory with the following files:
	// - main.json: the main structures of the event
	// - board.csv: the items that can appear on the board
	// - bonus_popup_order.csv: the order the bonus popup is shown
	// - ... (just figure it out)

	entries, err := os.ReadDir(config.AssetPath + "event/marathon")
	utils.CheckErr(err)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		path := config.AssetPath + "event/marathon/" + entry.Name() + "/"
		fmt.Println("Parsing event marathon: ", path)
		eventMarathon := EventMarathon{}
		parser.ParseJson(path+"main.json", &eventMarathon)
		_, err = session.Table("s_event_marathon").Insert(eventMarathon)
		utils.CheckErr(err)
		fmt.Println(eventMarathon)

		boardThings := []EventMarathonBoardThing{}
		parser.ParseCsv(path+"board.csv", &boardThings, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range boardThings {
			boardThings[i].EventId = eventMarathon.EventId
		}
		_, err = session.Table("s_event_marathon_board_thing").Insert(boardThings)
		utils.CheckErr(err)

		ruleDescriptionPages := []EventMarathonRuleDescriptionPage{}
		parser.ParseCsv(path+"rule_description_page.csv", &ruleDescriptionPages, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range ruleDescriptionPages {
			ruleDescriptionPages[i].EventId = eventMarathon.EventId
		}
		_, err = session.Table("s_event_marathon_rule_description_page").Insert(ruleDescriptionPages)
		utils.CheckErr(err)

		totalTopicRewards := []EventTopicReward{}
		parser.ParseCsv(path+"total_topic_reward.csv", &totalTopicRewards, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range totalTopicRewards {
			totalTopicRewards[i].EventId = eventMarathon.EventId
		}
		_, err = session.Table("s_event_marathon_total_topic_reward").Insert(totalTopicRewards)
		utils.CheckErr(err)

		rankingTopicRewards := []EventTopicReward{}
		parser.ParseCsv(path+"ranking_topic_reward.csv", &rankingTopicRewards, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range rankingTopicRewards {
			rankingTopicRewards[i].EventId = eventMarathon.EventId
		}
		_, err = session.Table("s_event_marathon_ranking_topic_reward").Insert(rankingTopicRewards)
		utils.CheckErr(err)

		bonusPopupOrders := []EventMarathonBonusPopupOrderCardMater{}
		parser.ParseCsv(path+"bonus_popup_order.csv", &bonusPopupOrders, &parser.CsvContext{
			StartField: 1,
			HasHeader:  true,
		})
		for i := range bonusPopupOrders {
			bonusPopupOrders[i].EventId = eventMarathon.EventId
		}
		_, err = session.Table("s_event_marathon_bonus_popup_order_card_mater").Insert(bonusPopupOrders)
		utils.CheckErr(err)

		// rewards
		eventMarathonPointRewards := []EventMarathonPointReward{}
		eventMarathonRankingRewards := []EventMarathonRankingReward{}
		eventMarathonRewards := []EventMarathonReward{}

		pointRewards := []struct {
			RequiredPoint int32
			ContentType   int32
			ContentId     int32
			ContentAmount int32
		}{}
		parser.ParseCsv(path+"point_reward.csv", &pointRewards, &parser.CsvContext{
			HasHeader: true,
		})
		if len(pointRewards) > 999 {
			panic("Can't have more than 999 rewards due to id convention")
		}
		for i, pointReward := range pointRewards {
			rewardGroupId := eventMarathon.EventId*10000 + 1 + int32(i)
			eventMarathonPointRewards = append(eventMarathonPointRewards, EventMarathonPointReward{
				EventId:       eventMarathon.EventId,
				RequiredPoint: pointReward.RequiredPoint,
				RewardGroupId: rewardGroupId,
			})
			eventMarathonRewards = append(eventMarathonRewards, EventMarathonReward{
				EventId:       eventMarathon.EventId,
				RewardGroupId: rewardGroupId,
				RewardContent: client.Content{
					ContentType:   pointReward.ContentType,
					ContentId:     pointReward.ContentId,
					ContentAmount: pointReward.ContentAmount,
				},
				DisplayOrder: 0,
			})
		}

		rankingRewards := []struct {
			UpperRank              int32
			ContentType            int32
			ContentId              int32
			ContentAmount          int32
			RankingResultPrizeType int32
		}{}
		parser.ParseCsv(path+"ranking_reward.csv", &rankingRewards, &parser.CsvContext{
			HasHeader: true,
		})
		n := len(rankingRewards)
		if n > 999 {
			panic("Can't have more than 999 rewards due to id convention")
		}
		i := 0
		rankingOrder := int32(0)
		for i < n {
			upperRank := rankingRewards[i].UpperRank
			j := i
			for ; (j+1 < n) && (rankingRewards[j+1].UpperRank == upperRank); j++ {
			}

			if (i > 0) && (rankingRewards[i].UpperRank <= rankingRewards[i-1].UpperRank) {
				panic("ranking reward need to be sorted")
			}

			rankingOrder++
			rankingRewardMasterId := eventMarathon.EventId*10000 + 1000 + rankingOrder
			eventMarathonRankingReward := EventMarathonRankingReward{
				EventId:                eventMarathon.EventId,
				RankingRewardMasterId:  rankingRewardMasterId,
				UpperRank:              upperRank,
				RewardGroupId:          rankingRewardMasterId,
				RankingResultPrizeType: rankingRewards[i].RankingResultPrizeType,
			}
			if j+1 < n {
				eventMarathonRankingReward.LowerRank = new(int32)
				*eventMarathonRankingReward.LowerRank = rankingRewards[j+1].UpperRank - 1
			}
			eventMarathonRankingRewards = append(eventMarathonRankingRewards, eventMarathonRankingReward)
			for k := i; k <= j; k++ {
				if rankingRewards[k].RankingResultPrizeType != rankingRewards[i].RankingResultPrizeType {
					panic("RankingResultPrizeType changed")
				}
				eventMarathonRewards = append(eventMarathonRewards, EventMarathonReward{
					EventId:       eventMarathon.EventId,
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

		_, err = session.Table("s_event_marathon_point_reward").Insert(eventMarathonPointRewards)
		utils.CheckErr(err)
		_, err = session.Table("s_event_marathon_ranking_reward").Insert(eventMarathonRankingRewards)
		utils.CheckErr(err)
		_, err = session.Table("s_event_marathon_reward").Insert(eventMarathonRewards)
		utils.CheckErr(err)
	}
}
