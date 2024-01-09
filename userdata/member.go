package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/utils"
)

func (session *Session) GetMember(memberMasterId int32) client.UserMember {
	pos, exist := session.UserMemberMapping.SetList(&session.UserModel.UserMemberByMemberId).Map[int64(memberMasterId)]
	if exist {
		return session.UserModel.UserMemberByMemberId.Objects[pos]
	}
	member := client.UserMember{}
	exist, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserId, memberMasterId).Get(&member)
	utils.CheckErr(err)
	if !exist {
		// always inserted at login if not exist
		panic("member not found")
	}
	return member
}

func (session *Session) UpdateMember(member client.UserMember) {
	session.UserMemberMapping.SetList(&session.UserModel.UserMemberByMemberId).Update(member)
}

func (session *Session) InsertMembers(members []client.UserMember) {
	for _, member := range members {
		session.UpdateMember(member)
	}
}

func memberFinalizer(session *Session) {
	for _, member := range session.UserModel.UserMemberByMemberId.Objects {
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserId, member.MemberMasterId).AllCols().Update(member)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_member", member)
		}
	}
}

func (session *Session) GetUserCommunicationMemberDetailBadge(memberMasterId int) client.UserCommunicationMemberDetailBadge {
	pos, exist := session.UserCommunicationMemberDetailBadgeMapping.
		SetList(&session.UserModel.UserCommunicationMemberDetailBadgeById).Map[int64(memberMasterId)]
	if exist {
		return session.UserModel.UserCommunicationMemberDetailBadgeById.Objects[pos]
	}
	badge := client.UserCommunicationMemberDetailBadge{}
	exist, err := session.Db.Table("u_communication_member_detail_badge").
		Where("user_id = ? AND member_master_id = ?", session.UserId, memberMasterId).Get(&badge)
	utils.CheckErr(err)
	if !exist {
		// always inserted at login if not exist
		panic("member not found")
	}
	return badge
}

func (session *Session) UpdateUserCommunicationMemberDetailBadge(badge client.UserCommunicationMemberDetailBadge) {
	session.UserCommunicationMemberDetailBadgeMapping.
		SetList(&session.UserModel.UserCommunicationMemberDetailBadgeById).Update(badge)
}

func communicationMemberDetailBadgeFinalizer(session *Session) {
	// TODO: this is only handled on the read side, new items won't change the badge
	for _, detailBadge := range session.UserModel.UserCommunicationMemberDetailBadgeById.Objects {
		affected, err := session.Db.Table("u_communication_member_detail_badge").
			Where("user_id = ? AND member_master_id = ?", session.UserId, detailBadge.MemberMasterId).
			AllCols().Update(detailBadge)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_communication_member_detail_badge", detailBadge)
		}
	}
}

// add love point and return the love point added (in case maxed out)
func (session *Session) AddLovePoint(memberId, point int32) int32 {
	member := session.GetMember(memberId)
	if point > member.LovePointLimit-member.LovePoint {
		point = member.LovePointLimit - member.LovePoint
	}
	member.LovePoint += point

	oldLoveLevel := member.LoveLevel
	member.LoveLevel = session.Gamedata.LoveLevelFromLovePoint(member.LovePoint)
	// unlock bond stories, unlock bond board
	if oldLoveLevel < member.LoveLevel {
		gamedata := session.Ctx.MustGet("gamedata").(*gamedata.Gamedata)
		masterMember := gamedata.Member[memberId]
		for loveLevel := oldLoveLevel + 1; loveLevel <= member.LoveLevel; loveLevel++ {
			for _, reward := range masterMember.LoveLevelRewards[loveLevel] {
				session.AddResource(reward)
			}
		}

		currentLovePanel := session.GetMemberLovePanel(memberId)
		if len(currentLovePanel.LovePanelLastLevelCellIds) == 5 {
			// the current bond board is full, check if we can unlock a new bond board
			masterLovePanel := gamedata.MemberLovePanelCell[currentLovePanel.LovePanelLastLevelCellIds[0]].MemberLovePanel
			if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
				// TODO: remove magic id from love panel system
				currentLovePanel.LevelUp()
				session.AddTriggerBasic(client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
					ParamInt:        generic.NewNullable(int32(currentLovePanel.LovePanelLevel*1000 + currentLovePanel.MemberId))})

				session.UpdateMemberLovePanel(currentLovePanel)
			}
		}
		session.AddTriggerMemberLoveLevelUp(client.UserInfoTriggerMemberLoveLevelUp{
			MemberMasterId:  memberId,
			BeforeLoveLevel: member.LoveLevel - 1})

	}
	session.UpdateMember(member)
	return point
}

func init() {
	addGenericTableFieldPopulator("u_member", "UserMemberByMemberId")
	addFinalizer(memberFinalizer)
	addGenericTableFieldPopulator("u_communication_member_detail_badge", "UserCommunicationMemberDetailBadgeById")
	addFinalizer(communicationMemberDetailBadgeFinalizer)
}
