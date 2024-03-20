package user_lesson

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/generic/drop"
	"elichika/item"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_lesson_deck"
	"elichika/subsystem/user_member_guild"
	"elichika/subsystem/user_mission"
	"elichika/subsystem/user_subscription_status"
	"elichika/userdata"

	"reflect"
	"sort"
)

// handle the lesson and write the result to the database
// drop is calculated using the following process:
// - First iterate over the lesson menu in the order sent, let's say A, B, C, and get a random drop count for each of them
// - Then generate the items using a generic.random list.
//   - We pick the list based on whether the user has the training enhancing items requested
//
// - Finally there is a chance to add megaphones, if it's applicable.
// - For 3 times, the order of A B C is not preserved for the later runs, instead it's sorted
//   - Don't really know why this is the case
//   - One plausible theory is that they sorted the list to use it for insight skills as the order doesn't matter
//
// - the amount of drop is assumed to be the following per instace of lesson (9 in total for x3), start with 15 and go up to 25
//   - 0.25
//   - 0.08
//   - 0.08
//   - 0.25
//   - 0.08
//   - 0.08
//   - 0.11
//   - 0.01
//   - 0.01
//   - 0.03
//   - 0.01
//   - 0.01
//
// - the amount of megaphone drop is assumed to be the following per instance of lesson menu, starting with 0, end with 3
//   - 0.81
//   - 0.1
//   - 0.075
//   - 0.015
//
// TODO(hard_coded): Maybe this should be in the database
var (
	dropCountList          drop.WeightedDropList[int32]
	megaphoneDropCountList drop.WeightedDropList[int32]
)

func init() {
	dropCountList.AddItem(15, 25)
	dropCountList.AddItem(16, 8)
	dropCountList.AddItem(17, 8)
	dropCountList.AddItem(18, 25)
	dropCountList.AddItem(19, 8)
	dropCountList.AddItem(20, 8)
	dropCountList.AddItem(21, 11)
	dropCountList.AddItem(22, 1)
	dropCountList.AddItem(23, 1)
	dropCountList.AddItem(24, 3)
	dropCountList.AddItem(25, 1)
	dropCountList.AddItem(26, 1)

	megaphoneDropCountList.AddItem(0, 810)
	megaphoneDropCountList.AddItem(1, 100)
	megaphoneDropCountList.AddItem(2, 75)
	megaphoneDropCountList.AddItem(3, 15)
}

