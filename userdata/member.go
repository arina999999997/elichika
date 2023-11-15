package userdata

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
	"elichika/utils"

	"fmt"
)

func (session *Session) GetMember(memberMasterID int) model.UserMember {
	member, exist := session.UserMemberDiffs[memberMasterID]
	if exist {
		return member
	}
	exists, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, memberMasterID).Get(&member)
	// inserted at login if not exist
	utils.CheckErr(err)
	if !exists {
		panic("member not found")
	}
	return member
}

func (session *Session) GetAllMembers() []model.UserMember {
	members := []model.UserMember{}
	err := session.Db.Table("u_member").Where("user_id = ?", session.UserStatus.UserID).Find(&members)
	if err != nil {
		panic(err)
	}
	return members
}

func (session *Session) UpdateMember(member model.UserMember) {
	session.UserMemberDiffs[member.MemberMasterID] = member
}

func (session *Session) InsertMembers(members []model.UserMember) {
	affected, err := session.Db.Table("u_member").Insert(&members)
	utils.CheckErr(err)
	fmt.Println("Inserted ", affected, " members")
}

func (session *Session) FinalizeUserMemberDiffs() []any {
	userMemberByMemberID := []any{}
	for memberMasterID, member := range session.UserMemberDiffs {
		session.UserModelCommon.UserMemberByMemberID.PushBack(member)
		userMemberByMemberID = append(userMemberByMemberID, memberMasterID)
		userMemberByMemberID = append(userMemberByMemberID, member)
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, memberMasterID).AllCols().Update(member)
		if (err != nil) || (affected != 1) {
			panic(err)
		}
	}
	return userMemberByMemberID
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
			session.AddResource(masterMember.LoveLevelRewards[loveLevel])
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
