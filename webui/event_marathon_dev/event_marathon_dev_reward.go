package event_marathon_dev

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/utils"

	"strconv"
	"strings"
)

func BuildRewardObjects() {
	// build the reward object
	TopStatus.EventMarathonRewardMasterRows.Slice = nil
	TopStatus.EventMarathonPointRewardMasterRows.Slice = nil
	TopStatus.EventMarathonRankingRewardMasterRows.Slice = nil
	for i := range PointRewardLines {
		if PointRewardLines[i] == "unknown_reward_name" {
			break
		}
		parts := strings.Split(PointRewardLines[i], ",")
		pointRequired, err := strconv.Atoi(parts[0])
		utils.CheckErr(err)
		contentType, err := strconv.Atoi(parts[1])
		utils.CheckErr(err)
		contentId, err := strconv.Atoi(parts[2])
		utils.CheckErr(err)
		contentAmount, err := strconv.Atoi(parts[3])
		utils.CheckErr(err)
		TopStatus.EventMarathonRewardMasterRows.Append(client.EventMarathonRewardMasterRow{
			RewardGroupId: TopStatus.EventId*10000 + int32(i) + 1,
			RewardContent: client.Content{
				ContentType:   int32(contentType),
				ContentId:     int32(contentId),
				ContentAmount: int32(contentAmount),
			},
			DisplayOrder: 0,
		})
		TopStatus.EventMarathonPointRewardMasterRows.Append(client.EventMarathonPointRewardMasterRow{
			RequiredPoint: int32(pointRequired),
			RewardGroupId: TopStatus.EventId*10000 + int32(i) + 1,
		})
	}
	rewardGroupId := TopStatus.EventId*10000 + 1000
	rewardGroupCount := 0
	displayOrder := int32(0)
	for i := range RankingRewardLines {
		tokens := strings.Split(RankingRewardLines[i], ",")
		highestRank, err := strconv.Atoi(tokens[0])
		utils.CheckErr(err)
		contentType, err := strconv.Atoi(tokens[1])
		utils.CheckErr(err)
		contentId, err := strconv.Atoi(tokens[2])
		utils.CheckErr(err)
		contentAmount, err := strconv.Atoi(tokens[3])
		utils.CheckErr(err)
		prizeType, err := strconv.Atoi(tokens[4])
		utils.CheckErr(err)
		// new reward group
		if int32(contentType) == enum.ContentTypeEmblem {
			rewardGroupId++
			rewardGroupCount++
			TopStatus.EventMarathonRankingRewardMasterRows.Append(client.EventMarathonRankingRewardMasterRow{
				RankingRewardMasterId:  rewardGroupId,
				UpperRank:              int32(highestRank),
				RewardGroupId:          rewardGroupId,
				RankingResultPrizeType: int32(prizeType),
			})
			if rewardGroupCount > 1 {
				TopStatus.EventMarathonRankingRewardMasterRows.Slice[rewardGroupCount-2].LowerRank = generic.NewNullable[int32](int32(highestRank - 1))
			}
			// bigger displayOrder means showing up first, kinda weird
			displayOrder = 7
		} else {
			displayOrder--
		}
		TopStatus.EventMarathonRewardMasterRows.Append(client.EventMarathonRewardMasterRow{
			RewardGroupId: rewardGroupId,
			RewardContent: client.Content{
				ContentType:   int32(contentType),
				ContentId:     int32(contentId),
				ContentAmount: int32(contentAmount),
			},
			DisplayOrder: displayOrder,
		})
	}
}
