package user_rule_description

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func userRuleDescriptionFinalizer(session *userdata.Session) {
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
			userdata.GenericDatabaseInsert(session, "u_rule_description", temp)
		}
	}
}

func init() {
	userdata.AddFinalizer(userRuleDescriptionFinalizer)
}
