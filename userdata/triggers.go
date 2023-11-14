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
func (session *Session) AddTriggerCardGradeUp(id int64, trigger *model.TriggerCardGradeUp) {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	if trigger != nil {
		trigger.UserID = session.UserStatus.UserID
		trigger.TriggerID = id
	}
	if trigger != nil {
		dbTrigger := model.TriggerCardGradeUp{}

		exists, err := session.Db.Table("u_trigger_card_grade_up").
			Where("user_id = ? AND card_master_id = ?", trigger.UserID, trigger.CardMasterID).Get(&dbTrigger)
		utils.CheckErr(err)
		currentPos := -1
		if exists { // if the card has a trigger, we have to remove it
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
			if currentPos == -1 { // not in the current session but at login
				session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, dbTrigger.TriggerID)
				session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, nil)
			}
		}
		if currentPos != -1 {
			// overwrite the current trigger, this happen when we get 2 of the same card in gacha
			session.TriggerCardGradeUps[currentPos] = id
			session.TriggerCardGradeUps[currentPos+1] = *trigger
		} else {
			// insert the trigger
			session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, id)
			session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, trigger)
		}

		// save the trigger in db
		dbTrigger = *trigger
		dbTrigger.BeforeLoveLevelLimit = dbTrigger.AfterLoveLevelLimit
		// db trigger when login have BeforeLoveLevelLimit = AfterLoveLevelLimit
		// if the 2 numbers are equal the level up don't show when we open the card.
		_, err = session.Db.Table("u_trigger_card_grade_up").Insert(&dbTrigger)
		utils.CheckErr(err)
	} else {
		// add trigger and remove from db
		// this is only caused by a infoTrigger/read
		session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, id)
		session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, trigger)
		_, err := session.Db.Table("u_trigger_card_grade_up").Where("trigger_id = ?", id).Delete(
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

func (session *Session) AddTriggerBasic(id int64, trigger *model.TriggerBasic) {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	if trigger != nil {
		trigger.TriggerID = id
		trigger.UserID = session.UserStatus.UserID
	}
	session.TriggerBasics = append(session.TriggerBasics, id)
	session.TriggerBasics = append(session.TriggerBasics, trigger)
	if trigger != nil {
		_, err := session.Db.Table("u_trigger_basic").Insert(trigger)
		utils.CheckErr(err)
	} else {
		_, err := session.Db.Table("u_trigger_basic").Where("trigger_id = ?", id).Delete(
			&model.TriggerBasic{})
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

func (session *Session) AddTriggerMemberLoveLevelUp(id int64, trigger *model.TriggerMemberLoveLevelUp) {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	if trigger != nil {
		trigger.TriggerID = id
		trigger.UserID = session.UserStatus.UserID
	}
	session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, id)
	session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, trigger)
	if trigger != nil {
		_, err := session.Db.Table("u_trigger_member_love_level_up").Insert(trigger)
		utils.CheckErr(err)
	} else {
		_, err := session.Db.Table("u_trigger_member_love_level_up").Where("trigger_id = ?", id).Delete(
			&model.TriggerMemberLoveLevelUp{})
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerMemberLoveLevelUps() generic.ObjectByObjectIDWrite[*model.TriggerMemberLoveLevelUp] {
	triggers := generic.ObjectByObjectIDWrite[*model.TriggerMemberLoveLevelUp]{}
	err := session.Db.Table("u_trigger_member_love_level_up").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers.Objects)
	utils.CheckErr(err)
	triggers.Length = len(triggers.Objects)
	return triggers
}
