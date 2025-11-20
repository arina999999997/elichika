//go:build dev

package event_mining_dev

import (
	"elichika/config"
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
func intBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
func nullString(ptr *string) string {
	if ptr == nil {
		return "null"
	}
	return *ptr
}
func EventMiningDevFinalGET(ctx *gin.Context) {
	archiveBuffer := new(bytes.Buffer)
	archiveWriter := zip.NewWriter(archiveBuffer)

	// generate the files
	main := serverdata.EventMining{
		EventId:                            TopStatus.EventId,
		TitleImagePath:                     TopStatus.TitleImagePath.V.Value,
		BackgroundImagePath:                TopStatus.BackgroundImagePath.V.Value,
		BgmAssetPath:                       TopStatus.BgmAssetPath.V.Value,
		GachaMasterId:                      TopStatus.GachaMasterId,
		VoltageRankingTopicRewardAssetPath: TopStatus.EventVoltageRankingTopicRewardInfo.Slice[0].SinglePictureAssetPath.V.Value,
	}
	for _, live := range TopStatus.EventMiningCompetitionMasterRows.Slice {
		main.EventCompetitionLiveIds = append(main.EventCompetitionLiveIds, live.LiveId)
	}
	// TODO(hardcoded): this can be 5 for earlier events
	// but because we don't have assets from there, let's just hard code it to 3 and maybe change it later
	main.EventCompetitionSelectionAmount = 3
	for _, page := range TopStatus.EventMiningRuleDescriptionPageMasterRows.Slice {
		main.RuleDescriptionPagesAssetPath = append(main.RuleDescriptionPagesAssetPath, page.ImageAssetPath.V.Value)
	}
	mainFile, err := archiveWriter.Create("main.json")
	utils.CheckErr(err)
	mainBytes, err := json.Marshal(main)
	utils.CheckErr(err)
	_, err = mainFile.Write(mainBytes)
	utils.CheckErr(err)

	pointRankingRewardFile, err := archiveWriter.Create("point_ranking_reward.csv")
	utils.CheckErr(err)
	_, err = pointRankingRewardFile.Write([]byte("highest_rank,content_type,content_id,content_amount,ranking_result_prize_type\n" + pointRankingRewardCsv))
	utils.CheckErr(err)
	pointRankingTopicRewardFile, err := archiveWriter.Create("point_ranking_topic_reward.csv")
	utils.CheckErr(err)
	_, err = pointRankingTopicRewardFile.Write([]byte("display_order,reward_card_id,reward_card_amount"))
	utils.CheckErr(err)
	for _, card := range TopStatus.EventPointRankingTopicRewardInfo.Slice {
		_, err = pointRankingTopicRewardFile.Write([]byte(fmt.Sprintf("\n%d,%d,1", card.DisplayOrder, card.RewardContent.ContentId)))
		utils.CheckErr(err)
	}

	voltageRankingRewardFile, err := archiveWriter.Create("voltage_ranking_reward.csv")
	utils.CheckErr(err)
	_, err = voltageRankingRewardFile.Write([]byte("highest_rank,content_type,content_id,content_amount,ranking_result_prize_type\n" + voltageRankingRewardCsv))
	utils.CheckErr(err)

	// // TODO(event): mission
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
	for _, bonusCard := range TopStatus.EventMiningBonusPopupOrderCardMaterRows.Slice {
		_, err = bonusPopupOrderFile.Write([]byte(fmt.Sprintf("\n%d,%d,%d,%d", bonusCard.CardMatserId, bonusCard.DisplayLine, bonusCard.DisplayOrder, toInt(bonusCard.IsGacha))))
		utils.CheckErr(err)
	}

	tradeProductFile, err := archiveWriter.Create("trade_product.csv")
	utils.CheckErr(err)
	_, err = tradeProductFile.Write([]byte("product_id,source_amount,stock_amount,content_type,content_id,content_amount"))
	utils.CheckErr(err)
	for i, tradeProduct := range tradeProducts {
		_, err = tradeProductFile.Write([]byte(fmt.Sprintf("\n%d,%d,%d,%d,%d,%d",
			TopStatus.EventId*1000+int32(i)+1,
			tradeProduct.Price,
			tradeProduct.Stock,
			tradeProduct.ContentType,
			tradeProduct.ContentId,
			tradeProduct.ContentAmount)))
		utils.CheckErr(err)
	}

	topStillCellBuffer := new(bytes.Buffer)
	topStillCellWriter := csv.NewWriter(topStillCellBuffer)
	err = topStillCellWriter.Write([]string{
		"thumbnail_cell_id",
		"add_story_number",
		"priority",
		"image_thumbnail_asset_path",
		"is_popup",
		"popup_image_thumbnail_asset_path",
		"popup_comment",
		"is_big"})
	utils.CheckErr(err)
	for _, cell := range TopStatus.EventMiningTopStillCellMasterRows.Slice {
		err = topStillCellWriter.Write([]string{
			fmt.Sprint(cell.ThumbnailCellId),
			fmt.Sprint(cell.AddStoryNumber),
			fmt.Sprint(cell.Priority),
			fmt.Sprint(cell.ImageThumbnailAssetPath.V.Value),
			fmt.Sprint(intBool(cell.IsPopup)),
			fmt.Sprint(cell.PopupImageThumbnailAssetPath.V.Value),
			fmt.Sprint("null"), // TODO(extra): figure it out if it is ever not null
			fmt.Sprint(intBool(cell.IsBig)),
		})
		utils.CheckErr(err)
	}
	topStillCellWriter.Flush()
	topStillCellFile, err := archiveWriter.Create("top_still_cell.csv")
	utils.CheckErr(err)
	_, err = topStillCellFile.Write(topStillCellBuffer.Bytes())
	utils.CheckErr(err)

	topStillSubCellBuffer := new(bytes.Buffer)
	topStillSubCellWriter := csv.NewWriter(topStillSubCellBuffer)
	err = topStillSubCellWriter.Write([]string{
		"panel_set_id",
		"thumbnail_cell_id",
		"priority",
		"image_thumbnail_asset_path",
		"is_popup",
		"popup_image_thumbnail_asset_path",
	})
	utils.CheckErr(err)
	for _, cell := range TopStatus.EventMiningTopStillSubCellMasterRows.Slice {
		err = topStillSubCellWriter.Write([]string{
			"2", // TODO(extra): doesn't really matter
			fmt.Sprint(cell.ThumbnailCellId),
			fmt.Sprint(cell.Priority),
			fmt.Sprint(cell.ImageThumbnailAssetPath.V.Value),
			fmt.Sprint(intBool(cell.IsPopup)),
			fmt.Sprint(cell.PopupImageThumbnailAssetPath.V.Value),
		})
		utils.CheckErr(err)
	}
	topStillSubCellWriter.Flush()
	topStillSubCellFile, err := archiveWriter.Create("top_still_sub_cell.csv")
	utils.CheckErr(err)
	_, err = topStillSubCellFile.Write(topStillSubCellBuffer.Bytes())
	utils.CheckErr(err)

	err = archiveWriter.Close()
	utils.CheckErr(err)
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%d.zip", TopStatus.EventId))
	ctx.Header("Content-Type", "application/zip")
	ctx.Writer.Write(archiveBuffer.Bytes())
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/final", EventMiningDevFinalGET)
	}
}
