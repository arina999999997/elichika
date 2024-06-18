package user_beginner_challenge

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

func GetBeginnerChallengeCells(session *userdata.Session) map[int32]*client.ChallengeCell {
	result := map[int32]*client.ChallengeCell{}
	err := session.Db.Table("u_beginner_challenge_cell").Where("user_id = ?", session.UserId).Find(&result)
	utils.CheckErr(err)
	// TODO(extra): This can reuse user_mission code, or both package can use another package all for goal initialization
	for _, cell := range session.Gamedata.BeginnerChallengeCell {
		_, exist := result[cell.Id]

		if !exist {
			result[cell.Id] = &client.ChallengeCell{
				CellId:           cell.Id,
				IsRewardReceived: false,
				Progress:         0,
			}
		}
		if result[cell.Id].IsRewardReceived {
			continue
		}
		changed := true
		switch cell.MissionClearConditionType {
		case enum.MissionClearConditionTypeClearedStoryMainChapter:
			chapterId := cell.MissionClearConditionCount
			requiredStoryCell := session.Gamedata.StoryMainChapter[chapterId].LastCellId
			userStoryMain := client.UserStoryMain{
				StoryMainMasterId: requiredStoryCell,
			}
			if userdata.GenericDatabaseExist(session, "u_story_main", userStoryMain) {
				result[cell.Id].Progress = chapterId
			} else {
				result[cell.Id].Progress = 0
			}
		case enum.MissionClearConditionTypeReadReferenceBook:
			// TODO(hardcored): this doesn't reference the masterdata database
			// so if/when the reference books are updated, there could be mistake
			count, err := session.Db.Table("u_reference_book").Where("user_id = ?", session.UserId).Count()
			utils.CheckErr(err)
			result[cell.Id].Progress = int32(count)
		case enum.MissionClearConditionTypeClearedEpisode:
			count, err := session.Db.Table("u_story_member").Where("user_id = ?", session.UserId).Count()
			utils.CheckErr(err)
			result[cell.Id].Progress = int32(count)
		case enum.MissionClearConditionTypeMemberLovePanelCell:
			loveMemberPanels := []client.MemberLovePanel{}
			err := session.Db.Table("u_member_love_panel").Where("user_id = ?", session.UserId).Find(&loveMemberPanels)
			utils.CheckErr(err)
			count := 0
			for _, panel := range loveMemberPanels {
				count += panel.MemberLovePanelCellIds.Size()
			}
			result[cell.Id].Progress = int32(count)
		case enum.MissionClearConditionTypeClearedStorySide:
			count, err := session.Db.Table("u_story_side").Where("user_id = ? AND is_new = 0", session.UserId).Count()
			utils.CheckErr(err)
			result[cell.Id].Progress = int32(count)
		default:
			changed = false
			// fmt.Println("default type: ", cell.MissionClearConditionType)
		}
		if changed || (!exist) {
			UpdateChallengeCell(session, *result[cell.Id])
		}
	}
	return result
}
