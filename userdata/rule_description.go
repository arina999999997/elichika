package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UpdateUserRuleDescription(ruleDescriptionID int) {
	// rule description is used for popup windows that tell you the rule of things
	// only encountered in /referenceBook for now, and all of them have display status 2
	// but some items can have display status 1 or 3
	// for now always use display status 2 until some exception happens
	userRuleDescription := model.UserRuleDescription{
		UserID:            session.UserStatus.UserID,
		RuleDescriptionID: ruleDescriptionID,
		DisplayStatus:     2,
	}
	session.UserModel.UserRuleDescriptionByID.PushBack(userRuleDescription)
	affected, err := session.Db.Table("u_rule_description").Where("user_id = ? AND rule_description_id = ?",
		userRuleDescription.UserID, userRuleDescription.RuleDescriptionID).AllCols().
		Update(userRuleDescription)
	utils.CheckErr(err)
	if affected > 0 {
		return
	}
	// need to insert
	_, err = session.Db.Table("u_rule_description").Insert(userRuleDescription)
	utils.CheckErr(err)
}

func init() {
	addGenericTableFieldPopulator("u_rule_description", "UserRuleDescriptionByID")
}
