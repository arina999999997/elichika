package userdata

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetMember(memberMasterID int) model.UserMember {
	pos, exist := session.UserMemberMapping.SetList(&session.UserModel.UserMemberByMemberID).Map[int64(memberMasterID)]
	if exist {
		return session.UserModel.UserMemberByMemberID.Objects[pos]
	}
	member := model.UserMember{}
	exist, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, memberMasterID).Get(&member)
	utils.CheckErr(err)
	if !exist {
		// always inserted at login if not exist
		panic("member not found")
	}
	return member
}

func (session *Session) UpdateMember(member model.UserMember) {
	session.UserMemberMapping.SetList(&session.UserModel.UserMemberByMemberID).Update(member)
}

func (session *Session) InsertMembers(members []model.UserMember) {
	for _, member := range members {
		session.UpdateMember(member)
	}
}

func memberFinalizer(session *Session) {
	for _, member := range session.UserModel.UserMemberByMemberID.Objects {
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, member.MemberMasterID).AllCols().Update(member)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_member").Insert(member)
			utils.CheckErr(err)
		}
	}
}

func (session *Session) GetUserCommunicationMemberDetailBadge(memberMasterID int) model.UserCommunicationMemberDetailBadge {
	pos, exist := session.UserCommunicationMemberDetailBadgeMapping.
		SetList(&session.UserModel.UserCommunicationMemberDetailBadgeByID).Map[int64(memberMasterID)]
	if exist {
		return session.UserModel.UserCommunicationMemberDetailBadgeByID.Objects[pos]
	}
	badge := model.UserCommunicationMemberDetailBadge{}
	exist, err := session.Db.Table("u_communication_member_detail_badge").
		Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, memberMasterID).Get(&badge)
	utils.CheckErr(err)
	if !exist {
		// always inserted at login if not exist
		panic("member not found")
	}
	return badge
}

func (session *Session) UpdateUserCommunicationMemberDetailBadge(badge model.UserCommunicationMemberDetailBadge) {
	session.UserCommunicationMemberDetailBadgeMapping.
		SetList(&session.UserModel.UserCommunicationMemberDetailBadgeByID).Update(badge)
}

func communicationMemberDetailBadgeFinalizer(session *Session) {
	// TODO: this is only handled on the read side, new items won't change the badge
	for _, detailBadge := range session.UserModel.UserCommunicationMemberDetailBadgeByID.Objects {
		affected, err := session.Db.Table("u_communication_member_detail_badge").
			Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, detailBadge.MemberMasterID).
			AllCols().Update(detailBadge)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_communication_member_detail_badge").Insert(detailBadge)
			utils.CheckErr(err)
		}
	}
}

// add love point and return the love point added (in case maxed out)
func (session *Session) AddLovePoint(memberID, point int) int {
	member := session.GetMember(memberID)
	if point > member.LovePointLimit-member.LovePoint {
		point = member.LovePointLimit - member.LovePoint
	}
	member.LovePoint += point

	oldLoveLevel := member.LoveLevel
	member.LoveLevel = session.Gamedata.LoveLevelFromLovePoint(member.LovePoint)
	// unlock bond stories, unlock bond board
	if oldLoveLevel < member.LoveLevel {
		gamedata := session.Ctx.MustGet("gamedata").(*gamedata.Gamedata)
		masterMember := gamedata.Member[memberID]
		for loveLevel := oldLoveLevel + 1; loveLevel <= member.LoveLevel; loveLevel++ {
			for _, reward := range masterMember.LoveLevelRewards[loveLevel] {
				session.AddResource(reward)
			}
		}

		currentLovePanel := session.GetMemberLovePanel(memberID)
		if len(currentLovePanel.LovePanelLastLevelCellIDs) == 5 {
			// the current bond board is full, check if we can unlock a new bond board
			masterLovePanel := gamedata.MemberLovePanelCell[currentLovePanel.LovePanelLastLevelCellIDs[0]].MemberLovePanel
			if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
				// TODO: remove magic id from love panel system
				currentLovePanel.LevelUp()
				session.AddTriggerBasic(model.TriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeUnlockBondBoard,
					ParamInt:        currentLovePanel.LovePanelLevel*1000 + currentLovePanel.MemberID})

				session.UpdateMemberLovePanel(currentLovePanel)
			}
		}
		session.AddTriggerMemberLoveLevelUp(model.TriggerMemberLoveLevelUp{
			MemberMasterID:  memberID,
			BeforeLoveLevel: member.LoveLevel - 1})

	}
	session.UpdateMember(member)
	return point
}

func init() {
	addGenericTableFieldPopulator("u_member", "UserMemberByMemberID")
	addFinalizer(memberFinalizer)
	addGenericTableFieldPopulator("u_communication_member_detail_badge", "UserCommunicationMemberDetailBadgeByID")
	addFinalizer(communicationMemberDetailBadgeFinalizer)
}
