package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/enum"
	"elichika/generic"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"
	"sort"

	"xorm.io/xorm"
)

// marathon event loading process:
// - load the main structure from s_event_marathon
// - load the board pictures and memos from s_event_marathon_board_thing
// - load the event stories from gamedata.StoryEventHistory
// - load the rewards from s_event_marathon_reward:
//   - this use gamedata.EventMarathonRewardGroups
//
// - load the event_ranking_topic_reward_info and event_total_topic_reward_info from the rewards:
//   - the event UR go first
//   - then the event SR that is not in final ranking
//   - then the other event SR
//
// - event_marathon_bonus_popup_order_card_mater_rows:
//   - the order seems to be gacha UR
//   - event UR
//   - gacha UR
//   - gacha UR
//   - event SR
//   - event SR
//   - gacha SR
//
// - event_marathon_bonus_popup_order_member_mater_rows is sorted by member id
type EventMarathon struct {
	Gamedata *Gamedata

	EventId       int32
	BoosterItemId int32

	// this is the top status template, COPY before use
	TopStatus client.EventMarathonTopStatus

	// relevant data that is can changed based one what user have done
	BoardMemos    []client.EventMarathonBoardMemorialThingsMasterRow
	BoardPictures []client.EventMarathonBoardMemorialThingsMasterRow

	// bonus mapping
	CardBonus   map[int32][]int32
	MemberBonus map[int32]int32

	// We don't need this right because we will only be using old event stories.
	// BoardStory []EventMarathonBoardStory

	// TODO(extra): Check if this data is available when the event start or only later on.
	// TODO(extra): check and implement loop rewards
}

func (em *EventMarathon) GetNextReward(eventPoint int32) (generic.Nullable[int32], generic.Nullable[client.Content]) {
	// TODO(optimization): this can be a binary search
	for _, reward := range em.TopStatus.EventMarathonPointRewardMasterRows.Slice {
		if reward.RequiredPoint > eventPoint {
			content := em.Gamedata.EventMarathonReward[reward.RewardGroupId][0]
			return generic.NewNullable(reward.RequiredPoint), generic.NewNullableFromPointer(content)
		}
	}
	return generic.Nullable[int32]{}, generic.Nullable[client.Content]{}
}

func (em *EventMarathon) GetRankingReward(rank int32) int32 {
	// TODO(optimization): this can be a binary search
	for _, reward := range em.TopStatus.EventMarathonRankingRewardMasterRows.Slice {
		if (!reward.LowerRank.HasValue) || (reward.LowerRank.Value >= rank) {
			return reward.RewardGroupId
		}
	}
	panic("wrong ranking reward")
	return 0
}

