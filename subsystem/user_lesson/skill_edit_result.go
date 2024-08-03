package user_lesson

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"
)

func SkillEditResult(session *userdata.Session, req request.SkillEditResultRequest) {
	session.UserStatus.LessonResumeStatus = enum.TopPriorityProcessStatusNone

	_, err := session.Db.Table("u_lesson").Where("user_id = ?", session.UserId).Delete(&response.LessonResultResponse{})
	utils.CheckErr(err)

	for cardMasterId, selectedSkills := range req.SelectedSkillIds.Map {
		card := user_card.GetUserCard(session, cardMasterId)
		for i, skillId := range selectedSkills.Slice {
			switch i {
			case 0:
				card.AdditionalPassiveSkill1Id = skillId
			case 1:
				card.AdditionalPassiveSkill2Id = skillId
			case 2:
				card.AdditionalPassiveSkill3Id = skillId
			case 3:
				card.AdditionalPassiveSkill4Id = skillId
			}
		}
		user_card.UpdateUserCard(session, card) // this is always updated even if no skill change happen
	}

}