func ExecuteLesson(session *userdata.Session, req request.ExecuteLessonRequest) response.ExecuteLessonResponse {
	resp := response.ExecuteLessonResponse{
		UserModelDiff: &session.UserModel,
	}

	result := response.LessonResultResponse{
		SelectedDeckId: req.SelectedDeckId,
	}

	deck := user_lesson_deck.GetUserLessonDeck(session, req.SelectedDeckId)
	repeatCount := 1
	if req.IsThreeTimes {
		repeatCount = 3
	}

	// update mission progress
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountLesson, nil, nil,
		func(session *userdata.Session, missionList []any, _ ...any) {
			for _, mission := range missionList {
				user_mission.AddMissionProgress(session, mission, int32(repeatCount))
			}
		})

	session.UserStatus.LessonResumeStatus = enum.TopPriorityProcessStatusLesson

	enhancingItems := map[int32]*client.Content{}

	for _, itemId := range req.ConsumedContentIds.Slice {
		item := user_content.GetUserContent(session, enum.ContentTypeLessonEnhancingItem, itemId)
		enhancingItems[itemId] = &item
	}

	resp.IsSubscription = user_subscription_status.HasSubscription(session)

	for lesson := int32(1); lesson <= 4; lesson++ {
		actions := generic.List[client.LessonMenuAction]{}
		for i := 1; i <= 9; i++ {
			cardMasterId := reflect.ValueOf(deck).Field(i + 1).Interface().(generic.Nullable[int32]).Value
			actions.Append(client.LessonMenuAction{
				CardMasterId: cardMasterId,
				Position:     int32(i),
			})
		}
		resp.LessonMenuActions.Set(lesson%4, actions)
		resp.LessonDropRarityList.Set(lesson%4, generic.List[int32]{})
	}

	isMemberGuildRankingPeriod := user_member_guild.IsMemberGuildRankingPeriod(session)

	for repeat := 1; repeat <= repeatCount; repeat++ {
		usedItems := []int32{}
		for _, itemId := range req.ConsumedContentIds.Slice {
			if enhancingItems[itemId].ContentAmount > 0 {
				enhancingItems[itemId].ContentAmount--
				usedItems = append(usedItems, itemId)
			}
		}

		// handle skill here if we want
		gainedItems := []client.LessonDropItem{}

		// use default drop, but switch to other drop if necessary

		for lesson := int32(1); lesson <= 3; lesson++ {
			lessonMenu := session.Gamedata.LessonMenu[req.ExecuteLessonIds.Slice[lesson-1]]
			dropList := lessonMenu.DefaultDrop
			for _, item := range usedItems {
				drop, exist := lessonMenu.Drop[item]
				if exist {
					dropList = drop
				}
			}

			dropCount := dropCountList.GetRandomItem()
			gainedRarity := []int32{}

			dropRarityList := resp.LessonDropRarityList.GetOnly(lesson)
			for i := int32(0); i < dropCount; i++ {
				drop := dropList.GetRandomItem()
				if drop.DropRarity > enum.LessonDropRarityTypeRare1 {
					gainedRarity = append(gainedRarity, enum.LessonDropRarityTypeRare2)
				} else {
					gainedRarity = append(gainedRarity, enum.LessonDropRarityTypeRare1)
				}
				gainedItems = append(gainedItems, drop)
			}

			// megaphone, only drop when ranking is on
			if isMemberGuildRankingPeriod {
				megaphoneDrop := megaphoneDropCountList.GetRandomItem()
				for i := int32(0); i < megaphoneDrop; i++ {
					gainedItems = append(gainedItems, client.LessonDropItem{
						ContentType:   item.RallyMegaphone.ContentType,
						ContentId:     item.RallyMegaphone.ContentId,
						ContentAmount: item.RallyMegaphone.ContentAmount,
						DropRarity:    4, // this field is not enum
					})
					gainedRarity = append(gainedRarity, enum.LessonDropRarityTypeRare2)
				}
			}

			for _, content := range gainedItems {
				user_content.AddContent(session, client.Content{
					ContentType:   content.ContentType,
					ContentId:     content.ContentId,
					ContentAmount: content.ContentAmount,
				})
			}
			for _, rarity := range gainedRarity {
				dropRarityList.Append(rarity)
			}

			if resp.IsSubscription {
				for _, rarity := range gainedRarity {
					dropRarityList.Append(rarity)
				}
				for _, content := range gainedItems {
					user_content.AddContent(session, client.Content{
						ContentType:   content.ContentType,
						ContentId:     content.ContentId,
						ContentAmount: content.ContentAmount,
					})
				}
			}
		}

		for _, drop := range gainedItems {
			result.DropItemList.Append(drop)
		}

		if resp.IsSubscription {
			for _, drop := range gainedItems {
				drop.IsSubscription = true
				result.DropItemList.Append(drop)
			}
		}

		if (repeat == 1) && (repeat < repeatCount) {
			sort.Slice(req.ExecuteLessonIds.Slice, func(i, j int) bool {
				return req.ExecuteLessonIds.Slice[i] < req.ExecuteLessonIds.Slice[j]
			})
		}
	}

	for _, item := range enhancingItems {
		user_content.UpdateUserContent(session, *item)
	}

	// insight skills

	// can only return 12 max
	skills := []int32{
		// 30000041, // Appeal+ (L)
		30000482, // Appeal+ (M):Group
		30000517, // Appeal+ (M):Same Attribute
		30000502, // Appeal+ (M):Same School
		30000507, // Appeal+ (M):Same Strategy
		30000492, // Appeal+ (M):Same Year
		30000512, // Appeal+ (M):Type
		30000044, // Skill Activation %+ (L)
		30000485, // Skill Activation %+ (M):Group
		30000520, // Skill Activation %+ (M):Same Attribute
		// 30000505, // Skill Activation %+ (M):Same School
		30000510, // Skill Activation %+ (M):Same Strategy
		30000495, // Skill Activation %+ (M):Same Year
		// 30000515, // Skill Activation %+ (M):Type
		30000045, // Type Effect+ (M)
	}
	for position := int32(1); position <= 9; position++ {
		for _, skillId := range skills {
			result.DropSkillList.Append(client.LessonResultDropPassiveSkill{
				Position:       position,
				PassiveSkillId: skillId,
			})
		}
	}
	userdata.GenericDatabaseInsert(session, "u_lesson", result)

	return resp
}
