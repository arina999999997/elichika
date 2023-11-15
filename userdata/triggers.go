package userdata

import (
	"elichika/generic"
	"elichika/model"
	"elichika/utils"

	"time"
)

// card grade up trigger is responsible for showing the pop-up animation when openning a card after getting a new copy
// or right after performing a limit break using items
// Getting a new trigger also destroy old trigger, and we might have to update it
func (session *Session) AddTriggerCardGradeUp(trigger model.TriggerCardGradeUp) {
	if trigger.TriggerID == 0 {
		trigger.TriggerID = time.Now().UnixNano()
	}
	trigger.UserID = session.UserStatus.UserID
	if !trigger.IsNull {
		dbTrigger := model.TriggerCardGradeUp{}

		exists, err := session.Db.Table("u_trigger_card_grade_up").
			Where("user_id = ? AND card_master_id = ?", trigger.UserID, trigger.CardMasterID).Get(&dbTrigger)
		utils.CheckErr(err)
		currentPos := -1
		userModelCommonPos := -1
		if exists { // if the card has a trigger, we have to remove it
			dbTrigger.IsNull = true
			session.Db.Table("u_trigger_card_grade_up").
				Where("user_id = ? AND card_master_id = ?", trigger.UserID, trigger.CardMasterID).Delete(&dbTrigger)
			// make the client remove the trigger
			for i, _ := range session.TriggerCardGradeUps {
				if i%2 == 0 {
					if session.TriggerCardGradeUps[i].(int64) == dbTrigger.TriggerID {
						currentPos = i
						break
					}
				}
			}
			for i, obj := range session.UserModelCommon.UserInfoTriggerCardGradeUpByTriggerID.Objects {
				if obj.TriggerID == dbTrigger.TriggerID {
					userModelCommonPos = i
					break
				}
			}
			if currentPos == -1 { // not in the current session but at login
				session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, dbTrigger.TriggerID)
				session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, nil)
			}
			if userModelCommonPos == -1 {
				session.UserModelCommon.UserInfoTriggerCardGradeUpByTriggerID.PushBack(dbTrigger)
			}
		}
		if currentPos != -1 {
			// overwrite the current trigger, this happen when we get 2 of the same card in gacha
			session.TriggerCardGradeUps[currentPos] = trigger.TriggerID
			session.TriggerCardGradeUps[currentPos+1] = trigger
		} else {
			// insert the trigger
			session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, trigger.TriggerID)
			session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, trigger)
		}
		if userModelCommonPos != -1 {
			// overwrite the current trigger, this happen when we get 2 of the same card in gacha
			session.UserModelCommon.UserInfoTriggerCardGradeUpByTriggerID.Objects[userModelCommonPos] = trigger
		} else {
			// insert the trigger
			session.UserModelCommon.UserInfoTriggerCardGradeUpByTriggerID.PushBack(trigger)
		}

		// save the trigger in db
		dbTrigger = trigger
		dbTrigger.BeforeLoveLevelLimit = dbTrigger.AfterLoveLevelLimit
		// db trigger when login have BeforeLoveLevelLimit = AfterLoveLevelLimit
		// if the 2 numbers are equal the level up don't show when we open the card.
		_, err = session.Db.Table("u_trigger_card_grade_up").Insert(&dbTrigger)
		utils.CheckErr(err)
	} else {
		// add trigger and remove from db
		// this is only caused by infoTrigger/read
		session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, trigger.TriggerID)
		session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, nil)
		session.UserModelCommon.UserInfoTriggerCardGradeUpByTriggerID.PushBack(trigger)
		_, err := session.Db.Table("u_trigger_card_grade_up").Where("trigger_id = ?", trigger.TriggerID).Delete(
			&model.TriggerCardGradeUp{})
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerCardGradeUps() generic.ObjectByObjectIDWrite[*model.TriggerCardGradeUp] {
	triggers := generic.ObjectByObjectIDWrite[*model.TriggerCardGradeUp]{}
	err := session.Db.Table("u_trigger_card_grade_up").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers.Objects)
	utils.CheckErr(err)
	triggers.Length = len(triggers.Objects)
	return triggers
}

func (session *Session) AddTriggerBasic(trigger model.TriggerBasic) {
	if trigger.TriggerID == 0 {
		trigger.TriggerID = time.Now().UnixNano()
	}
	trigger.UserID = session.UserStatus.UserID
	session.UserModelCommon.UserInfoTriggerBasicByTriggerID.PushBack(trigger)
	session.TriggerBasics = append(session.TriggerBasics, trigger.TriggerID)
	if trigger.IsNull {
		session.TriggerBasics = append(session.TriggerBasics, nil)
	} else {
		session.TriggerBasics = append(session.TriggerBasics, trigger)
	}
	if trigger.IsNull { // delete
		_, err := session.Db.Table("u_trigger_basic").Where("trigger_id = ?", trigger.TriggerID).Delete(
			&model.TriggerBasic{})
		utils.CheckErr(err)
	} else { // add
		_, err := session.Db.Table("u_trigger_basic").Insert(trigger)
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerBasics() generic.ObjectByObjectIDWrite[*model.TriggerBasic] {
	triggers := generic.ObjectByObjectIDWrite[*model.TriggerBasic]{}
	err := session.Db.Table("u_trigger_basic").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers.Objects)
	utils.CheckErr(err)
	triggers.Length = len(triggers.Objects)
	return triggers
}

func (session *Session) AddTriggerMemberLoveLevelUp(trigger model.TriggerMemberLoveLevelUp) {
	if trigger.TriggerID == 0 {
		trigger.TriggerID = time.Now().UnixNano()
	}

	trigger.UserID = session.UserStatus.UserID

	session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, trigger.TriggerID)
	if trigger.IsNull {
		session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, nil)
	} else {
		session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, trigger)
	}

	session.UserModelCommon.UserInfoTriggerMemberLoveLevelUpByTriggerID.PushBack(trigger)
	if !trigger.IsNull {
		_, err := session.Db.Table("u_trigger_member_love_level_up").Insert(trigger)
		utils.CheckErr(err)
	} else {
		_, err := session.Db.Table("u_trigger_member_love_level_up").Where("trigger_id = ?", trigger.TriggerID).Delete(
			&model.TriggerMemberLoveLevelUp{})
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerMemberLoveLevelUps() []model.TriggerMemberLoveLevelUp {
	triggers := []model.TriggerMemberLoveLevelUp{}
	err := session.Db.Table("u_trigger_member_love_level_up").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers)
	utils.CheckErr(err)
	return triggers
}
