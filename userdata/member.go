package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/utils"

	"fmt"
)

func (session *Session) GetMember(memberMasterId int32) client.UserMember {
	ptr, exist := session.UserModel.UserMemberByMemberId.Get(memberMasterId)
	if exist {
		return *ptr
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
	session.UserModel.UserMemberByMemberId.Set(member.MemberMasterId, member)
}

func (session *Session) InsertMembers(members []client.UserMember) {
	for _, member := range members {
		session.UpdateMember(member)
	}
}

func memberFinalizer(session *Session) {
	for _, member := range session.UserModel.UserMemberByMemberId.Map {
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserId, member.MemberMasterId).AllCols().Update(*member)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_member", *member)
		}
	}
}

func (session *Session) GetUserCommunicationMemberDetailBadge(memberMasterId int32) client.UserCommunicationMemberDetailBadge {
	ptr, exist := session.UserModel.UserCommunicationMemberDetailBadgeById.Get(memberMasterId)
	if exist {
		return *ptr
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
	session.UserModel.UserCommunicationMemberDetailBadgeById.Set(badge.MemberMasterId, badge)
}

func communicationMemberDetailBadgeFinalizer(session *Session) {
	// TODO: this is only handled on the read side, new items won't change the badge
	for _, detailBadge := range session.UserModel.UserCommunicationMemberDetailBadgeById.Map {
		affected, err := session.Db.Table("u_communication_member_detail_badge").
			Where("user_id = ? AND member_master_id = ?", session.UserId, detailBadge.MemberMasterId).
			AllCols().Update(*detailBadge)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_communication_member_detail_badge", *detailBadge)
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
		masterMember := session.Gamedata.Member[memberId]
		for loveLevel := oldLoveLevel + 1; loveLevel <= member.LoveLevel; loveLevel++ {
			for _, reward := range masterMember.LoveLevelRewards[loveLevel] {
				session.AddResource(reward)
			}
		}

		currentLovePanel := session.GetMemberLovePanel(memberId)
		unlockCount := currentLovePanel.MemberLovePanelCellIds.Size()
		if (unlockCount > 0) && (unlockCount%5 == 0) {
			// love panel is maxed out
			lastCell := currentLovePanel.MemberLovePanelCellIds.Slice[unlockCount-1]
			masterLovePanel := session.Gamedata.MemberLovePanelCell[lastCell].MemberLovePanel

			if (masterLovePanel.NextPanel != nil) &&
				(masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel > oldLoveLevel) {
				nextPanel := masterLovePanel.NextPanel
				fmt.Println(nextPanel.Id)
				session.AddTriggerBasic(client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
					ParamInt:        generic.NewNullable(nextPanel.Id)})
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
	addFinalizer(memberFinalizer)
	addFinalizer(communicationMemberDetailBadgeFinalizer)
}
