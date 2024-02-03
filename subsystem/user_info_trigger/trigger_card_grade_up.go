package user_info_trigger

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func RemoveTriggerCardGradeUp(session *userdata.Session, triggerId int64) {
	session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.SetNull(triggerId)
}

// card grade up trigger is responsible for showing the pop-up animation when openning a card after getting a new copy
// or right after performing a limit break using items
// Getting a new trigger also destroy old trigger, and we might have to update it
func AddTriggerCardGradeUp(session *userdata.Session, trigger client.UserInfoTriggerCardGradeUp) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.NextUniqueId()
	}
	session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Set(trigger.TriggerId, trigger)
}

func triggerCardGradeUpFinalizer(session *userdata.Session) {
	// keep only the latest one for each card
	keep := map[int32]int64{}
	for _, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map {
		if trigger != nil {
			if keep[trigger.CardMasterId] < trigger.TriggerId {
				keep[trigger.CardMasterId] = trigger.TriggerId
			}
		}
	}
	for triggerId, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map {
		if trigger != nil && triggerId != keep[trigger.CardMasterId] {
			delete(session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map, triggerId)
		}
	}
	// remove existing trigger in the database
	for _, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map {
		if trigger != nil {
			existingTrigger := client.UserInfoTriggerCardGradeUp{}
			exist, err := session.Db.Table("u_info_trigger_card_grade_up").
				Where("user_id = ? AND card_master_id = ?", session.UserId, trigger.CardMasterId).Get(&existingTrigger)
			utils.CheckErr(err)
			if exist {
				RemoveTriggerCardGradeUp(session, existingTrigger.TriggerId)
			}
		}
	}

	// finally make the change
	for triggerId, trigger := range session.UserModel.UserInfoTriggerCardGradeUpByTriggerId.Map {
		if trigger != nil {
			// triggers for login have BeforeLoveLevelLimit = AfterLoveLevelLimit
			// if the 2 numbers are equal the level up isn't shown when we open the card.
			dbTrigger := *trigger
			dbTrigger.BeforeLoveLevelLimit = dbTrigger.AfterLoveLevelLimit
			userdata.GenericDatabaseInsert(session, "u_info_trigger_card_grade_up", dbTrigger)
		} else {
			// remove from db
			_, err := session.Db.Table("u_info_trigger_card_grade_up").
				Where("user_id = ? AND trigger_id = ?", session.UserId, triggerId).
				Delete(&client.UserInfoTriggerCardGradeUp{})
			utils.CheckErr(err)
		}
	}
}

func init() {
	userdata.AddFinalizer(triggerCardGradeUpFinalizer)
}
