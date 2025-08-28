package event_marathon_dev

import (
	"elichika/config"
	"elichika/enum"
	"elichika/router"
	"elichika/serverdata"
	"elichika/utils"

	"archive/zip"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// final page, we can extract the data here

func EventMarathonDevFinalGET(ctx *gin.Context) {
	archiveBuffer := new(bytes.Buffer)
	archiveWriter := zip.NewWriter(archiveBuffer)

	// generate the files
	main := serverdata.EventMarathon{
		EventId:             TopStatus.EventId,
		EventName:           EventName,
		BoosterItemId:       BoosterItemId,
		TitleImagePath:      TopStatus.TitleImagePath.V.ToPointer(),
		BackgroundImagePath: TopStatus.BackgroundImagePath.V.ToPointer(),
		BoardBaseImagePath:  TopStatus.BoardStatus.BoardBaseImagePath.V.ToPointer(),
		BoardDecoImagePath:  TopStatus.BoardStatus.BoardDecoImagePath.V.ToPointer(),
		BgmAssetPath:        TopStatus.BgmAssetPath.V.ToPointer(),
		GachaMasterId:       TopStatus.GachaMasterId,
	}
	for _, page := range TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice {
		main.RuleDescriptionPagesAssetPath = append(main.RuleDescriptionPagesAssetPath, page.ImageAssetPath.V.Value)
	}
	mainFile, err := archiveWriter.Create("main.json")
	utils.CheckErr(err)
	mainBytes, err := json.Marshal(main)
	utils.CheckErr(err)
	_, err = mainFile.Write(mainBytes)
	utils.CheckErr(err)

	boardBuffer := new(bytes.Buffer)
	boardWriter := csv.NewWriter(boardBuffer)
	err = boardWriter.Write([]string{"event_marathon_board_position_type", "position", "add_story_number", "priority", "image_thumbnail_asset_path"})
	utils.CheckErr(err)
	addStory := int32(0)
	for i, boardThing := range TopStatus.BoardStatus.BoardThingMasterRows.Slice {
		if boardThing.EventMarathonBoardPositionType == enum.EventMarathonBoardPositionTypePicture {
			addStory += 1
		}
		priority := int32(0)
		if i >= 3 {
			priority = int32(i - 2)
		}
		boardWriter.Write([]string{fmt.Sprint(boardThing.EventMarathonBoardPositionType), fmt.Sprint(boardThing.Position), fmt.Sprint(addStory), fmt.Sprint(priority), fmt.Sprint(boardThing.ImageThumbnailAssetPath.V.Value)})
	}
	boardWriter.Flush()
	boardFile, err := archiveWriter.Create("board.csv")
	utils.CheckErr(err)
	_, err = boardFile.Write(boardBuffer.Bytes())
	utils.CheckErr(err)

	// point rewards and ranking rewards are just lines that we can write directly
	pointRewardFile, err := archiveWriter.Create("point_reward.csv")
	utils.CheckErr(err)
	_, err = pointRewardFile.Write([]byte("point_required,content_type,content_id,content_amount"))
	utils.CheckErr(err)
	for _, line := range PointRewardLines {
		_, err = pointRewardFile.Write([]byte("\n" + line))
		utils.CheckErr(err)
	}
	rankingRewardFile, err := archiveWriter.Create("ranking_reward.csv")
	utils.CheckErr(err)
	_, err = rankingRewardFile.Write([]byte("highest_rank,content_type,content_id,content_amount,ranking_result_prize_type"))
	utils.CheckErr(err)
	for _, line := range RankingRewardLines {
		_, err = rankingRewardFile.Write([]byte("\n" + line))
		utils.CheckErr(err)
	}
	// TODO(event): mission
	missionFile, err := archiveWriter.Create("mission.csv")
	utils.CheckErr(err)
	_, err = missionFile.Write([]byte("mission_master_id"))
	utils.CheckErr(err)

	bonusPopupOrderFile, err := archiveWriter.Create("bonus_popup_order.csv")
	utils.CheckErr(err)
	_, err = bonusPopupOrderFile.Write([]byte("card_matser_id,display_line,display_order,is_gacha"))
	utils.CheckErr(err)
	toInt := func(b bool) int {
		if b {
			return 1
		} else {
			return 0
		}
	}
	for _, bonusCard := range TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Slice {
		_, err = bonusPopupOrderFile.Write([]byte(fmt.Sprintf("\n%d,%d,%d,%d", bonusCard.CardMatserId, bonusCard.DisplayLine, bonusCard.DisplayOrder, toInt(bonusCard.IsGacha))))
		utils.CheckErr(err)
	}

	rankingTopicRewardFile, err := archiveWriter.Create("ranking_topic_reward.csv")
	utils.CheckErr(err)
	_, err = rankingTopicRewardFile.Write([]byte("display_order,reward_card_id,reward_card_amount"))
	utils.CheckErr(err)
	for _, card := range TopStatus.EventRankingTopicRewardInfo.Slice {
		_, err = rankingTopicRewardFile.Write([]byte(fmt.Sprintf("\n%d,%d,1", card.DisplayOrder, card.RewardContent.ContentId)))
		utils.CheckErr(err)
	}

	totalTopicRewardFile, err := archiveWriter.Create("total_topic_reward.csv")
	utils.CheckErr(err)
	_, err = totalTopicRewardFile.Write([]byte("display_order,reward_card_id,reward_card_amount"))
	utils.CheckErr(err)
	for _, card := range TopStatus.EventTotalTopicRewardInfo.Slice {
		_, err = totalTopicRewardFile.Write([]byte(fmt.Sprintf("\n%d,%d,1", card.DisplayOrder, card.RewardContent.ContentId)))
		utils.CheckErr(err)
	}

	err = archiveWriter.Close()
	utils.CheckErr(err)
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%d.zip", TopStatus.EventId))
	ctx.Header("Content-Type", "application/zip")
	ctx.Writer.Write(archiveBuffer.Bytes())
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/final", EventMarathonDevFinalGET)
	}
}
