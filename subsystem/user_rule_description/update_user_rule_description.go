package user_rule_description

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
)

func UpdateUserRuleDescription(session *userdata.Session, ruleDescriptionId int32) {
	// rule description is used for popup windows that tell you the rule of things
	session.UserModel.UserRuleDescriptionById.Set(ruleDescriptionId, client.UserRuleDescription{
		DisplayStatus: enum.RuleDescriptionDisplayStatusDisplay,
	})
}
