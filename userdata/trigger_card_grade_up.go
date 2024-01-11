package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
)

func (session *Session) RemoveTriggerCardGradeUp(triggerId int64) {
	session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.SetZero(triggerId)
}

// card grade up trigger is responsible for showing the pop-up animation when openning a card after getting a new copy
// or right after performing a limit break using items
// Getting a new trigger also destroy old trigger, and we might have to update it
func (session *Session) AddTriggerCardGradeUp(trigger client.UserInfoTriggerCardGradeUp) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.Time.UnixNano() + session.UniqueCount
		session.UniqueCount++
	}
	session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Set(trigger.TriggerId, generic.NewNullable(trigger))
}

func triggerCardGradeUpFinalizer(session *Session) {
	// remove existing triggers for the cards first
	for _, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map {
		if trigger.HasValue {
			existingTrigger := client.UserInfoTriggerCardGradeUp{}
			exist, err := session.Db.Table("u_info_trigger_card_grade_up").
				Where("user_id = ? AND card_master_id = ?", session.UserId, trigger.Value.CardMasterId).Get(&existingTrigger)
			utils.CheckErr(err)
			if exist {
				session.RemoveTriggerCardGradeUp(existingTrigger.TriggerId)
			}
		}
	}
	// make the change in the database
	for triggerId, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map {
		if trigger.HasValue {
			trigger.Value.BeforeLoveLevelLimit = trigger.Value.AfterLoveLevelLimit
			// db trigger when login have BeforeLoveLevelLimit = AfterLoveLevelLimit
			// if the 2 numbers are equal the level up isn't shown when we open the card.
			genericDatabaseInsert(session, "u_info_trigger_card_grade_up", trigger.Value)
		} else {
			// remove from db
			_, err := session.Db.Table("u_info_trigger_card_grade_up").Where("trigger_id = ?", triggerId).Delete(
				&client.UserInfoTriggerCardGradeUp{})
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(triggerCardGradeUpFinalizer)
}