func loadEventMarathon(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading EventMarathon")
	gamedata.EventMarathon = make(map[int32]*EventMarathon)
	events := []serverdata.EventMarathon{}
	err := serverdata_db.Table("s_event_marathon").Find(&events)
	utils.CheckErr(err)
	for _, event := range events {
		eventMarathon := EventMarathon{
			Gamedata:      gamedata,
			EventId:       event.EventId,
			BoosterItemId: event.BoosterItemId,
			TopStatus: client.EventMarathonTopStatus{
				EventId: event.EventId,
				TitleImagePath: client.TextureStruktur{
					V: generic.NewNullableFromPointer(event.TitleImagePath),
				},
				BackgroundImagePath: client.TextureStruktur{
					V: generic.NewNullableFromPointer(event.BackgroundImagePath),
				},
				BgmAssetPath: client.SoundStruktur{
					V: generic.NewNullableFromPointer(event.BgmAssetPath),
				},
				GachaMasterId: event.GachaMasterId,
			},
		}
		eventMarathon.TopStatus.BoardStatus.BoardBaseImagePath.V = generic.NewNullableFromPointer(event.BoardBaseImagePath)
		eventMarathon.TopStatus.BoardStatus.BoardDecoImagePath.V = generic.NewNullableFromPointer(event.BoardDecoImagePath)

		{
			err = serverdata_db.Table("s_event_marathon_board_thing").
				Where("event_id = ? AND event_marathon_board_position_type = ?", event.EventId, enum.EventMarathonBoardPositionTypeMemo).
				OrderBy("priority").Find(&eventMarathon.BoardMemos)
			utils.CheckErr(err)
			err = serverdata_db.Table("s_event_marathon_board_thing").
				Where("event_id = ? AND event_marathon_board_position_type = ?", event.EventId, enum.EventMarathonBoardPositionTypePicture).
				OrderBy("priority").Find(&eventMarathon.BoardPictures)
			utils.CheckErr(err)
		}

		{
			eventStoryIds := []int32{}
			err = masterdata_db.Table("m_story_event_history_detail").Where("event_master_id = ?", event.EventId).
				OrderBy("story_number DESC").Cols("story_event_id").Find(&eventStoryIds)
			utils.CheckErr(err)
			for _, storyId := range eventStoryIds {
				eventMarathon.TopStatus.StoryStatus.Stories.Append(gamedata.EventStory[storyId].GetEventMarathonStory())
			}
		}

		{
			topicRewards := []serverdata.EventTopicReward{}
			err = serverdata_db.Table("s_event_marathon_total_topic_reward").Where("event_id = ?", event.EventId).
				OrderBy("display_order").Find(&topicRewards)
			utils.CheckErr(err)
			for _, topicReward := range topicRewards {
				member := gamedata.Card[topicReward.RewardCardId].Member
				eventMarathon.TopStatus.EventTotalTopicRewardInfo.Append(client.EventTopicReward{
					DisplayOrder: topicReward.DisplayOrder,
					RewardContent: client.Content{
						ContentType:   enum.ContentTypeCard,
						ContentId:     topicReward.RewardCardId,
						ContentAmount: topicReward.RewardCardAmount,
					},
					MainNameTopAssetPath:    member.MainNameTopAssetPath,
					MainNameBottomAssetPath: member.MainNameBottomAssetPath,
					SubNameTopAssetPath:     member.SubNameTopAssetPath,
					SubNameBottomAssetPath:  member.SubNameBottomAssetPath,
				})
			}
		}

		{
			topicRewards := []serverdata.EventTopicReward{}
			err = serverdata_db.Table("s_event_marathon_ranking_topic_reward").Where("event_id = ?", event.EventId).
				OrderBy("display_order").Find(&topicRewards)
			utils.CheckErr(err)
			for _, topicReward := range topicRewards {
				member := gamedata.Card[topicReward.RewardCardId].Member
				eventMarathon.TopStatus.EventRankingTopicRewardInfo.Append(client.EventTopicReward{
					DisplayOrder: topicReward.DisplayOrder,
					RewardContent: client.Content{
						ContentType:   enum.ContentTypeCard,
						ContentId:     topicReward.RewardCardId,
						ContentAmount: topicReward.RewardCardAmount,
					},
					MainNameTopAssetPath:    member.MainNameTopAssetPath,
					MainNameBottomAssetPath: member.MainNameBottomAssetPath,
					SubNameTopAssetPath:     member.SubNameTopAssetPath,
					SubNameBottomAssetPath:  member.SubNameBottomAssetPath,
				})
			}
		}

		err = serverdata_db.Table("s_event_marathon_point_reward").Where("event_id = ?", event.EventId).
			OrderBy("required_point").Find(&eventMarathon.TopStatus.EventMarathonPointRewardMasterRows.Slice)
		utils.CheckErr(err)

		err = serverdata_db.Table("s_event_marathon_ranking_reward").Where("event_id = ?", event.EventId).
			OrderBy("ranking_reward_master_id").Find(&eventMarathon.TopStatus.EventMarathonRankingRewardMasterRows.Slice)
		utils.CheckErr(err)

		err = serverdata_db.Table("s_event_marathon_reward").Where("event_id = ?", event.EventId).
			OrderBy("reward_group_id").OrderBy("display_order").Find(&eventMarathon.TopStatus.EventMarathonRewardMasterRows.Slice)
		utils.CheckErr(err)

		err = serverdata_db.Table("s_event_marathon_rule_description_page").Where("event_id = ?", event.EventId).
			OrderBy("page").Find(&eventMarathon.TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice)
		utils.CheckErr(err)

		err = serverdata_db.Table("s_event_marathon_bonus_popup_order_card_mater").Where("event_id = ?", event.EventId).
			Find(&eventMarathon.TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Slice)
		utils.CheckErr(err)

		{
			eventMarathon.CardBonus = map[int32][]int32{}
			type cardBonusValue struct {
				CardMasterId int32 `xorm:"'card_master_id'"`
				Grade        int32 `xorm:"'grade'"`
				Value        int32 `xorm:"'value'"`
			}
			bonuses := []cardBonusValue{}
			err = masterdata_db.Table("m_event_marathon_bonus_card").Where("event_marathon_master_id = ?", event.EventId).
				Find(&bonuses)
			utils.CheckErr(err)
			for _, bonus := range bonuses {
				_, exist := eventMarathon.CardBonus[bonus.CardMasterId]
				if !exist {
					eventMarathon.CardBonus[bonus.CardMasterId] = make([]int32, 6)
				}
				eventMarathon.CardBonus[bonus.CardMasterId][bonus.Grade] = bonus.Value
			}
		}

		{
			eventMarathon.MemberBonus = map[int32]int32{}
			type memberBonusValue struct {
				MemberMasterId int32 `xorm:"'member_master_id'"`
				Value          int32 `xorm:"'value'"`
			}
			bonuses := []memberBonusValue{}
			err = masterdata_db.Table("m_event_marathon_bonus_member").Where("event_marathon_master_id = ?", event.EventId).
				Find(&bonuses)
			utils.CheckErr(err)
			for _, bonus := range bonuses {
				eventMarathon.MemberBonus[bonus.MemberMasterId] = bonus.Value
			}
		}

		// generate the event_marathon_bonus_popup_order_member_mater_rows field, which are always sorted on member_matser_id (typo)
		for memberId := range eventMarathon.MemberBonus {
			eventMarathon.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Append(
				client.EventMarathonBonusPopupOrderMemberMaterRow{
					MemberMatserId: memberId,
					DisplayLine:    3,
					DisplayOrder:   memberId,
				})
		}
		sort.Slice(eventMarathon.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice, func(i, j int) bool {
			return eventMarathon.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice[i].DisplayOrder <
				eventMarathon.TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice[j].DisplayOrder
		})

		gamedata.EventMarathon[event.EventId] = &eventMarathon
	}
}

func init() {
	addLoadFunc(loadEventMarathon)
	addPrequisite(loadEventMarathon, loadCard)
	addPrequisite(loadEventMarathon, loadEventStory)
}
