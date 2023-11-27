package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UpdateTriggerCardGradeUp(trigger model.TriggerCardGradeUp) {
	session.UserTriggerCardGradeUpMapping.SetList(&session.UserModel.UserInfoTriggerCardGradeUpByTriggerID).Update(trigger)
}

// card grade up trigger is responsible for showing the pop-up animation when openning a card after getting a new copy
// or right after performing a limit break using items
// Getting a new trigger also destroy old trigger, and we might have to update it
func (session *Session) AddTriggerCardGradeUp(trigger model.TriggerCardGradeUp) {
	if trigger.TriggerID == 0 {
		trigger.TriggerID = session.TimeStamp + session.UniqueCount
		session.UniqueCount++
	}
	trigger.UserID = session.UserStatus.UserID
	// update the trigger
	session.UpdateTriggerCardGradeUp(trigger)
}

func triggerCardGradeUpFinalizer(session *Session) {
	for _, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerID.Objects {
		if !trigger.IsNull {
			dbTrigger := model.TriggerCardGradeUp{}
			exist, err := session.Db.Table("u_trigger_card_grade_up").
				Where("user_id = ? AND card_master_id = ?", trigger.UserID, trigger.CardMasterID).Get(&dbTrigger)
			utils.CheckErr(err)
			if exist { // if the card has a trigger, we have to remove it
				dbTrigger.IsNull = true
				session.UpdateTriggerCardGradeUp(dbTrigger)
				session.Db.Table("u_trigger_card_grade_up").
					Where("user_id = ? AND card_master_id = ?", trigger.UserID, trigger.CardMasterID).Delete(&dbTrigger)
			}
			trigger.BeforeLoveLevelLimit = trigger.AfterLoveLevelLimit
			// db trigger when login have BeforeLoveLevelLimit = AfterLoveLevelLimit
			// if the 2 numbers are equal the level up isn't shown when we open the card.
			_, err = session.Db.Table("u_trigger_card_grade_up").Insert(&trigger)
			utils.CheckErr(err)
		} else {
			// remove from db
			// this is only caused by infoTrigger/read
			_, err := session.Db.Table("u_trigger_card_grade_up").Where("trigger_id = ?", trigger.TriggerID).Delete(
				&model.TriggerCardGradeUp{})
			utils.CheckErr(err)
		}
	}
}

func (session *Session) UpdateTriggerBasic(trigger model.TriggerBasic) {
	session.UserTriggerBasicMapping.SetList(&session.UserModel.UserInfoTriggerBasicByTriggerID).Update(trigger)
}

func (session *Session) AddTriggerBasic(trigger model.TriggerBasic) {
	if trigger.TriggerID == 0 {
		trigger.TriggerID = session.TimeStamp + session.UniqueCount
		session.UniqueCount++
	}
	trigger.UserID = session.UserStatus.UserID
	session.UpdateTriggerBasic(trigger)
}

func triggerBasicFinalizer(session *Session) {
	for _, trigger := range session.UserModel.UserInfoTriggerBasicByTriggerID.Objects {
		if trigger.IsNull { // delete
			_, err := session.Db.Table("u_trigger_basic").Where("trigger_id = ?", trigger.TriggerID).Delete(
				&model.TriggerBasic{})
			utils.CheckErr(err)
		} else { // add
			_, err := session.Db.Table("u_trigger_basic").Insert(trigger)
			utils.CheckErr(err)
		}
	}
}

func triggerMemberGuildSupportItemExpiredFinalizer(session *Session) {
	for _, trigger := range session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID.Objects {
		if trigger.IsNull { // delete
			_, err := session.Db.Table("u_trigger_member_guild_support_item_expired").Where("trigger_id = ?", trigger.TriggerID).Delete(
				&model.TriggerMemberGuildSupportItemExpired{})
			utils.CheckErr(err)
		} else { // add
			_, err := session.Db.Table("u_trigger_member_guild_support_item_expired").Insert(trigger)
			utils.CheckErr(err)
		}
	}
}

func (session *Session) ReadMemberGuildSupportItemExpired() {
	err := session.Db.Table("u_trigger_member_guild_support_item_expired").
		Where("user_id = ?", session.UserStatus.UserID).
		Find(&session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID.Objects)
	utils.CheckErr(err)
	for i := range session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID.Objects {
		session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID.Objects[i].IsNull = true
	}
	// already marked as removed, the finalizer will take care of things
	// there's also no need to remove the item, the client won't show them if they're expired
}

// TODO: Trigger member love level up isn't really that persistent, so it's probably better to only keep it in ram
// This could be done by keeping a full user model in ram too.

func (session *Session) UpdateTriggerMemberLoveLevelUp(trigger model.TriggerMemberLoveLevelUp) {
	session.UserTriggerMemberLoveLevelUpMapping.SetList(&session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerID).Update(trigger)
}

func (session *Session) AddTriggerMemberLoveLevelUp(trigger model.TriggerMemberLoveLevelUp) {
	if trigger.TriggerID == 0 {
		trigger.TriggerID = session.TimeStamp + session.UniqueCount
		session.UniqueCount++
	}
	trigger.UserID = session.UserStatus.UserID
	session.UpdateTriggerMemberLoveLevelUp(trigger)
	if !trigger.IsNull {
		_, err := session.Db.Table("u_trigger_member_love_level_up").Insert(trigger)
		utils.CheckErr(err)
	} else {
		_, err := session.Db.Table("u_trigger_member_love_level_up").Where("trigger_id = ?", trigger.TriggerID).Delete(
			&model.TriggerMemberLoveLevelUp{})
		utils.CheckErr(err)
	}
}

func (session *Session) ReadAllMemberLoveLevelUpTriggers() {

	err := session.Db.Table("u_trigger_member_love_level_up").
		Where("user_id = ?", session.UserStatus.UserID).Find(&session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerID.Objects)
	utils.CheckErr(err)
	for i := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerID.Objects {
		session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerID.Objects[i].IsNull = true
	}
	_, err = session.Db.Table("u_trigger_member_love_level_up").Where("user_id = ?", session.UserStatus.UserID).Delete(
		&model.TriggerMemberLoveLevelUp{})
	utils.CheckErr(err)
}

func init() {
	addFinalizer(triggerCardGradeUpFinalizer)
	addFinalizer(triggerBasicFinalizer)
	addFinalizer(triggerMemberGuildSupportItemExpiredFinalizer)
	addGenericTableFieldPopulator("u_trigger_basic", "UserInfoTriggerBasicByTriggerID")
	addGenericTableFieldPopulator("u_trigger_card_grade_up", "UserInfoTriggerCardGradeUpByTriggerID")
	addGenericTableFieldPopulator("u_trigger_member_love_level_up", "UserInfoTriggerMemberLoveLevelUpByTriggerID")
	addGenericTableFieldPopulator("u_trigger_member_guild_support_item_expired", "UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID")
}
