package userdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/utils"
)

func (session *Session) UpdateUserRuleDescription(ruleDescriptionId int32) {
	// rule description is used for popup windows that tell you the rule of things
	// only encountered in /referenceBook for now
	session.UserModel.UserRuleDescriptionById.Set(ruleDescriptionId, client.UserRuleDescription{
		DisplayStatus: enum.RuleDescriptionDisplayStatusDisplay,
	})
}

func ruleDescriptionFinalizer(session *Session) {
	for ruleDescriptionId, userRuleDescription := range session.UserModel.UserRuleDescriptionById.Map {
		affected, err := session.Db.Table("u_rule_description").Where("user_id = ? AND rule_description_id = ?",
			session.UserId, ruleDescriptionId).AllCols().
			Update(userRuleDescription)
		utils.CheckErr(err)
		if affected == 0 {
			// need to insert
			type Temp struct {
				RuleDescriptionId   int32                      `xorm:"pk 'rule_description_id'"`
				UserRuleDescription client.UserRuleDescription `xorm:"extends"`
			}
			temp := Temp{
				RuleDescriptionId:   ruleDescriptionId,
				UserRuleDescription: *userRuleDescription,
			}
			GenericDatabaseInsert(session, "u_rule_description", temp)
		}
	}
}

func init() {
	AddContentFinalizer(ruleDescriptionFinalizer)
}
